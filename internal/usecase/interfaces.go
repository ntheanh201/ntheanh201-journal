package usecase

import (
	"context"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/response"
)

type (
	Page interface {
		GetPages(context.Context) (response.DatabaseQueryResponse, error)
		GetPage(context.Context, entity.ObjectID) (entity.Page, error)
		GetBlockChildren(ctx context.Context, id entity.ObjectID) (response.BlockChildrenResponse, error)
		LoadPageChunkV3(ctx context.Context, slug string) (response.LoadPageChunkResponse, error)
	}

	PageNotionWebAPI interface {
		GetNotionPages(context.Context) (response.DatabaseQueryResponse, error)
		GetNotionPage(ctx context.Context, id entity.ObjectID) (entity.Page, error)
		GetNotionBlockChildren(ctx context.Context, id entity.ObjectID, pageSize int) (response.BlockChildrenResponse, error)
		QueryNotionPageBySlug(ctx context.Context, slug string) (response.DatabaseQueryResponse, error)
		LoadPageChunkV3(ctx context.Context, id entity.ObjectID) (response.LoadPageChunkResponse, error)
	}
)
