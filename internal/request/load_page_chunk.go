package request

type Stack struct {
	ID    string `json:"id"`
	Index int    `json:"index"`
	Table string `json:"table"`
}

type Cursor struct {
	Stack [][]Stack `json:"stack"`
}

type LoadPageChunkRequest struct {
	PageId          string `json:"pageId"`
	Limit           int    `json:"limit"`
	ChunkNumber     int    `json:"chunkNumber"`
	VerticalColumns bool   `json:"verticalColumns"`
	Cursor          Cursor `json:"cursor"`
}
