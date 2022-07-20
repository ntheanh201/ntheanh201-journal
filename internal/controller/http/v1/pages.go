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
		h.GET("/blocks/:slug", r.getBlockChildrenBySlug)
		h.GET("/record/:slug", r.getPageChunkV3)
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

func (r *pageRoutes) getBlockChildrenBySlug(ctx *gin.Context) {
	res, err := r.p.GetBlockChildren(ctx.Request.Context(), entity.ObjectID(ctx.Param("slug")))
	if err != nil {
		r.l.Error(err, "http - v1 - block children by slug")
		errorResponse(ctx, http.StatusNotFound, "fetching block children by slug problems")
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (r *pageRoutes) getPageChunkV3(ctx *gin.Context) {
	res, err := r.p.LoadPageChunkV3(ctx.Request.Context(), ctx.Param("slug"))
	if err != nil {
		r.l.Error(err, "http - v1 - load page chunk")
		errorResponse(ctx, http.StatusNotFound, "fetching page chunk problems")
		return
	}
	ctx.JSON(http.StatusOK, res)
}
