package response

import "ntheanh201-journal/internal/entity"

type QueryDatabase struct {
}

type DatabaseQueryResponse struct {
	Results    []entity.Page `json:"results"`
	HasMore    bool          `json:"has_more"`
	NextCursor *string       `json:"next_cursor"`
}
