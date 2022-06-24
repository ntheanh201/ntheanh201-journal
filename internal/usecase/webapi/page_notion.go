package webapi

import (
	"ntheanh201-journal/internal/entity"
)

type PageNotionWebAPI struct {
}

func New() *PageNotionWebAPI {
	return &PageNotionWebAPI{}
}

func (p *PageNotionWebAPI) GetPages() ([]entity.Page, error) {
	// TODO: implement me
	return []entity.Page{{}}, nil
}

func (p *PageNotionWebAPI) GetPage(id entity.ObjectID) (entity.Page, error) {
	// TODO: implement me
	return entity.Page{}, nil
}
