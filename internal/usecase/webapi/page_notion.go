package webapi

import (
	"context"
	"fmt"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/response"
	"os"
)

const databaseId = "68f077a6dfb346358f219875e80ea72c"

type PageNotionWebAPI struct {
	notionClient *Client
}

func New() *PageNotionWebAPI {
	notionAPIKey := os.Getenv("NOTION_API_KEY")
	return &PageNotionWebAPI{
		notionClient: NewClient(notionAPIKey),
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
