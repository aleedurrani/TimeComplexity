package routes

import (
	"net/http"

	"github.com/aleedurrani/TimeComplexity/pkg/parallel"
	"github.com/aleedurrani/TimeComplexity/internal/utils/commonFunctions"
)

func ParallelHandler(w http.ResponseWriter, r *http.Request) {
	commonFunctions.HandleSingleMethod(w, r, parallel.ParallelCountAll)
}