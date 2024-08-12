package routes

import (
	"net/http"

	"github.com/aleedurrani/TimeComplexity/pkg/optimized"
	"github.com/aleedurrani/TimeComplexity/internal/utils/commonFunctions"
)

func OptimizedHandler(w http.ResponseWriter, r *http.Request) {
	commonFunctions.HandleSingleMethod(w, r, optimized.OptimizedCountAll)
}