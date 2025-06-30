package index

import (
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

// Document представляет текстовый узел как отдельный документ индекса.
// Node – указатель на исходный html.Node, Vector – TF веса (без IDF).
// Безопасно хранить map[string]float64, потому что размер словаря небольшой.
// Comments: рус.
type Document struct {
	Node   *html.Node
	Vector map[string]float64 // term -> tf
}

// tokenize простым делителем по буквам, приводим к lower-case.
func tokenize(text string) []string {
	return strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		// Разделяем всё, что не является буквой (unicode) или цифрой.
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

// NewDocument создаёт Document из текстового содержимого узла.
func NewDocument(n *html.Node) Document {
	tokens := tokenize(n.Data)
	tf := make(map[string]float64, len(tokens))
	for _, tok := range tokens {
		tf[tok] += 1
	}
	// нормализуем на длину документа (l2 или просто длину) – выберем простую норму.
	length := float64(len(tokens))
	if length > 0 {
		for k, v := range tf {
			tf[k] = v / length
		}
	}

	return Document{Node: n, Vector: tf}
}
