package routes

import (
	"net/http"

	"github.com/aleedurrani/TimeComplexity/pkg/unoptimized"
	"github.com/aleedurrani/TimeComplexity/pkg/utils/helperFunctions"
	"github.com/aleedurrani/TimeComplexity/internal/utils/commonFunctions"
)

func UnoptimizedHandler(w http.ResponseWriter, r *http.Request) {
	commonFunctions.HandleSingleMethod(w, r, func() helperFunctions.Counts {
		return helperFunctions.Counts{
			Word:      unoptimized.CountWords(),
			Punct:     unoptimized.CountPunctuation(),
			Vowel:     unoptimized.CountVowels(),
			Sentence:  unoptimized.CountSentences(),
			Paragraph: unoptimized.CountParagraphs(),
			Digit:     unoptimized.CountDigits(),
		}
	})
}