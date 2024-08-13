package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	var err error
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)
    
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        return err
    }
	
    err = db.Ping()
    if err != nil {
        return err
    }

    // Create tables
    tables := []string{"all", "unoptimized", "optimized", "parallel", "parallelExtended"}
    for _, table := range tables {
        query := fmt.Sprintf(`
            CREATE TABLE IF NOT EXISTS public.%s (
                id SERIAL PRIMARY KEY,
                results JSONB NOT NULL,
                created_at TIMESTAMP NOT NULL
            )
        `, table)
        _, err = db.Exec(query)
        if err != nil {
            return fmt.Errorf("error creating table %s: %v", table, err)
        }
    }

    return nil
}

// StoreResponse stores the response in the database
func StoreResponse(tableName string, response map[string]interface{}) error {
	jsonResponse, err := json.Marshal(response)
    if err != nil {
        return fmt.Errorf("error marshaling response: %v", err)
    }

    // Insert the response into the specified table
    query := fmt.Sprintf(`
        INSERT INTO public.%s (results, created_at)
        VALUES ($1, $2)
    `, tableName)

    _, err = db.Exec(query, string(jsonResponse), time.Now())
    if err != nil {
        return fmt.Errorf("error inserting into %s: %v", tableName, err)
    }

    return nil
	
}

// RetrieveRecords retrieves all records from the specified table
func RetrieveRecords(tableName string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`
		SELECT id, results, created_at
		FROM public.%s
		ORDER BY created_at DESC
	`, tableName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying %s: %v", tableName, err)
	}
	defer rows.Close()

	var records []map[string]interface{}

	for rows.Next() {
		var id int
		var results string
		var createdAt time.Time

		err := rows.Scan(&id, &results, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		var resultsMap map[string]interface{}
		err = json.Unmarshal([]byte(results), &resultsMap)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling results: %v", err)
		}

		record := map[string]interface{}{
			"id":         id,
			"results":    resultsMap,
			"created_at": createdAt,
		}

		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("no records found")
	}

	return records, nil
}


// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}