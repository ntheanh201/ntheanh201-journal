package request

type LoadPageChunkRequest struct {
	PageId          string `json:"pageId"`
	Limit           int    `json:"limit"`
	ChunkNumber     int    `json:"chunkNumber"`
	VerticalColumns bool   `json:"verticalColumns"`
}
