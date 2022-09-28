package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (s *Server) dataAPI(rg *gin.RouterGroup) {

	rg.GET("/", s.DataX(s.loadGroup))

	rg.POST("/:tid/row", s.DataX(s.newRow))
	rg.GET("/:tid/row/:id", s.DataX(s.getRow))
	rg.POST("/:tid/row/:id", s.DataX(s.updateRow))
	rg.DELETE("/:tid/row/:id", s.DataX(s.deleteRow))
	rg.POST("/:tid/simple_query", s.DataX(s.simpleQuery))
	rg.POST("/:tid/fts_query", s.DataX(s.FTSQuery)) // fixme => remove this and consolidate this to simple_query ?
	rg.POST("/:tid/ref_load", s.DataX(s.refLoad))
	rg.POST("/:tid/ref_resolve", s.DataX(s.refResolve))
	rg.POST("/:tid/rev_ref_load", s.DataX(s.reverseRefLoad))
	rg.GET("/:tid/activity/:row_id", s.DataX(s.listActivity))
	rg.POST("/:tid/activity/:row_id", s.DataX(s.commentRow))
}

type newRowReq struct {
	Cells map[string]interface{} `json:"cells,omitempty"`
}

func (s *Server) loadGroup(uclaim *claim.DataTkt, ctx *gin.Context) {
	gr, err := s.cData.LoadGroup(uclaim)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteJSON(ctx, gr, err)
}

func (s *Server) newRow(uclaim *claim.DataTkt, ctx *gin.Context) {

	data := &newRowReq{}
	err := ctx.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	id, err := s.cData.NewRow(uclaim, ctx.Param("tid"), data.Cells)
	httpx.WriteJSON(ctx, id, err)
}

func (s *Server) getRow(uclaim *claim.DataTkt, ctx *gin.Context) {
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

func (s *Server) updateRow(uclaim *claim.DataTkt, ctx *gin.Context) {
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

func (s *Server) deleteRow(uclaim *claim.DataTkt, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}
	err = s.cData.DeleteRow(uclaim, ctx.Param("tid"), id)
	httpx.WriteFinal(ctx, err)
}

func (s *Server) simpleQuery(uclaim *claim.DataTkt, ctx *gin.Context) {
	query := store.SimpleQueryReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.SimpleQuery(uclaim, ctx.Param("tid"), query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) FTSQuery(uclaim *claim.DataTkt, ctx *gin.Context) {
	query := store.FTSQueryReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.FTSQuery(uclaim, ctx.Param("tid"), query.SearchTerm)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) refLoad(uclaim *claim.DataTkt, ctx *gin.Context) {
	query := &store.RefLoadReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.RefLoad(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) refResolve(uclaim *claim.DataTkt, ctx *gin.Context) {
	query := &store.RefResolveReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.RefResolve(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) reverseRefLoad(uclaim *claim.DataTkt, ctx *gin.Context) {
	query := &store.RevRefLoadReq{}

	err := ctx.BindJSON(query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ReverseRefLoad(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) listActivity(uclaim *claim.DataTkt, ctx *gin.Context) {
	rid, err := strconv.ParseInt(ctx.Param("row_id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ListActivity(uclaim, ctx.Param("tid"), int(rid))
	httpx.WriteJSON(ctx, resp, err)
}

type commentRowReq struct {
	Message string `json:"message,omitempty"`
}

func (s *Server) commentRow(uclaim *claim.DataTkt, ctx *gin.Context) {
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

func (s *Server) DataX(fn func(uclaim *claim.DataTkt, ctx *gin.Context)) func(*gin.Context) {

	return nil
}

// private
