package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (s *Server) dataSheetAPI(rg *gin.RouterGroup) {
	rg.POST("/export", s.dx(s.export))
	rg.POST("/list", s.dx(s.listSheetGroup))
	rg.POST("/:id/load", s.dx(s.loadSheet))
	rg.POST("/:id/search", s.dx(s.searchSheet))
	rg.POST("/:id/query", s.dx(s.querySheet))
	rg.POST("/:id/ref_query", s.dx(s.refSheet))

	rg.GET("/", s.dx(s.listSheet))
	rg.GET("/:id", s.dx(s.getSheet))
	rg.POST("/", s.dx(s.newSheet))
	rg.POST("/:id", s.dx(s.updateSheet))
	rg.DELETE("/:id", s.dx(s.deleteSheet))

	rg.GET("/:id/column", s.dx(s.listSheetColumn))
	rg.POST("/:id/column", s.dx(s.newSheetColumn))
	rg.GET("/:id/column/:cid", s.dx(s.getSheetColumn))
	rg.POST("/:id/column/:cid", s.dx(s.updateSheetColumn))
	rg.DELETE("/:id/column/:cid", s.dx(s.deleteSheetColumn))

	rg.POST("/:id/row_cell", s.dx(s.NewRowWithCell))
	rg.POST("/:id/row_cell/:rid", s.dx(s.UpdateRowWithCell))
	rg.GET("/:id/row_cell/:rid", s.dx(s.GetRowWithCell))
	rg.DELETE("/:id/row_cell/:rid", s.dx(s.DeleteRowWithCell))
	rg.GET("/:id/relation/:rid/ref/:refsheet/column/:refcol", s.dx(s.GetRowRelations))

	rg.GET("/:id/history/:rid", s.dx(s.getRowHistory))

}

func (s *Server) export(uclaim *claim.Data, ctx *gin.Context) {

	sheets := make([]int64, 0)

	err := ctx.BindJSON(&sheets)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ExportSheets(uclaim, sheets)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) listSheetGroup(uclaim *claim.Data, ctx *gin.Context) {
	resp, err := s.cData.ListSheetGroup(uclaim)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) searchSheet(uclaim *claim.Data, ctx *gin.Context) {
	data := dyndb.FTSQuerySheet{}
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	data.SheetId = id

	resp, err := s.cData.FTSQuerySheet(uclaim, &data)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) refSheet(uclaim *claim.Data, ctx *gin.Context) {
	data := dyndb.RefQuerySheet{}
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	data.SheetId = id

	resp, err := s.cData.RefQuery(uclaim, &data)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) querySheet(uclaim *claim.Data, ctx *gin.Context) {
	data := dyndb.QuerySheetReq{}
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	data.SheetId = id

	resp, err := s.cData.QuerySheet(uclaim, &data)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) loadSheet(uclaim *claim.Data, ctx *gin.Context) {
	data := dyndb.LoadSheetReq{}
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	data.SheetId = id

	resp, err := s.cData.LoadSheet(uclaim, &data)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) listSheet(uclaim *claim.Data, ctx *gin.Context) {
	resp, err := s.cData.ListSheet(uclaim)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) newSheet(uclaim *claim.Data, ctx *gin.Context) {
	data := make(map[string]any, 0)
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cData.NewSheet(uclaim, data)
	httpx.WriteJSON(ctx, nil, err)
}

func (s *Server) getSheet(uclaim *claim.Data, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	resp, err := s.cData.GetSheet(uclaim, id)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) updateSheet(uclaim *claim.Data, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[string]any, 0)
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cData.UpdateSheet(uclaim, id, data)
	httpx.WriteJSON(ctx, nil, err)

}

func (s *Server) deleteSheet(uclaim *claim.Data, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := s.cData.DeleteSheet(uclaim, id)
	httpx.WriteJSON(ctx, nil, err)

}

// columns

func (s *Server) listSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	resp, err := s.cData.ListSheetColumn(uclaim, sid)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) newSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[string]any, 0)
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.NewSheetColumn(uclaim, sid, data)
	httpx.WriteJSON(ctx, resp, err)

}

func (s *Server) getSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(ctx.Param("cid"), 10, 64)

	resp, err := s.cData.GetSheetColumn(uclaim, sid, cid)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) updateSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(ctx.Param("cid"), 10, 64)

	data := make(map[string]any, 0)
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cData.UpdateSheetColumn(uclaim, sid, cid, data)
	httpx.WriteJSON(ctx, nil, err)
}

func (s *Server) deleteSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(ctx.Param("cid"), 10, 64)

	err := s.cData.DeleteSheetColumn(uclaim, sid, cid)
	httpx.WriteJSON(ctx, nil, err)
}

// cells

func (s *Server) NewRowWithCell(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[int64]map[string]any, 0)

	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.NewRowWithCell(uclaim, sid, data)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) UpdateRowWithCell(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	rid, _ := strconv.ParseInt(ctx.Param("rid"), 10, 64)

	data := make(map[int64]map[string]any, 0)
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.UpdateRowWithCell(uclaim, sid, rid, data)
	httpx.WriteJSON(ctx, resp, err)

}

func (s *Server) GetRowWithCell(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) DeleteRowWithCell(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	rid, _ := strconv.ParseInt(ctx.Param("rid"), 10, 64)

	err := s.cData.DeleteRowWithCell(uclaim, sid, rid)
	httpx.WriteJSON(ctx, nil, err)
}

func (s *Server) GetRowRelations(uclaim *claim.Data, ctx *gin.Context) {

	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	rid, _ := strconv.ParseInt(ctx.Param("rid"), 10, 64)

	refsheet, _ := strconv.ParseInt(ctx.Param("refsheet"), 10, 64)
	refcol, _ := strconv.ParseInt(ctx.Param("refcol"), 10, 64)

	resp, err := s.cData.GetRowRelations(uclaim, sid, rid, refsheet, refcol)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) getRowHistory(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	rid, _ := strconv.ParseInt(ctx.Param("rid"), 10, 64)

	resp, err := s.cData.GetRowHistory(uclaim, sid, rid)
	httpx.WriteJSON(ctx, resp, err)
}
