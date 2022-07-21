package webapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/request"
	"ntheanh201-journal/internal/response"
)

var jsonit = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	apiUrlV3     = "https://www.notion.so"
	apiTag       = "api"
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
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/%s%s", apiTag, c.apiVersion, url))
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

func (c *Client) loadPageChunk(ctx context.Context, reqBody request.LoadPageChunkRequest) (result response.LoadPageChunkResponse, err error) {
	body := &bytes.Buffer{}

	err = json.NewEncoder(body).Encode(reqBody)
	if err != nil {
		return response.LoadPageChunkResponse{}, fmt.Errorf("notion: failed to encode filter to JSON: %w", err)
	}

	req, err := c.newRequestV3(ctx, http.MethodPost, "/loadCachedPageChunk", body)
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

	//err = json.NewDecoder(res.Body).Decode(&result)
	//if err != nil {
	//	return response.LoadPageChunkResponse{}, fmt.Errorf("notion: failed to parse HTTP response: %w", err)
	//}
	resBody, err := ioutil.ReadAll(res.Body)
	var rsp response.LoadPageChunkResponse

	err = jsonit.Unmarshal(resBody, &rsp)
	if err != nil {
		fmt.Sprintf("Error: json.Unmarshal() failed with %s\n. Body:\n%s\n", err, string(resBody))
		return response.LoadPageChunkResponse{}, err
	}

	//if err = ParseRecordMap(result.RecordMap); err != nil {
	//	return response.LoadPageChunkResponse{}, err
	//}

	return rsp, nil

}

const (
	// TableSpace those are Record.Type and determine the type of Record.Value
	TableSpace          = "space"
	TableActivity       = "activity"
	TableBlock          = "block"
	TableNotionUser     = "notion_user"
	TableUserRoot       = "user_root"
	TableUserSettings   = "user_settings"
	TableCollection     = "collection"
	TableCollectionView = "collection_view"
	TableComment        = "comment"
	TableDiscussion     = "discussion"
)

func parseRecord(table string, r *entity.Record) error {
	// it's ok if some records don't return a value
	if len(r.Value) == 0 {
		return nil
	}
	if r.Table == "" {
		r.Table = table
	} else {
		// TODO: probably never happens
		//panicIf(r.Table != table)
	}

	// set Block/Space etc. based on TableView type
	var pRawJSON *map[string]interface{}
	var obj interface{}
	switch table {
	case TableActivity:
		r.Activity = &entity.Activity{}
		obj = r.Activity
		pRawJSON = &r.Activity.RawJSON
	case TableBlock:
		r.Block = &entity.BlockV3{}
		obj = r.Block
		pRawJSON = &r.Block.RawJSON
	//case TableNotionUser:
	//	r.NotionUser = &NotionUser{}
	//	obj = r.NotionUser
	//	pRawJSON = &r.NotionUser.RawJSON
	//case TableUserRoot:
	//	r.UserRoot = &UserRoot{}
	//	obj = r.UserRoot
	//	pRawJSON = &r.UserRoot.RawJSON
	//case TableUserSettings:
	//	r.UserSettings = &UserSettings{}
	//	obj = r.UserSettings
	//	pRawJSON = &r.UserSettings.RawJSON
	//case TableSpace:
	//	r.Space = &Space{}
	//	obj = r.Space
	//	pRawJSON = &r.Space.RawJSON
	case TableCollection:
		r.Collection = &entity.Collection{}
		obj = r.Collection
		pRawJSON = &r.Collection.RawJSON
	case TableCollectionView:
		r.CollectionView = &entity.CollectionView{}
		obj = r.CollectionView
		pRawJSON = &r.CollectionView.RawJSON
		//case TableDiscussion:
		//	r.Discussion = &Discussion{}
		//	obj = r.Discussion
		//	pRawJSON = &r.Discussion.RawJSON
		//case TableComment:
		//	r.Comment = &Comment{}
		//	obj = r.Comment
		//	pRawJSON = &r.Comment.RawJSON
	}
	if obj == nil {
		return fmt.Errorf("unsupported table '%s'", r.Table)
	}
	if err := jsonit.Unmarshal(r.Value, pRawJSON); err != nil {
		return err
	}
	id := (*pRawJSON)["id"]
	if id != nil {
		r.ID = id.(string)
	}
	if err := jsonit.Unmarshal(r.Value, &obj); err != nil {
		return err
	}
	return nil
}

func ParseRecordMap(recordMap *entity.RecordMap) error {
	for _, r := range recordMap.Activities {
		if err := parseRecord(TableActivity, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.Blocks {
		if err := parseRecord(TableBlock, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.Spaces {
		if err := parseRecord(TableSpace, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.NotionUsers {
		if err := parseRecord(TableNotionUser, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.UsersRoot {
		if err := parseRecord(TableUserRoot, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.UserSettings {
		if err := parseRecord(TableUserSettings, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.CollectionViews {
		if err := parseRecord(TableCollectionView, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.Collections {
		if err := parseRecord(TableCollection, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.Discussions {
		if err := parseRecord(TableDiscussion, r); err != nil {
			return err
		}
	}

	for _, r := range recordMap.Comments {
		if err := parseRecord(TableComment, r); err != nil {
			return err
		}
	}

	return nil
}

//func (c *Client) queryCollection(ctx context.Context, req *request.QueryCollectionRequest, query *request.Query) (result response.QueryCollectionResponse, err error) {
//	body := &bytes.Buffer{}
//
//	query := &request
//
//	err = json.NewEncoder(body).Encode(query)
//	if err != nil {
//		return response.QueryCollectionResponse{}, fmt.Errorf("notion: failed to encode filter to JSON: %w", err)
//	}
//
//	req, err := c.newRequestV3(ctx, http.MethodPost, "/loadPageChunk", body)
//	if err != nil {
//		return response.QueryCollectionResponse{}, fmt.Errorf("notion: invalid load page chunk request: %w", err)
//	}
//	res, err := c.httpClient.Do(req)
//	if err != nil {
//		return response.QueryCollectionResponse{}, fmt.Errorf("notion: failed to make HTTP request: %w", err)
//	}
//	defer res.Body.Close()
//
//	if res.StatusCode != http.StatusOK {
//		return response.QueryCollectionResponse{}, fmt.Errorf("notion: failed to load page chunk: %w", res)
//	}
//
//	err = json.NewDecoder(res.Body).Decode(&result)
//	if err != nil {
//		return response.QueryCollectionResponse{}, fmt.Errorf("notion: failed to parse HTTP response: %w", err)
//	}
//
//	return result, nil
//}

func (c *Client) getSignedFileUrls(ctx context.Context, input request.GetSignedFileUrlsRequest) (result response.GetSignedUrlsResponse, err error) {
	body := &bytes.Buffer{}

	err = json.NewEncoder(body).Encode(input)
	if err != nil {
		return response.GetSignedUrlsResponse{}, fmt.Errorf("notion: failed to encode filter to JSON: %w", err)
	}

	req, err := c.newRequestV3(ctx, http.MethodPost, "/getSignedFileUrls", body)
	if err != nil {
		return response.GetSignedUrlsResponse{}, fmt.Errorf("notion: invalid load page chunk request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return response.GetSignedUrlsResponse{}, fmt.Errorf("notion: failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response.GetSignedUrlsResponse{}, fmt.Errorf("notion: failed to load page chunk: %w", res)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return response.GetSignedUrlsResponse{}, fmt.Errorf("notion: failed to parse HTTP response: %w", err)
	}

	return result, nil
}
