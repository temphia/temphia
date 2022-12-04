package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (s *Server) dataAPI(rg *gin.RouterGroup) {

	rg.GET("/", s.dx(s.loadGroup))

	rg.POST("/:tid/row", s.dx(s.newRow))
	rg.GET("/:tid/row/:id", s.dx(s.getRow))
	rg.POST("/:tid/row/:id", s.dx(s.updateRow))
	rg.DELETE("/:tid/row/:id", s.dx(s.deleteRow))
	rg.POST("/:tid/simple_query", s.dx(s.simpleQuery))
	rg.POST("/:tid/fts_query", s.dx(s.FTSQuery)) // fixme => remove this and consolidate this to simple_query ?
	rg.POST("/:tid/ref_load", s.dx(s.refLoad))
	rg.POST("/:tid/ref_resolve", s.dx(s.refResolve))
	rg.POST("/:tid/rev_ref_load", s.dx(s.reverseRefLoad))
	rg.GET("/:tid/activity/:row_id", s.dx(s.listActivity))
	rg.POST("/:tid/activity/:row_id", s.dx(s.commentRow))
}

func (s *Server) dataWSAPI(rg *gin.RouterGroup) {
	rg.GET("/", s.sockdDataWS)
	rg.POST("/", s.sockdDataUpdateWS)
}

func (s *Server) loadGroup(uclaim *claim.Data, ctx *gin.Context) {
	gr, err := s.cData.LoadGroup(uclaim)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteJSON(ctx, gr, err)
}

func (s *Server) newRow(uclaim *claim.Data, ctx *gin.Context) {

	data := &newRowReq{}
	err := ctx.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	id, err := s.cData.NewRow(uclaim, ctx.Param("tid"), data.Cells)
	httpx.WriteJSON(ctx, id, err)
}

func (s *Server) getRow(uclaim *claim.Data, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {

		httpx.WriteErr(ctx, err)
		return
	}
	cells, err := s.cData.GetRow(uclaim, ctx.Param("tid"), id)
	httpx.WriteJSON(ctx, cells, err)
}

type updateRowReq struct {
	Version int64                  `json:"version,omitempty"`
	Cells   map[string]interface{} `json:"cells,omitempty"`
}

func (s *Server) updateRow(uclaim *claim.Data, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	data := &updateRowReq{}
	err = ctx.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	cells, err := s.cData.UpdateRow(uclaim, ctx.Param("tid"), id, data.Version, data.Cells)
	httpx.WriteJSON(ctx, cells, err)
}

func (s *Server) deleteRow(uclaim *claim.Data, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}
	err = s.cData.DeleteRow(uclaim, ctx.Param("tid"), id)
	httpx.WriteFinal(ctx, err)
}

func (s *Server) simpleQuery(uclaim *claim.Data, ctx *gin.Context) {
	query := store.SimpleQueryReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.SimpleQuery(uclaim, ctx.Param("tid"), query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) FTSQuery(uclaim *claim.Data, ctx *gin.Context) {
	query := store.FTSQueryReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.FTSQuery(uclaim, ctx.Param("tid"), query.SearchTerm)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) refLoad(uclaim *claim.Data, ctx *gin.Context) {
	query := &store.RefLoadReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.RefLoad(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) refResolve(uclaim *claim.Data, ctx *gin.Context) {
	query := &store.RefResolveReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.RefResolve(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) reverseRefLoad(uclaim *claim.Data, ctx *gin.Context) {
	query := &store.RevRefLoadReq{}

	err := ctx.BindJSON(query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ReverseRefLoad(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) listActivity(uclaim *claim.Data, ctx *gin.Context) {
	rid, err := strconv.ParseInt(ctx.Param("row_id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ListActivity(uclaim, ctx.Param("tid"), int(rid))
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) commentRow(uclaim *claim.Data, ctx *gin.Context) {
	rid, err := strconv.ParseInt(ctx.Param("row_id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	reqdata := commentRowReq{}

	err = ctx.BindJSON(&reqdata)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cData.CommentRow(uclaim, ctx.Param("tid"), reqdata.Message, int(rid))
	httpx.WriteFinal(ctx, err)
}

func (s *Server) dx(fn func(uclaim *claim.Data, ctx *gin.Context)) func(*gin.Context) {
	return s.middleware.DataX(fn)
}

// data sockd

func (s *Server) sockdDataWS(ctx *gin.Context) {

}

func (s *Server) sockdDataUpdateWS(ctx *gin.Context) {

}

// models

type newRowReq struct {
	Cells map[string]interface{} `json:"cells,omitempty"`
}

type commentRowReq struct {
	Message string `json:"message,omitempty"`
}
