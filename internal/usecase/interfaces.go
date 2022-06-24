package usecase

import (
	"context"
	"ntheanh201-journal/internal/entity"
)

type (
	Page interface {
		GetPages(context.Context) ([]entity.Page, error)
		GetPage(context.Context, entity.ObjectID) (entity.Page, error)
	}

	PageNotionWebAPI interface {
		GetPages() ([]entity.Page, error)
		GetPage(id entity.ObjectID) (entity.Page, error)
	}
)
