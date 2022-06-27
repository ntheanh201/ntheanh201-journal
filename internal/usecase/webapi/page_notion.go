package webapi

import (
	"context"
	"fmt"
	"log"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/response"
	"os"
)

const databaseId = "68f077a6dfb346358f219875e80ea72c"

type PageNotionWebAPI struct {
	notionClient *Client
}

func New() *PageNotionWebAPI {
	notionAPIKey, err := os.LookupEnv("NOTION_API_KEY")
	if err {
		log.Fatalf("journal: environment variable not declared: PG_URL")
	}
	return &PageNotionWebAPI{
		notionClient: NewClient(notionAPIKey),
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
