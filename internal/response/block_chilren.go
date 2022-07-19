package response

import "ntheanh201-journal/internal/entity"

type BlockChildrenResponse struct {
	Object     *string        `json:"object"`
	Results    []entity.Block `json:"results"`
	NextCursor *string        `json:"next_cursor"`
	HasMore    bool           `json:"has_more"`
	Type       *string        `json:"type"`
	Block      interface{}    `json:"block"`
}
