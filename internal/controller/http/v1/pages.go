package v1

import (
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"ntheanh201-journal/internal/entity"
	"ntheanh201-journal/internal/usecase"
)

type pageRoutes struct {
	p usecase.Page
	l logger.Interface
}

func NewPageRoutes(handler *gin.RouterGroup, p usecase.Page, l logger.Interface) {
	r := &pageRoutes{p, l}
	h := handler.Group("/pages")
	{
		h.GET("", r.getPages)
		h.GET("/:pageId", r.getPage)
		h.GET("/:pageId/children", r.getBlockChildren)
	}
}

func (r *pageRoutes) getPages(ctx *gin.Context) {
	res, err := r.p.GetPages(ctx.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - pages")
		errorResponse(ctx, http.StatusInternalServerError, "fetching pages problems")
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (r *pageRoutes) getPage(ctx *gin.Context) {
	page, err := r.p.GetPage(ctx.Request.Context(), entity.ObjectID(ctx.Param("pageId")))
	if err != nil {
		r.l.Error(err, "http - v1 - page")
		errorResponse(ctx, http.StatusInternalServerError, "fetching page problems")
		return
	}
	ctx.JSON(http.StatusOK, page)
}

func (r *pageRoutes) getBlockChildren(ctx *gin.Context) {
	res, err := r.p.GetBlockChildren(ctx.Request.Context(), entity.ObjectID(ctx.Param("pageId")))
	if err != nil {
		r.l.Error(err, "http - v1 - block children")
		errorResponse(ctx, http.StatusInternalServerError, "fetching block children problems")
		return
	}
	ctx.JSON(http.StatusOK, res)
}
