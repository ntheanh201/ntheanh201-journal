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
	}

	PageNotionWebAPI interface {
		GetNotionPages(context.Context) (response.DatabaseQueryResponse, error)
		GetNotionPage(ctx context.Context, id entity.ObjectID) (entity.Page, error)
	}
)