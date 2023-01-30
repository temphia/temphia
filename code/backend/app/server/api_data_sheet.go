package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/data"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (s *Server) dataSheetAPI(rg *gin.RouterGroup) {
	rg.POST("/sheet/list", s.dx(s.listSheetGroup))
	rg.POST("/sheet/:id/load", s.dx(s.loadSheet))

	rg.GET("/sheet", s.dx(s.listSheet))
	rg.GET("/sheet/:id", s.dx(s.getSheet))
	rg.POST("/sheet", s.dx(s.newSheet))
	rg.POST("/sheet/:id", s.dx(s.updateSheet))
	rg.DELETE("/sheet/:id", s.dx(s.deleteSheet))

	rg.GET("/sheet/:id/column", s.dx(s.listSheetColumn))
	rg.POST("/sheet/:id/column", s.dx(s.newSheetColumn))
	rg.GET("/sheet/:id/column/:cid", s.dx(s.getSheetColumn))
	rg.POST("/sheet/:id/column/:cid", s.dx(s.updateSheetColumn))
	rg.DELETE("/sheet/:id/column/:cid", s.dx(s.deleteSheetColumn))

	rg.POST("/sheet/:id/row_cell", s.dx(s.NewRowWithCell))
	rg.POST("/sheet/:id/row_cell/:rid", s.dx(s.UpdateRowWithCell))
	rg.GET("/sheet/:id/row_cell/:rid", s.dx(s.GetRowWithCell))
	rg.DELETE("/sheet/:id/row_cell/:rid", s.dx(s.DeleteRowWithCell))

}

func (s *Server) listSheetGroup(uclaim *claim.Data, ctx *gin.Context) {
	resp, err := s.cData.ListSheetGroup(uclaim)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) loadSheet(uclaim *claim.Data, ctx *gin.Context) {
	data := data.LoadSheetReq{}
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

}

func (s *Server) getSheet(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) updateSheet(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) deleteSheet(uclaim *claim.Data, ctx *gin.Context) {

}

// columns

func (s *Server) listSheetColumn(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) newSheetColumn(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) getSheetColumn(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) updateSheetColumn(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) deleteSheetColumn(uclaim *claim.Data, ctx *gin.Context) {

}

// cells

func (s *Server) NewRowWithCell(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make([]map[string]any, 0)

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
	resp, err := s.cData.UpdateRowWithCell(uclaim, sid, rid, data)
	httpx.WriteJSON(ctx, resp, err)

}

func (s *Server) GetRowWithCell(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Server) DeleteRowWithCell(uclaim *claim.Data, ctx *gin.Context) {

}
