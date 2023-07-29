package datahub

import (
	"sync"

	"github.com/temphia/temphia/code/backend/hub/dyndb/handle"
	"github.com/temphia/temphia/code/backend/hub/dyndb/sheet"
	"github.com/temphia/temphia/code/backend/hub/dyndb/table"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type DataSource struct {
	inner  dyndb.DynDB
	name   string
	handle *handle.Handle
	groups map[string]dyndb.DataTableHub
	gLock  sync.RWMutex
	sheets map[string]dyndb.DataSheetHub
	sLock  sync.RWMutex
}

func (ds *DataSource) Name() string { return ds.name }

func (ds *DataSource) NewGroup(tenantId string, model *xbprint.NewTableGroup) error {
	return ds.inner.NewGroup(tenantId, model)
}

func (ds *DataSource) EditGroup(tenantId, gslug string, model *entities.TableGroupPartial) error {
	return ds.inner.EditGroup(tenantId, gslug, model)
}

func (ds *DataSource) ListGroup(tenantId string, cond map[string]any) ([]*entities.TableGroup, error) {
	return ds.inner.ListGroup(tenantId, cond)
}

func (ds *DataSource) GetGroup(tenantId, gslug string) (*entities.TableGroup, error) {
	return ds.inner.GetGroup(tenantId, gslug)
}

func (ds *DataSource) DeleteGroup(tenantId, gslug string) error {
	return ds.inner.DeleteGroup(tenantId, gslug)
}

func (ds *DataSource) EditTable(tenantId, gslug, tslug string, model *entities.TablePartial) error {
	return ds.inner.EditTable(tenantId, gslug, tslug, model)
}
func (ds *DataSource) GetTable(tenantId, gslug, tslug string) (*entities.Table, error) {
	return ds.inner.GetTable(tenantId, gslug, tslug)
}

func (ds *DataSource) ListTables(tenantId, gslug string) ([]*entities.Table, error) {
	return ds.inner.ListTables(tenantId, gslug)
}

func (ds *DataSource) DeleteTable(tenantId, gslug, tslug string) error {
	return ds.inner.DeleteTable(tenantId, gslug, tslug)
}

func (ds *DataSource) EditColumn(tenantId, gslug, tslug, cslug string, model *entities.ColumnPartial) error {
	return ds.inner.EditColumn(tenantId, gslug, tslug, cslug, model)
}

func (ds *DataSource) GetColumn(tenantId, gslug, tslug, cslug string) (*entities.Column, error) {
	return ds.inner.GetColumn(tenantId, gslug, tslug, cslug)
}

func (ds *DataSource) ListColumns(tenantId, gslug, tslug string) ([]*entities.Column, error) {
	return ds.inner.ListColumns(tenantId, gslug, tslug)
}

func (ds *DataSource) ListReverseColumnRef(tenantId, gslug, tslug string) ([]*entities.Column, error) {
	return ds.inner.ListReverseColumnRef(tenantId, gslug, tslug)
}

func (ds *DataSource) DeleteColumn(tenantId, gslug, tslug, cslug string) error {
	return ds.inner.DeleteColumn(tenantId, gslug, tslug, cslug)
}

func (ds *DataSource) NewView(tenantId string, model *entities.DataView) error {
	return ds.inner.NewView(tenantId, model)
}

func (ds *DataSource) GetView(tenantId, gslug, tslug string, id int64) (*entities.DataView, error) {
	return ds.inner.GetView(tenantId, gslug, tslug, id)
}

func (ds *DataSource) ModifyView(tenantId, gslug, tslug string, id int64, data map[string]any) error {
	return ds.inner.ModifyView(tenantId, gslug, tslug, id, data)
}

func (ds *DataSource) ListView(tenantId, gslug, tslug string) ([]*entities.DataView, error) {
	return ds.inner.ListView(tenantId, gslug, tslug)
}

func (ds *DataSource) DelView(tenantId, gslug, tslug string, id int64) error {
	return ds.inner.DelView(tenantId, gslug, tslug, id)
}

func (ds *DataSource) QueryActivity(tenantId, group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error) {
	return ds.inner.QueryActivity(tenantId, group, table, query)
}

func (ds *DataSource) ListActivity(tenantId, group, table string, rowId int) ([]*entities.DynActivity, error) {
	return ds.inner.ListActivity(tenantId, group, table, rowId)
}

func (ds *DataSource) ListActivityByAlt(tenantId, group, table string, alt string) ([]*entities.DynActivity, error) {
	return ds.inner.ListActivityByAlt(tenantId, group, table, alt)
}

func (ds *DataSource) NewActivity(tenantId, group, table string, record *entities.DynActivity) error {
	_, err := ds.inner.NewActivity(tenantId, group, table, record)
	return err
}

func (ds *DataSource) MigrateSchema(tenantId string, opts step.MigrateOptions) error {
	return ds.inner.MigrateSchema(tenantId, opts)
}

func (ds *DataSource) ListDataUsers(source, tenantId, group, ttype, target string) ([]entities.UserInfo, error) {

	// fixme use ttype and target
	// check if that column only have some user access

	users, err := ds.handle.CoreHub.ListUsers(tenantId)
	if err != nil {
		return nil, err
	}

	iusers := make([]entities.UserInfo, 0, len(users))

	for _, u := range users {
		iusers = append(iusers, entities.UserInfo{
			UserId:    u.UserId,
			FullName:  u.FullName,
			Bio:       u.Bio,
			PublicKey: u.PublicKey,
			Email:     u.Email,
			GroupId:   u.GroupID,
		})
	}

	return iusers, nil
}

func (ds *DataSource) GetDataTableHub(tenantId, group string) dyndb.DataTableHub {

	ds.gLock.RLock()
	dh := ds.groups[tenantId+group]
	ds.gLock.RUnlock()

	if dh != nil {
		return dh
	}

	// fixme => we are creating hub without validating if group exists
	dh = table.New(ds.name, tenantId, group, ds.inner, ds.handle)

	ds.gLock.Lock()
	ds.groups[tenantId+group] = dh
	ds.gLock.Unlock()

	return dh
}

func (ds *DataSource) GetDataSheetHub(tenantId, group string) dyndb.DataSheetHub {
	ds.sLock.RLock()
	dh := ds.sheets[tenantId+group]
	ds.sLock.RUnlock()

	if dh != nil {
		return dh
	}

	dthub := ds.GetDataTableHub(tenantId, group)

	// fixme => we are creating hub without validating if group exists
	dh = sheet.New(dthub, ds.handle, ds.name, tenantId, group)

	ds.sLock.Lock()
	ds.sheets[tenantId+group] = dh
	ds.sLock.Unlock()

	return dh

}
