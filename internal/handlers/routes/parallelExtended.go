package routes

import (
	"net/http"

	"github.com/aleedurrani/TimeComplexity/internal/utils/commonFunctions"
	"github.com/aleedurrani/TimeComplexity/pkg/parallelExtended"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
)

func ParallelExtendedHandler(w http.ResponseWriter, r *http.Request) {
	numRoutines := commonFunctions.GetNumRoutines(w, r)
	if numRoutines == 0 {
		return
	}
	commonFunctions.HandleSingleMethod(w, r, func() helperFunctions.Counts {
		return parallelExtended.ParallelCountAll(numRoutines)
	})
}