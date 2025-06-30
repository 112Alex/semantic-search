package semhtml

import (
	"io"

	"github.com/yourname/semantic-search/internal/index"
	"github.com/yourname/semantic-search/internal/parser"
	"github.com/yourname/semantic-search/internal/search"
	"github.com/yourname/semantic-search/internal/utils"
)

// NewEngineFromReader парсит HTML из reader и строит индекс.
func NewEngineFromReader(r io.Reader) (*Engine, error) {
	nodes, err := parser.ParseTextNodes(r)
	if err != nil {
		return nil, err
	}

	docs := make([]index.Document, len(nodes))
	for i, n := range nodes {
		docs[i] = index.NewDocument(n)
	}
	idx := index.Build(docs)
	return &Engine{idx: idx}, nil
}

// Search выполняет семантический поиск по запросу и возвращает топ N результатов.
func (e *Engine) Search(query string, topN int) []Result {
	tokens := tokenize(query)
	res := search.Query(e.idx, tokens, topN)
	// заполняем CSS-селектор
	out := make([]Result, len(res))
	for i, r := range res {
		out[i] = Result{
			Score:    r.Score,
			Selector: utils.CSSSelector(r.Document.Node),
			Text:     r.Document.Node.Data,
		}
	}
	return out
}

// копия токенизации из internal/index, оставляем приватной здесь для удобства.
func tokenize(text string) []string {
	// простой split по не-буквам/цифрам
	tokens := make([]string, 0)
	current := ""
	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') {
			current += string(r)
		} else {
			if len(current) > 0 {
				tokens = append(tokens, current)
				current = ""
			}
		}
	}
	if len(current) > 0 {
		tokens = append(tokens, current)
	}
	return tokens
}
