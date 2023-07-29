package data

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (c *Controller) LoadGroup(uclaim *claim.Data) (*dyndb.LoadDgroupResp, error) {

	dynDB := c.dynHub.GetDynDB()

	tg, err := dynDB.GetGroup(uclaim.TenantId, uclaim.DataGroup)
	if err != nil {
		return nil, err
	}

	tables, err := dynDB.ListTables(uclaim.TenantId, uclaim.DataGroup)
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

func (c *Controller) NewRow(uclaim *claim.Data, tslug string, cells map[string]any) (int64, error) {

	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.NewRow(0, dyndb.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    uclaim.DataGroup,
		Table:    tslug,
		Data:     cells,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (c *Controller) GetRow(uclaim *claim.Data, tslug string, id int64) (map[string]any, error) {
	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.GetRow(0, dyndb.GetRowReq{
		TenantId: uclaim.TenantId,
		Group:    uclaim.DataGroup,
		Table:    tslug,
		Id:       id,
	})
}

func (c *Controller) UpdateRow(uclaim *claim.Data, tslug string, id, version int64, cells map[string]any) (map[string]any, error) {
	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.UpdateRow(0, dyndb.UpdateRowReq{
		TenantId: uclaim.TenantId,
		Version:  version,
		Group:    uclaim.DataGroup,
		Table:    tslug,
		Data:     cells,
		Id:       id,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (c *Controller) DeleteRowBatch(uclaim *claim.Data, tslug string, filterOpts []dyndb.FilterCond) error {
	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	_, err := thub.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId:    uclaim.TenantId,
		Group:       uclaim.DataGroup,
		Table:       tslug,
		FilterConds: filterOpts,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return err
}

func (c *Controller) DeleteRowMulti(uclaim *claim.Data, tslug string, ids []int64) error {
	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.DeleteRowMulti(0, dyndb.DeleteRowMultiReq{
		TenantId: uclaim.TenantId,
		Group:    uclaim.DataGroup,
		Table:    tslug,
		Ids:      ids,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

}

func (c *Controller) DeleteRow(uclaim *claim.Data, tslug string, id int64) error {
	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.DeleteRow(0, dyndb.DeleteRowReq{
		TenantId: uclaim.TenantId,
		Group:    uclaim.DataGroup,
		Table:    tslug,
		Id:       id,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

}

func (c *Controller) LoadTable(uclaim *claim.Data, req dyndb.LoadTableReq, tslug string) (*dyndb.LoadTableResp, error) {

	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	req.TenantId = uclaim.TenantId
	req.Table = tslug
	req.Group = uclaim.DataGroup

	return thub.LoadTable(0, req)

}

func (c *Controller) SimpleQuery(uclaim *claim.Data, tslug string, query dyndb.SimpleQueryReq) (*dyndb.QueryResult, error) {

	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	query.TenantId = uclaim.TenantId
	query.Table = tslug
	query.Group = uclaim.DataGroup

	return thub.SimpleQuery(0, query)
}

func (c *Controller) FTSQuery(uclaim *claim.Data, req dyndb.FTSQueryReq) (*dyndb.QueryResult, error) {

	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	req.Group = uclaim.DataGroup
	req.TenantId = uclaim.TenantId

	return thub.FTSQuery(0, req)
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

func (c *Controller) RefResolve(uclaim *claim.Data, req *dyndb.RefResolveReq) (*dyndb.QueryResult, error) {

	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.RefResolve(0, uclaim.DataGroup, req)
}

func (c *Controller) ReverseRefLoad(uclaim *claim.Data, req *dyndb.RevRefLoadReq) (*dyndb.QueryResult, error) {

	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.ReverseRefLoad(0, uclaim.DataGroup, req)
}

func (c *Controller) RefLoad(uclaim *claim.Data, req *dyndb.RefLoadReq) (*dyndb.QueryResult, error) {

	thub := c.dynHub.GetDataTableHub(uclaim.TenantId, uclaim.DataGroup)

	return thub.RefLoad(0, uclaim.DataGroup, req)
}

func (c *Controller) ListActivity(uclaim *claim.Data, table string, rowId int) ([]*entities.DynActivity, error) {

	return c.dynHub.GetDynDB().ListActivity(uclaim.TenantId, uclaim.DataGroup, table, rowId)
}

func (c *Controller) CommentRow(uclaim *claim.Data, table, msg string, rowId int) error {

	dyndb := c.dynHub.GetDynDB()

	_, err := dyndb.NewActivity(uclaim.TenantId, uclaim.DataGroup, table, &entities.DynActivity{
		Type:      "comment",
		RowId:     int64(rowId),
		RowVerson: 0,
		UserId:    uclaim.UserID,
		Payload:   msg,
	})

	return err

}

func (c *Controller) folderTicket(group *entities.TableGroup, uclaim *claim.Data) (string, error) {

	if group.CabinetFolder == "" {
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

func (c *Controller) ListDataUsers(uclaim *claim.Data, ttype, target string) ([]entities.UserInfo, error) {

	//	return c.dynHub.GetDynDB().ListDataUsers("default", uclaim.TenantId, uclaim.DataGroup, ttype, target)

	return nil, nil
}
