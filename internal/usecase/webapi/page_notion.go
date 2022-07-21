package webapi

import (
	"context"
	"fmt"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/request"
	"ntheanh201-journal/internal/response"
	"os"
)

const databaseId = "68f077a6dfb346358f219875e80ea72c"

type PageNotionWebAPI struct {
	notionClient   *Client
	notionClientV3 *Client
}

func New() *PageNotionWebAPI {
	notionAPIKey := os.Getenv("NOTION_API_KEY")
	return &PageNotionWebAPI{
		notionClient:   NewClient(notionAPIKey),
		notionClientV3: NewClientV3(),
	}
}

func (p *PageNotionWebAPI) GetNotionPages(ctx context.Context) (response.DatabaseQueryResponse, error) {
	res, err := p.notionClient.queryDatabase(ctx, databaseId, nil)
	if err != nil {
		_ = fmt.Errorf("journal: failed to get Notion database: %w", err)
	}
	return res, nil
}

func (p *PageNotionWebAPI) GetNotionPage(ctx context.Context, id entity.ObjectID) (entity.Page, error) {
	res, err := p.notionClient.retrievePage(ctx, id)
	if err != nil {
		_ = fmt.Errorf("journal: failed to retrieve Notion page: %s - %w", id, err)
	}
	return res, nil
}

func (p *PageNotionWebAPI) GetNotionBlockChildren(ctx context.Context, id entity.ObjectID, pageSize int) (response.BlockChildrenResponse, error) {
	res, err := p.notionClient.retrieveBlockChildren(ctx, id, pageSize)
	if err != nil {
		_ = fmt.Errorf("journal: failed to retrieve Notion block children: %s - %w", id, err)
	}
	return res, nil
}

func (p *PageNotionWebAPI) QueryNotionPageBySlug(ctx context.Context, slug string) (response.DatabaseQueryResponse, error) {
	query := response.QueryDatabase{
		Filter: &response.QueryDatabaseFilter{
			Property: "slug",
			Value:    response.FilterContains{Contains: slug},
		},
	}
	res, err := p.notionClient.queryDatabase(ctx, databaseId, &query)
	if err != nil {
		_ = fmt.Errorf("journal: failed to retrieve Notion page by slug: %s - %w", slug, err)
		return response.DatabaseQueryResponse{}, err
	}
	return res, nil
}

func (p *PageNotionWebAPI) LoadPageChunkV3(ctx context.Context, chunkRequest request.LoadPageChunkRequest) (response.LoadPageChunkResponse, error) {
	res, err := p.notionClientV3.loadPageChunk(ctx, chunkRequest)
	if err != nil {
		_ = fmt.Errorf("journal: failed to load page chunk Notion: %s - %w", err)
		return response.LoadPageChunkResponse{}, err
	}
	return res, nil
}

func (p *PageNotionWebAPI) GetSignedFileUrls(ctx context.Context, req request.GetSignedFileUrlsRequest) (response.GetSignedUrlsResponse, error) {
	res, err := p.notionClientV3.getSignedFileUrls(ctx, req)
	if err != nil {
		_ = fmt.Errorf("journal: failed to get signed file urls Notion: %w", err)
		return response.GetSignedUrlsResponse{}, err
	}
	return res, nil
}
