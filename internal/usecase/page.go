package usecase

import (
	"context"
	"errors"
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
		return entity.Page{}, fmt.Errorf("PageUseCase - webAPI.GetPage: %w", err)
	}
	return page, nil
}

func (uc *PageUseCase) GetBlockChildren(ctx context.Context, id entity.ObjectID) (response.BlockChildrenResponse, error) {
	pageSize := 100
	res, err := uc.webAPI.GetNotionBlockChildren(ctx, id, pageSize)
	if err != nil {
		return response.BlockChildrenResponse{}, fmt.Errorf("PageUseCase - webAPI.GetBlockChildren: %w", err)
	}
	return res, nil
}

func (uc *PageUseCase) GetBlockChildrenBySlug(ctx context.Context, slug string) (response.BlockChildrenResponse, error) {
	pages, err := uc.webAPI.QueryNotionPageBySlug(ctx, slug)

	if err != nil || len(pages.Results) == 0 {
		_ = fmt.Errorf("PageUseCase - webAPI.QueryNotionPageBySlug: %w", err)
		return response.BlockChildrenResponse{}, errors.New("cannot query page")
	}

	page := pages.Results[0]
	id := page.ID

	pageSize := 100
	res, err := uc.webAPI.GetNotionBlockChildren(ctx, id, pageSize)
	if err != nil {
		return response.BlockChildrenResponse{}, fmt.Errorf("PageUseCase - webAPI.GetBlockChildren: %w", err)
	}
	return res, nil
}

func (uc *PageUseCase) LoadPageChunkV3(ctx context.Context, slug string) (response.LoadPageChunkResponse, error) {
	pages, err := uc.webAPI.QueryNotionPageBySlug(ctx, slug)

	if err != nil || len(pages.Results) == 0 {
		_ = fmt.Errorf("PageUseCase - webAPI.QueryNotionPageBySlug: %w", err)
		return response.LoadPageChunkResponse{}, errors.New("cannot load page chunk")
	}

	page := pages.Results[0]
	id := page.ID

	res, err := uc.webAPI.LoadPageChunkV3(ctx, id)
	if err != nil {
		return response.LoadPageChunkResponse{}, fmt.Errorf("PageUseCase - webAPI.LoadPageChunkV3: %w", err)
	}
	return res, nil
}

func New(w PageNotionWebAPI) *PageUseCase {
	return &PageUseCase{
		w,
	}
}
