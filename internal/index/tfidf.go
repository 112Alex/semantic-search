package index

// Term Frequency — Inverse Document Frequency (TFIDF)

import "math"

// Index хранит документы и глобальную статистику IDF.
type Index struct {
	Docs []Document
	IDF  map[string]float64 // term -> idf
}

// Build строит индекс по коллекции документов.
func Build(docs []Document) Index {
	df := make(map[string]int) // term -> document frequency
	for _, doc := range docs {
		seen := make(map[string]struct{})
		for term := range doc.Vector {
			if _, ok := seen[term]; !ok {
				df[term]++
				seen[term] = struct{}{}
			}
		}
	}

	idf := make(map[string]float64, len(df))
	N := float64(len(docs))
	for term, freq := range df {
		idf[term] = math.Log((N+1)/(float64(freq)+1)) + 1 // сглажённая формула
	}

	// умножаем tf на idf для каждого документа (in-place копия)
	indexedDocs := make([]Document, len(docs))
	for i, doc := range docs {
		vec := make(map[string]float64, len(doc.Vector))
		for term, tf := range doc.Vector {
			vec[term] = tf * idf[term]
		}
		indexedDocs[i] = Document{Node: doc.Node, Vector: vec}
	}

	return Index{Docs: indexedDocs, IDF: idf}
}
