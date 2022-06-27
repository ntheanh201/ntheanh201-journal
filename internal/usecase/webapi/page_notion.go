package webapi

import (
	"context"
	"fmt"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/response"
)

const databaseId = "68f077a6dfb346358f219875e80ea72c"

type PageNotionWebAPI struct {
	notionClient *Client
}

func New() *PageNotionWebAPI {
	return &PageNotionWebAPI{
		notionClient: NewClient("secret_V9aVN7aMqWiLBYXxzPbollmYWlmVeyqmhi1Y9nRfQbd"),
	}
}

func (p *PageNotionWebAPI) GetNotionPages(ctx context.Context) (response.DatabaseQueryResponse, error) {
	res, err := p.notionClient.queryDatabase(ctx, databaseId, &response.QueryDatabase{})
	if err != nil {
		_ = fmt.Errorf("journal: failed to get Notion database: %w", err)
	}
	return res, nil
}

func (p *PageNotionWebAPI) GetNotionPage(ctx context.Context, id entity.ObjectID) (entity.Page, error) {
	// TODO: implement me
	return entity.Page{}, nil
}
