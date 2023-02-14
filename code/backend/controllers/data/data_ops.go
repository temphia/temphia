package data

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (c *Controller) LoadGroup(uclaim *claim.Data) (*dyndb.LoadDgroupResp, error) {
	dynDB := c.dynHub.GetSource(uclaim.DataSource, uclaim.TenantId)

	tg, err := dynDB.GetGroup(uclaim.DataGroup)
	if err != nil {
		return nil, err
	}

	tables, err := dynDB.ListTables(uclaim.DataGroup)
	if err != nil {
		return nil, err
	}

	fclaim, err := c.folderTicket(tg, uclaim)
	if err != nil {
		return nil, err
	}

	resp := &dyndb.LoadDgroupResp{
		Tables:       tables,
		FolderTicket: fclaim,
	}

	return resp, nil
}

func (d *Controller) NewRow(uclaim *claim.Data, tslug string, cells map[string]any) (int64, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.NewRow(0, dyndb.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    tslug,
		Data:     cells,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (d *Controller) GetRow(uclaim *claim.Data, tslug string, id int64) (map[string]any, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.GetRow(0, dyndb.GetRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    tslug,
		Id:       id,
	})
}

func (d *Controller) UpdateRow(uclaim *claim.Data, tslug string, id, version int64, cells map[string]any) (map[string]any, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)
	return dynDb.UpdateRow(0, dyndb.UpdateRowReq{
		TenantId: uclaim.TenantId,
		Version:  version,
		Group:    group,
		Table:    tslug,
		Data:     cells,
		Id:       id,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (d *Controller) DeleteRowBatch(uclaim *claim.Data, tslug string, filterOpts []*dyndb.FilterCond) error {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId:    uclaim.TenantId,
		Group:       group,
		Table:       tslug,
		FilterConds: filterOpts,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (d *Controller) DeleteRowMulti(uclaim *claim.Data, tslug string, ids []int64) error {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.DeleteRowMulti(0, dyndb.DeleteRowMultiReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    tslug,
		Ids:      ids,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

}

func (d *Controller) DeleteRow(uclaim *claim.Data, tslug string, id int64) error {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.DeleteRow(0, dyndb.DeleteRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    tslug,
		Id:       id,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

}

func (c *Controller) LoadTable(uclaim *claim.Data, req dyndb.LoadTableReq, tslug string) (*dyndb.LoadTableResp, error) {

	source, group := getTarget(uclaim)
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	req.TenantId = uclaim.TenantId
	req.Table = tslug
	req.Group = group

	resp, err := dynDB.LoadTable(0, req)
	if err != nil {
		return nil, err
	}

	// fixme => load user and folder tokens here

	return resp, nil
}

func (d *Controller) SimpleQuery(uclaim *claim.Data, tslug string, query dyndb.SimpleQueryReq) (*dyndb.QueryResult, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)

	query.TenantId = uclaim.TenantId
	query.Table = tslug
	query.Group = group

	return dynDb.SimpleQuery(0, query)
}

func (d *Controller) FTSQuery(uclaim *claim.Data, tslug, qstr string) (*dyndb.QueryResult, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)
	return dynDb.FTSQuery(0, dyndb.FTSQueryReq{
		TenantId:   uclaim.TenantId,
		Table:      tslug,
		Group:      group,
		SearchTerm: qstr,
		Count:      10,
	})
}

func (d *Controller) TemplateQuery(uclaim *claim.Data, tslug string, query any) (*dyndb.QueryResult, error) {

	// source, group := getTarget(uclaim)

	// dynDb := d.hub.GetSource(source, uclaim.TenantId)
	// // fixme
	// return dynDb.TemplateQuery(0, dyndb.TemplateQueryReq{
	// 	TenantId:  uclaim.TenantId,
	// 	Group:     group,
	// 	Fragments: nil,
	// 	Name:      "",
	// })

	return nil, nil
}

func (d *Controller) RefResolve(uclaim *claim.Data, req *dyndb.RefResolveReq) (*dyndb.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)
	return dynDb.RefResolve(0, group, req)
}

func (d *Controller) ReverseRefLoad(uclaim *claim.Data, req *dyndb.RevRefLoadReq) (*dyndb.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)
	return dynDb.ReverseRefLoad(0, group, req)
}

func (d *Controller) RefLoad(uclaim *claim.Data, req *dyndb.RefLoadReq) (*dyndb.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)
	return dynDb.RefLoad(0, group, req)
}

func (d *Controller) ListActivity(uclaim *claim.Data, table string, rowId int) ([]*entities.DynActivity, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)
	return dynDb.ListActivity(group, table, rowId)
}

func (d *Controller) CommentRow(uclaim *claim.Data, table, msg string, rowId int) error {

	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.NewActivity(group, table, &entities.DynActivity{
		Type:      "comment",
		RowId:     int64(rowId),
		RowVerson: 0,
		UserId:    uclaim.UserID,
		Payload:   msg,
	})
}

func (c *Controller) folderTicket(group *entities.TableGroup, uclaim *claim.Data) (string, error) {

	if group.CabinetSource == "" || group.CabinetFolder == "" {
		group.CabinetSource = c.cabHub.DefaultName(uclaim.TenantId)
		group.CabinetFolder = store.DefaultDataAssetsFolder
	}

	fcalim := &claim.Folder{
		Folder:    group.CabinetFolder,
		Source:    group.CabinetSource,
		Expiry:    0,
		TenantId:  uclaim.TenantId,
		UserId:    uclaim.UserID,
		SessionID: uclaim.SessionID,
	}

	return c.signer.SignFolder(uclaim.TenantId, fcalim)

}
