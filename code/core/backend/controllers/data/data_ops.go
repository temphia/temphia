package data

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (c *Controller) LoadGroup(uclaim *claim.DataTkt) (*store.LoadDgroupResp, error) {
	dynDB := c.dynHub.GetSource(uclaim.DataSource, uclaim.TenentId)

	tg, err := dynDB.GetGroup(uclaim.DataGroup)
	if err != nil {
		return nil, err
	}

	tables, err := dynDB.ListTables(uclaim.DataGroup)

	if err != nil {
		return nil, err
	}

	if tg.CabinetSource == "" || tg.CabinetFolder == "" {
		tg.CabinetSource = c.cabHub.DefaultName(uclaim.TenentId)
		tg.CabinetFolder = "data_common"
	}

	fcalim := &claim.FolderTkt{
		Folder: tg.CabinetFolder,
		Source: tg.CabinetSource,
		Expiry: 0,
		Prefix: "",
		//	DeviceId: uclaim.DeviceId,
	}

	cabToken, err := c.signer.SignFolderTkt(uclaim.TenentId, fcalim)
	if err != nil {
		return nil, err
	}

	resp := &store.LoadDgroupResp{
		Tables:          tables,
		CabinetTicket:   cabToken,
		SockdRoomTicket: "",
	}

	return resp, nil
}

func (d *Controller) NewRow(uclaim *claim.DataTkt, tslug string, cells map[string]any) (int64, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)

	return dynDb.NewRow(0, store.NewRowReq{
		TenantId: uclaim.TenentId,
		Group:    group,
		Table:    tslug,
		Data:     cells,
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (d *Controller) GetRow(uclaim *claim.DataTkt, tslug string, id int64) (map[string]any, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)

	return dynDb.GetRow(0, store.GetRowReq{
		TenantId: uclaim.TenentId,
		Group:    group,
		Table:    tslug,
		Id:       id,
	})
}

func (d *Controller) UpdateRow(uclaim *claim.DataTkt, tslug string, id, version int64, cells map[string]any) (map[string]any, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.UpdateRow(0, store.UpdateRowReq{
		TenantId: uclaim.TenentId,
		Version:  version,
		Group:    group,
		Table:    tslug,
		Data:     cells,
		Id:       id,
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (d *Controller) DeleteRow(uclaim *claim.DataTkt, tslug string, id int64) error {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.DeleteRows(0, store.DeleteRowReq{
		TenantId: uclaim.TenentId,
		Group:    group,
		Table:    tslug,
		Id:       []int64{id},
	})
}

func (d *Controller) SimpleQuery(uclaim *claim.DataTkt, tslug string, query store.SimpleQueryReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)

	query.TenantId = uclaim.TenentId
	query.Table = tslug
	query.Group = group

	return dynDb.SimpleQuery(0, query)
}

func (d *Controller) FTSQuery(uclaim *claim.DataTkt, tslug, qstr string) (*store.QueryResult, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.FTSQuery(0, store.FTSQueryReq{
		TenantId:   uclaim.TenentId,
		Table:      tslug,
		Group:      group,
		SearchTerm: qstr,
		Count:      10,
	})
}

func (d *Controller) TemplateQuery(uclaim *claim.DataTkt, tslug string, query any) (*store.QueryResult, error) {

	// source, group := getTarget(uclaim)

	// dynDb := d.hub.GetSource(source, uclaim.TenentId)
	// // fixme
	// return dynDb.TemplateQuery(0, store.TemplateQueryReq{
	// 	TenantId:  uclaim.TenentId,
	// 	Group:     group,
	// 	Fragments: nil,
	// 	Name:      "",
	// })

	return nil, nil
}

func (d *Controller) RefResolve(uclaim *claim.DataTkt, req *store.RefResolveReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.RefResolve(0, group, req)
}

func (d *Controller) ReverseRefLoad(uclaim *claim.DataTkt, req *store.RevRefLoadReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.ReverseRefLoad(0, group, req)
}

func (d *Controller) RefLoad(uclaim *claim.DataTkt, req *store.RefLoadReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.RefLoad(0, group, req)
}

func (d *Controller) ListActivity(uclaim *claim.DataTkt, table string, rowId int) ([]*entities.DynActivity, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.ListActivity(group, table, rowId)
}

func (d *Controller) CommentRow(uclaim *claim.DataTkt, table, msg string, rowId int) error {

	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)

	return dynDb.NewActivity(group, table, &entities.DynActivity{
		Type:      "comment",
		RowId:     int64(rowId),
		RowVerson: 0,
		UserId:    uclaim.UserID,
		Payload:   msg,
	})
}
