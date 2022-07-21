package request

type QuerySort struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Property  string `json:"property"`
	Direction string `json:"direction"`
}

type QueryAggregate struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Property        string `json:"property"`
	ViewType        string `json:"view_type"`
	AggregationType string `json:"aggregation_type"`
}

type QueryAggregation struct {
	Property   string `json:"property"`
	Aggregator string `json:"aggregator"`
}

type Query struct {
	Sort         []QuerySort            `json:"sort"`
	Aggregate    []QueryAggregate       `json:"aggregate"`
	Aggregations []QueryAggregation     `json:"aggregations"`
	Filter       map[string]interface{} `json:"filter"`
}

type QueryCollectionRequest struct {
	Collection struct {
		ID      string `json:"id"`
		SpaceID string `json:"spaceId,omitempty"`
	} `json:"collection"`
	CollectionView struct {
		ID      string `json:"id"`
		SpaceID string `json:"spaceId,omitempty"`
	} `json:"collectionView"`
	Loader interface{} `json:"loader"`
}
