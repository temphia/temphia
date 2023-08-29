package apidata

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Data) loadGroup(uclaim *claim.Data, ctx *gin.Context) {
	gr, err := s.cData.LoadGroup(uclaim)
	httpx.WriteJSON(ctx, gr, err)
}

func (s *Data) newRow(uclaim *claim.Data, ctx *gin.Context) {

	data := &newRowReq{}
	err := ctx.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	id, err := s.cData.NewRow(uclaim, ctx.Param("tid"), data.Cells)
	httpx.WriteJSON(ctx, id, err)
}

func (s *Data) getRow(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) updateRow(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) deleteRow(uclaim *claim.Data, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}
	err = s.cData.DeleteRow(uclaim, ctx.Param("tid"), id)
	httpx.WriteFinal(ctx, err)
}

func (s *Data) loadTable(uclaim *claim.Data, ctx *gin.Context) {
	req := dyndb.LoadTableReq{}
	err := ctx.BindJSON(&req)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.LoadTable(uclaim, req, ctx.Param("tid"))
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) simpleQuery(uclaim *claim.Data, ctx *gin.Context) {
	query := dyndb.SimpleQueryReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.SimpleQuery(uclaim, ctx.Param("tid"), query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) FTSQuery(uclaim *claim.Data, ctx *gin.Context) {
	query := dyndb.FTSQueryReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	query.Table = ctx.Param("tid")

	resp, err := s.cData.FTSQuery(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) refLoad(uclaim *claim.Data, ctx *gin.Context) {
	query := &dyndb.RefLoadReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.RefLoad(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) refResolve(uclaim *claim.Data, ctx *gin.Context) {
	query := &dyndb.RefResolveReq{}
	err := ctx.BindJSON(&query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.RefResolve(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) reverseRefLoad(uclaim *claim.Data, ctx *gin.Context) {
	query := &dyndb.RevRefLoadReq{}

	err := ctx.BindJSON(query)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ReverseRefLoad(uclaim, query)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) listActivity(uclaim *claim.Data, ctx *gin.Context) {
	rid, err := strconv.ParseInt(ctx.Param("row_id"), 10, 64)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ListActivity(uclaim, ctx.Param("tid"), int(rid))
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) commentRow(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) dx(fn func(uclaim *claim.Data, ctx *gin.Context)) func(*gin.Context) {
	return s.middleware.DataX(fn)
}

// models

type newRowReq struct {
	Cells map[string]interface{} `json:"cells,omitempty"`
	// Assets map[string]string      `json:"assets,omitempty"`
}

type commentRowReq struct {
	Message string `json:"message,omitempty"`
}

// utils

type DataUserReq struct {
	TargetType string `json:"target_type,omitempty"` // sheet | data
	Target     string `json:"target,omitempty"`      // column id/slug
}

func (s *Data) listDataUsers(uclaim *claim.Data, ctx *gin.Context) {
	req := DataUserReq{}
	err := ctx.BindJSON(&req)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ListDataUsers(uclaim, req.TargetType, req.Target)
	httpx.WriteJSON(ctx, resp, err)
}
