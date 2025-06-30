package semhtml

import (
	"github.com/yourname/semantic-search/internal/index"
	"github.com/yourname/semantic-search/internal/search"
)

// Result публично экспортируется из pkg.
// Содержит счёт, текст и CSS-селектор.
// Текст берём из html.Node.Data.
type Result struct {
	Score    float64
	Selector string
	Text     string
}

// convertResults internal->public
func convertResults(in []search.Result) []Result {
	out := make([]Result, len(in))
	for i, r := range in {
		out[i] = Result{
			Score:    r.Score,
			Selector: "", // TODO позже
			Text:     r.Document.Node.Data,
		}
	}
	return out
}

// Engine – единая точка работы: Parse -> Index -> Query.
// Пока минимальный API.
type Engine struct {
	idx index.Index
}
