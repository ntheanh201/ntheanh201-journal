package request

type PermissionRecord struct {
	ID      string `json:"id"`
	Table   string `json:"table"`
	SpaceID string `json:"spaceId"`
}

type SignedURLRequest struct {
	URL              string            `json:"url"`
	PermissionRecord *PermissionRecord `json:"permissionRecord"`
}

type GetSignedFileUrlsRequest struct {
	URLs []SignedURLRequest `json:"urls"`
}
