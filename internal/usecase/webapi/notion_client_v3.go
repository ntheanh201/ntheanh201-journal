package webapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/request"
	"ntheanh201-journal/internal/response"
	"ntheanh201-journal/internal/utils"
)

const (
	apiUrlV3     = "https://www.notion.so/api"
	apiVersionV3 = "v3"
)

func NewClientV3(opts ...ClientOption) *Client {
	u, err := url.Parse(apiUrlV3)
	if err != nil {
		panic(err)
	}
	c := &Client{
		baseUrl:    u,
		httpClient: http.DefaultClient,
		apiVersion: apiVersionV3,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) newRequestV3(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s%s", c.apiVersion, url))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) loadPageChunk(ctx context.Context, id entity.ObjectID) (result response.LoadPageChunkResponse, err error) {
	body := &bytes.Buffer{}

	query := &request.LoadPageChunkRequest{
		PageId:          utils.ParsePageId(id.String()),
		Limit:           100,
		ChunkNumber:     0,
		VerticalColumns: false,
	}

	err = json.NewEncoder(body).Encode(query)
	if err != nil {
		return response.LoadPageChunkResponse{}, fmt.Errorf("notion: failed to encode filter to JSON: %w", err)
	}

	req, err := c.newRequestV3(ctx, http.MethodPost, "/loadPageChunk", body)
	if err != nil {
		return response.LoadPageChunkResponse{}, fmt.Errorf("notion: invalid load page chunk request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return response.LoadPageChunkResponse{}, fmt.Errorf("notion: failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response.LoadPageChunkResponse{}, fmt.Errorf("notion: failed to load page chunk: %w", res)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return response.LoadPageChunkResponse{}, fmt.Errorf("notion: failed to parse HTTP response: %w", err)
	}

	return result, nil
}
