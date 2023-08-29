package apidata

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Data) export(uclaim *claim.Data, ctx *gin.Context) {

	sheets := make([]int64, 0)

	err := ctx.BindJSON(&sheets)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := s.cData.ExportSheets(uclaim, sheets)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) listSheetGroup(uclaim *claim.Data, ctx *gin.Context) {
	resp, err := s.cData.ListSheetGroup(uclaim)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) searchSheet(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) refSheet(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) querySheet(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) loadSheet(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) listSheet(uclaim *claim.Data, ctx *gin.Context) {
	resp, err := s.cData.ListSheet(uclaim)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) newSheet(uclaim *claim.Data, ctx *gin.Context) {
	data := make(map[string]any, 0)
	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cData.NewSheet(uclaim, data)
	httpx.WriteJSON(ctx, nil, err)
}

func (s *Data) getSheet(uclaim *claim.Data, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	resp, err := s.cData.GetSheet(uclaim, id)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) updateSheet(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) deleteSheet(uclaim *claim.Data, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := s.cData.DeleteSheet(uclaim, id)
	httpx.WriteJSON(ctx, nil, err)

}

// columns

func (s *Data) listSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	resp, err := s.cData.ListSheetColumn(uclaim, sid)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) newSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) getSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(ctx.Param("cid"), 10, 64)

	resp, err := s.cData.GetSheetColumn(uclaim, sid, cid)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) updateSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) deleteSheetColumn(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(ctx.Param("cid"), 10, 64)

	err := s.cData.DeleteSheetColumn(uclaim, sid, cid)
	httpx.WriteJSON(ctx, nil, err)
}

// cells

func (s *Data) NewRowWithCell(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) UpdateRowWithCell(uclaim *claim.Data, ctx *gin.Context) {
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

func (s *Data) GetRowWithCell(uclaim *claim.Data, ctx *gin.Context) {

}

func (s *Data) DeleteRowWithCell(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	rid, _ := strconv.ParseInt(ctx.Param("rid"), 10, 64)

	err := s.cData.DeleteRowWithCell(uclaim, sid, rid)
	httpx.WriteJSON(ctx, nil, err)
}

func (s *Data) GetRowRelations(uclaim *claim.Data, ctx *gin.Context) {

	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	rid, _ := strconv.ParseInt(ctx.Param("rid"), 10, 64)

	refsheet, _ := strconv.ParseInt(ctx.Param("refsheet"), 10, 64)
	refcol, _ := strconv.ParseInt(ctx.Param("refcol"), 10, 64)

	resp, err := s.cData.GetRowRelations(uclaim, sid, rid, refsheet, refcol)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Data) getRowHistory(uclaim *claim.Data, ctx *gin.Context) {
	sid, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	rid, _ := strconv.ParseInt(ctx.Param("rid"), 10, 64)

	resp, err := s.cData.GetRowHistory(uclaim, sid, rid)
	httpx.WriteJSON(ctx, resp, err)
}
