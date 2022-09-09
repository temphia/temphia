package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (s *Server) dataAPI(rg *gin.RouterGroup) {

	rg.GET("/load", s.X(s.loadGroup))

	rg.POST("/:table_id/row", s.X(s.newRow))
	rg.GET("/:table_id/row/:id", s.X(s.getRow))
	rg.POST("/:table_id/row/:id", s.X(s.updateRow))
	rg.DELETE("/:table_id/row/:id", s.X(s.deleteRow))
	rg.POST("/:table_id/simple_query", s.X(s.simpleQuery))
	rg.POST("/:table_id/fts_query", s.X(s.FTSQuery)) // fixme => remove this and consolidate this to simple_query ?
	rg.POST("/:table_id/ref_load", s.X(s.refLoad))
	rg.POST("/:table_id/ref_resolve", s.X(s.refResolve))
	rg.POST("/:table_id/rev_ref_load", s.X(s.reverseRefLoad))
	rg.GET("/:table_id/activity/:row_id", s.X(s.listActivity))
	rg.POST("/:table_id/activity/:row_id", s.X(s.commentRow))
}

type newRowReq struct {
	Cells map[string]interface{} `json:"cells,omitempty"`
}

func (s *Server) loadGroup(ctx httpx.Request) {
	gr, err := s.cData.LoadGroup(ctx.Session)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	httpx.WriteJSON(ctx.Http, gr, err)
}

func (s *Server) newRow(ctx httpx.Request) {

	data := &newRowReq{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := s.cData.NewRow(ctx.Session, ctx.MustParam("table_id"), data.Cells)
	httpx.WriteJSON(ctx.Http, id, err)
}

func (s *Server) getRow(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {

		httpx.WriteErr(ctx.Http, err.Error())
		return
	}
	cells, err := s.cData.GetRow(ctx.Session, ctx.MustParam("table_id"), id)
	httpx.WriteJSON(ctx.Http, cells, err)
}

type updateRowReq struct {
	Version int64                  `json:"version,omitempty"`
	Cells   map[string]interface{} `json:"cells,omitempty"`
}

func (s *Server) updateRow(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	data := &updateRowReq{}
	err = ctx.Http.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	cells, err := s.cData.UpdateRow(ctx.Session, ctx.MustParam("table_id"), id, data.Version, data.Cells)
	httpx.WriteJSON(ctx.Http, cells, err)
}

func (s *Server) deleteRow(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}
	err = s.cData.DeleteRow(ctx.Session, ctx.MustParam("table_id"), id)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) simpleQuery(ctx httpx.Request) {
	query := store.SimpleQueryReq{}
	err := ctx.Http.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := s.cData.SimpleQuery(ctx.Session, ctx.MustParam("table_id"), query)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) FTSQuery(ctx httpx.Request) {
	query := store.FTSQueryReq{}
	err := ctx.Http.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := s.cData.FTSQuery(ctx.Session, ctx.MustParam("table_id"), query.SearchTerm)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) refLoad(ctx httpx.Request) {
	query := &store.RefLoadReq{}
	err := ctx.Http.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := s.cData.RefLoad(ctx.Session, query)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) refResolve(ctx httpx.Request) {
	query := &store.RefResolveReq{}
	err := ctx.Http.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := s.cData.RefResolve(ctx.Session, query)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) reverseRefLoad(ctx httpx.Request) {
	query := &store.RevRefLoadReq{}

	err := ctx.Http.BindJSON(query)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := s.cData.ReverseRefLoad(ctx.Session, query)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) listActivity(ctx httpx.Request) {
	rid, err := strconv.ParseInt(ctx.MustParam("row_id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := s.cData.ListActivity(ctx.Session, ctx.MustParam("table_id"), int(rid))
	httpx.WriteJSON(ctx.Http, resp, err)
}

type commentRowReq struct {
	Message string `json:"message,omitempty"`
}

func (s *Server) commentRow(ctx httpx.Request) {
	rid, err := strconv.ParseInt(ctx.MustParam("row_id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	reqdata := commentRowReq{}

	err = ctx.Http.BindJSON(&reqdata)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = s.cData.CommentRow(ctx.Session, ctx.MustParam("table_id"), reqdata.Message, int(rid))
	httpx.WriteFinal(ctx.Http, err)
}
