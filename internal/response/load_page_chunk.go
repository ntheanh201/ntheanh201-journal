package response

import "ntheanh201-journal/internal/entity"

type stack struct {
	ID    string `json:"id"`
	Index int    `json:"index"`
	Table string `json:"table"`
}

type cursor struct {
	Stack [][]stack `json:"stack"`
}

type LoadPageChunkResponse struct {
	RecordMap *entity.RecordMap      `json:"recordMap"`
	Cursor    cursor                 `json:"cursor"`
	RawJSON   map[string]interface{} `json:"-"`
}
