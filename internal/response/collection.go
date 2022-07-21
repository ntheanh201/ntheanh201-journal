package response

import "ntheanh201-journal/internal/entity"

type CollectionGroupResults struct {
	Type     string   `json:"type"`
	BlockIds []string `json:"blockIds"`
	Total    int      `json:"total"`
}
type ReducerResults struct {
	CollectionGroupResults *CollectionGroupResults `json:"collection_group_results"`
}

type QueryCollectionResponse struct {
	RecordMap *entity.RecordMap `json:"recordMap"`
	Result    struct {
		Type           string          `json:"type"`
		ReducerResults *ReducerResults `json:"reducerResults"`
	} `json:"result"`
	RawJSON map[string]interface{} `json:"-"`
}
