package response

type GetSignedUrlsResponse struct {
	SignedURLS []string               `json:"signedUrls"`
	RawJSON    map[string]interface{} `json:"-"`
}
