package response

import "ntheanh201-journal/internal/entity"

type FilterContains struct {
	Contains string `json:"contains,omitempty"`
}

type QueryDatabaseFilter struct {
	Property string         `json:"property,omitempty"`
	Value    FilterContains `json:"rich_text,omitempty"`
}

type QueryDatabase struct {
	Filter *QueryDatabaseFilter `json:"filter,omitempty"`
}

type DatabaseQueryResponse struct {
	Results    []entity.Page `json:"results"`
	HasMore    bool          `json:"has_more"`
	NextCursor *string       `json:"next_cursor"`
}
