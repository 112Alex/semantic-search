package search

import (
	"math"
	"sort"

	"github.com/yourname/semantic-search/internal/index"
)

// Result представляет одну совпавшую запись с оценкой.
type Result struct {
	Score    float64
	Document index.Document
}

// cosineSimilarity считает косинус между двумя векторами map[string]float64.
func cosineSimilarity(a, b map[string]float64) float64 {
	var dot, normA, normB float64
	for k, va := range a {
		vb := b[k]
		dot += va * vb
		normA += va * va
	}
	for _, vb := range b {
		normB += vb * vb
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}

// Query выполняет поиск по индексу. queryTokens – уже токенизированные слова.
func Query(idx index.Index, queryTokens []string, topN int) []Result {
	// строим вектор запроса (tf * idf)
	qVecRaw := make(map[string]float64)
	for _, t := range queryTokens {
		qVecRaw[t] += 1
	}
	// нормализуем tf
	l := float64(len(queryTokens))
	for k, v := range qVecRaw {
		qVecRaw[k] = v / l * idx.IDF[k]
	}

	res := make([]Result, 0, len(idx.Docs))
	for _, doc := range idx.Docs {
		score := cosineSimilarity(doc.Vector, qVecRaw)
		if score > 0 {
			res = append(res, Result{Score: score, Document: doc})
		}
	}

	sort.Slice(res, func(i, j int) bool { return res[i].Score > res[j].Score })
	if len(res) > topN {
		res = res[:topN]
	}
	return res
}
