package data

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (c *Controller) ListSheetGroup(uclaim *claim.Data) (*dyndb.ListSheetGroupResp, error) {

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	tg, err := ddb.GetGroup(uclaim.TenantId, uclaim.DataGroup)
	if err != nil {
		return nil, err
	}

	resp, err := thub.ListSheetGroup(0)
	if err != nil {
		return nil, err
	}

	ftok, err := c.folderTicket(tg, uclaim)
	if err != nil {
		return nil, err
	}

	resp.FolderTicket = ftok
	return resp, nil

}

func (c *Controller) FTSQuerySheet(uclaim *claim.Data, req *dyndb.FTSQuerySheet) (*dyndb.QuerySheetResp, error) {

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	req.Group = uclaim.DataGroup
	req.TenantId = uclaim.TenantId

	return thub.FTSQuery(0, req)
}

func (c *Controller) LoadSheet(uclaim *claim.Data, data *dyndb.LoadSheetReq) (*dyndb.LoadSheetResp, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.LoadSheet(0, data)
}

func (c *Controller) QuerySheet(uclaim *claim.Data, data *dyndb.QuerySheetReq) (*dyndb.QuerySheetResp, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.Query(0, &dyndb.QuerySheetReq{
		TenantId:    uclaim.TenantId,
		Group:       uclaim.DataGroup,
		SheetId:     data.SheetId,
		View:        data.View,
		FilterConds: data.FilterConds,
		Desc:        data.Desc,
		RowCursorId: data.RowCursorId,
	})
}

func (c *Controller) RefQuery(uclaim *claim.Data, data *dyndb.RefQuerySheet) (*dyndb.QuerySheetResp, error) {

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.RefQuery(0, &dyndb.RefQuerySheet{
		TenantId:      uclaim.TenantId,
		Group:         uclaim.DataGroup,
		SheetId:       data.SheetId,
		ColumnId:      data.ColumnId,
		RowCursorId:   data.RowCursorId,
		TargetSource:  data.TargetSource,
		TargetGroup:   data.TargetGroup,
		TargetSheetId: data.TargetSheetId,
	})

}

// sheets

func (c *Controller) ListSheet(uclaim *claim.Data) ([]map[string]any, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.ListSheet(0)
}

func (c *Controller) NewSheet(uclaim *claim.Data, data map[string]any) error {

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.NewSheet(0, uclaim.UserID, data)

}

func (c *Controller) GetSheet(uclaim *claim.Data, id int64) (map[string]any, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.GetSheet(0, id)
}

func (c *Controller) UpdateSheet(uclaim *claim.Data, id int64, data map[string]any) error {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.UpdateSheet(0, id, uclaim.UserID, data)

}

func (c *Controller) DeleteSheet(uclaim *claim.Data, id int64) error {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.DeleteSheet(0, id, uclaim.UserID)
}

// columns

func (c *Controller) ListSheetColumn(uclaim *claim.Data, sid int64) ([]map[string]any, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.ListSheetColumn(0, sid)

}

func (c *Controller) NewSheetColumn(uclaim *claim.Data, sid int64, data map[string]any) (int64, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.NewSheetColumn(0, sid, uclaim.UserID, data)
}

func (c *Controller) GetSheetColumn(uclaim *claim.Data, sid, cid int64) (map[string]any, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.GetSheetColumn(0, sid, cid)
}

func (c *Controller) UpdateSheetColumn(uclaim *claim.Data, sid, cid int64, data map[string]any) error {

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.UpdateSheetColumn(0, sid, cid, uclaim.UserID, data)
}

func (c *Controller) DeleteSheetColumn(uclaim *claim.Data, sid, cid int64) error {

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.DeleteSheetColumn(0, sid, cid, uclaim.UserID)

}

// row_cells

func (c *Controller) NewRowWithCell(uclaim *claim.Data, sid int64, data map[int64]map[string]any) (any, error) {

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.NewRowWithCell(0, sid, uclaim.UserID, data)
}

func (c *Controller) UpdateRowWithCell(uclaim *claim.Data, sid, rid int64, data map[int64]map[string]any) (map[int64]map[string]any, error) {

	pp.Println("@update", data)

	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.UpdateRowWithCell(0, sid, rid, uclaim.UserID, data)
}

func (c *Controller) DeleteRowWithCell(uclaim *claim.Data, sid, rid int64) error {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)
	return thub.DeleteRowWithCell(0, sid, rid, uclaim.UserID)
}

func (c *Controller) GetRowRelations(uclaim *claim.Data, sid, rid, refsheet, refcol int64) (*dyndb.Relation, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.GetRowRelations(0, sid, rid, refsheet, refcol)
}

func (c *Controller) GetRowHistory(uclaim *claim.Data, sid, rid int64) ([]*entities.DynActivity, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.GetRowHistory(rid)
}

func (c *Controller) ExportSheets(uclaim *claim.Data, sheets []int64) (*dyndb.ExportData, error) {
	source, _ := getTarget(uclaim)
	ddb := c.dynHub.GetSource(source, uclaim.TenantId)
	thub := ddb.GetDataSheetHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.ExportSheets(0, dyndb.ExportOptions{
		TenantId: uclaim.TenantId,
		Group:    uclaim.DataGroup,
		Sheets:   sheets,
	})
}
