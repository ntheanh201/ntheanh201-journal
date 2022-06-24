package usecase

import (
	"context"
	"fmt"
	"ntheanh201-journal/internal/entity"
)

type PageUseCase struct {
	webAPI PageNotionWebAPI
}

func (p *PageUseCase) GetPages(ctx context.Context) ([]entity.Page, error) {
	pages, err := p.webAPI.GetPages()
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}
	return pages, nil
}

func (p *PageUseCase) GetPage(ctx context.Context, id entity.ObjectID) (entity.Page, error) {
	page, err := p.webAPI.GetPage(id)
	if err != nil {
		return entity.Page{}, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}
	return page, nil
}

func New(w PageNotionWebAPI) *PageUseCase {
	return &PageUseCase{
		w,
	}
}
