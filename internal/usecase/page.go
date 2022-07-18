package usecase

import (
	"context"
	"fmt"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/response"
)

type PageUseCase struct {
	webAPI PageNotionWebAPI
}

func (uc *PageUseCase) GetPages(ctx context.Context) (response.DatabaseQueryResponse, error) {
	pages, err := uc.webAPI.GetNotionPages(ctx)
	if err != nil {
		return response.DatabaseQueryResponse{}, fmt.Errorf("PageUseCase - webAPI.GetPages: %w", err)
	}
	return pages, nil
}

func (uc *PageUseCase) GetPage(ctx context.Context, id entity.ObjectID) (entity.Page, error) {
	page, err := uc.webAPI.GetNotionPage(ctx, id)
	if err != nil {
		return entity.Page{}, fmt.Errorf("PageUseCase - webAPI.GetPages: %w", err)
	}
	return page, nil
}

func New(w PageNotionWebAPI) *PageUseCase {
	return &PageUseCase{
		w,
	}
}
