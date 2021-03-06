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
	"ntheanh201-journal/internal/response"
)

const (
	apiUrl        = "https://api.notion.com"
	apiVersion    = "v1"
	notionVersion = "2022-02-22"
)

type Client struct {
	baseUrl    *url.URL
	httpClient *http.Client
	apiVersion string
	databaseId string
	apiKey     string
}

// ClientOption to configure API client
type ClientOption func(*Client)

func NewClient(apiKey string, opts ...ClientOption) *Client {
	u, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}
	c := &Client{
		apiKey:     apiKey,
		baseUrl:    u,
		httpClient: http.DefaultClient,
		apiVersion: apiVersion,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithHTTPClient overrides the default http.Client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s%s", c.apiVersion, url))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
	req.Header.Set("Notion-Version", notionVersion)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) queryDatabase(ctx context.Context, id string, query *response.QueryDatabase) (result response.DatabaseQueryResponse, err error) {
	body := &bytes.Buffer{}

	if query != nil {
		err = json.NewEncoder(body).Encode(query)
		if err != nil {
			return response.DatabaseQueryResponse{}, fmt.Errorf("notion: failed to encode filter to JSON: %w", err)
		}
	}

	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("/databases/%v/query", id), body)
	if err != nil {
		return response.DatabaseQueryResponse{}, fmt.Errorf("notion: invalid request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return response.DatabaseQueryResponse{}, fmt.Errorf("notion: failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response.DatabaseQueryResponse{}, fmt.Errorf("notion: failed to query database: %w", res)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return response.DatabaseQueryResponse{}, fmt.Errorf("notion: failed to parse HTTP response: %w", err)
	}

	return result, nil
}

func (c *Client) retrievePage(ctx context.Context, id entity.ObjectID) (result entity.Page, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/pages/%v", id), &bytes.Buffer{})
	if err != nil {
		return entity.Page{}, fmt.Errorf("notion: invalid retrieve page request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return entity.Page{}, fmt.Errorf("notion: failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return entity.Page{}, fmt.Errorf("notion: failed to retrieve page: %w", res)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return entity.Page{}, fmt.Errorf("notion: failed to parse HTTP response: %w", err)
	}

	return result, nil
}

func (c *Client) retrieveBlockChildren(ctx context.Context, id entity.ObjectID, pageSize int) (result response.BlockChildrenResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/blocks/%v/children?page_size=%d", id, pageSize), &bytes.Buffer{})
	if err != nil {
		return response.BlockChildrenResponse{}, fmt.Errorf("notion: invalid retrieve block children request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return response.BlockChildrenResponse{}, fmt.Errorf("notion: failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response.BlockChildrenResponse{}, fmt.Errorf("notion: failed to retrieve block children: %w", res)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return response.BlockChildrenResponse{}, fmt.Errorf("notion: failed to parse HTTP response: %w", err)
	}

	return result, nil
}
