package dtable

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func getTarget(uclaim *claim.Session) (string, string) {
	return uclaim.Path[1], uclaim.Path[2]
}

func (d *Controller) NewRow(uclaim *claim.Session, tslug string, cells map[string]interface{}) (int64, error) {

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

func (d *Controller) GetRow(uclaim *claim.Session, tslug string, id int64) (map[string]interface{}, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)

	return dynDb.GetRow(0, store.GetRowReq{
		TenantId: uclaim.TenentId,
		Group:    group,
		Table:    tslug,
		Id:       id,
	})
}

func (d *Controller) UpdateRow(uclaim *claim.Session, tslug string, id, version int64, cells map[string]interface{}) (map[string]interface{}, error) {

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

func (d *Controller) DeleteRow(uclaim *claim.Session, tslug string, id int64) error {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.DeleteRows(0, store.DeleteRowReq{
		TenantId: uclaim.TenentId,
		Group:    group,
		Table:    tslug,
		Id:       []int64{id},
	})
}

func (d *Controller) SimpleQuery(uclaim *claim.Session, tslug string, query store.SimpleQueryReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)

	query.TenantId = uclaim.TenentId
	query.Table = tslug
	query.Group = group

	return dynDb.SimpleQuery(0, query)
}

func (d *Controller) FTSQuery(uclaim *claim.Session, tslug, qstr string) (*store.QueryResult, error) {

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

func (d *Controller) TemplateQuery(uclaim *claim.Session, tslug string, query interface{}) (*store.QueryResult, error) {

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

func (d *Controller) RefResolve(uclaim *claim.Session, req *store.RefResolveReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.RefResolve(0, group, req)
}

func (d *Controller) ReverseRefLoad(uclaim *claim.Session, req *store.RevRefLoadReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.ReverseRefLoad(0, group, req)
}

func (d *Controller) RefLoad(uclaim *claim.Session, req *store.RefLoadReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.RefLoad(0, group, req)
}

func (d *Controller) ListActivity(uclaim *claim.Session, table string, rowId int) ([]*entities.DynActivity, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentId)
	return dynDb.ListActivity(group, table, rowId)
}

func (d *Controller) CommentRow(uclaim *claim.Session, table, msg string, rowId int) error {

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
