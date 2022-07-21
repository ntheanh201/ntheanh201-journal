package response

import (
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/request"
)

type LoadPageChunkResponse struct {
	RecordMap *entity.RecordMap `json:"recordMap"`
	Cursor    request.Cursor    `json:"cursor"`
	//RawJSON   map[string]interface{} `json:"-"`
}
