package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (c *Controller) ListSources(uclaim *claim.Session) ([]string, error) {
	return c.dynHub.ListSources(uclaim.TenantId)
}

// dyn_table_group
func (c *Controller) NewGroup(uclaim *claim.Session, source string, model *xbprint.NewTableGroup) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.NewGroup(model)
}

func (c *Controller) EditGroup(uclaim *claim.Session, source, gslug string, model *entities.TableGroupPartial) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.EditGroup(gslug, model)
}

func (c *Controller) GetGroup(uclaim *claim.Session, source, gslug string) (*entities.TableGroup, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.GetGroup(gslug)
}

func (c *Controller) ListGroup(uclaim *claim.Session, source string) ([]*entities.TableGroup, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.ListGroup()
}

func (c *Controller) DeleteGroup(uclaim *claim.Session, source, gslug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.DeleteGroup(gslug)
}

// dyn_table
func (c *Controller) AddTable(uclaim *claim.Session, source, group string, model *xbprint.NewTable) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.AddTable(group, model)
}

func (c *Controller) EditTable(uclaim *claim.Session, source, group, tslug string, model *entities.TablePartial) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.EditTable(group, tslug, model)
}

func (c *Controller) GetTable(uclaim *claim.Session, source, group, tslug string) (*entities.Table, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.GetTable(group, tslug)
}

func (c *Controller) ListTables(uclaim *claim.Session, source, group string) ([]*entities.Table, error) {

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.ListTables(group)
}

func (c *Controller) DeleteTable(uclaim *claim.Session, source, group, tslug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.DeleteTable(group, tslug)
}

// dyn_table_column
func (c *Controller) AddColumn(uclaim *claim.Session, source, group, tslug string, model *xbprint.NewColumn) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.AddColumn(group, tslug, model)
}

func (c *Controller) GetColumn(uclaim *claim.Session, source, group, tslug string, cslug string) (*entities.Column, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.GetColumn(group, tslug, cslug)
}

func (c *Controller) EditColumn(uclaim *claim.Session, source, group, tslug string, cslug string, model *entities.ColumnPartial) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.EditColumn(group, tslug, cslug, model)
}

func (c *Controller) ListColumns(uclaim *claim.Session, source, group, tslug string) ([]*entities.Column, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.ListColumns(group, tslug)
}

func (c *Controller) DeleteColumn(uclaim *claim.Session, source, group, tslug string, cslug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.DeleteColumn(group, tslug, cslug)
}

func (c *Controller) AddIndex(uclaim *claim.Session, source, group, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.AddIndex(group, tslug, model)
}

// dyn_table_meta
func (c *Controller) AddUniqueIndex(uclaim *claim.Session, source, group, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.AddUniqueIndex(group, tslug, model)
}

func (c *Controller) AddFTSIndex(uclaim *claim.Session, source, group, tslug string, model *entities.FTSIndex) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.AddFTSIndex(group, tslug, model)
}

func (c *Controller) AddColumnFRef(uclaim *claim.Session, source, group, tslug string, model *entities.ColumnFKRef) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.AddColumnFRef(group, tslug, model)
}

func (c *Controller) ListIndex(uclaim *claim.Session, source, group, tslug string) ([]*entities.Index, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.ListIndex(group, tslug)
}

func (c *Controller) RemoveIndex(uclaim *claim.Session, source, group, tslug, slug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.RemoveIndex(group, tslug, slug)
}

func (c *Controller) DataActivityQuery(uclaim *claim.Session, source, group, tslug string, offset int64) ([]*entities.DynActivity, error) {

	return c.dynHub.GetSource(source, uclaim.TenantId).QueryActivity(group, tslug, &entities.ActivityQuery{
		Types:       nil,
		UserId:      "",
		BetweenTime: [2]string{},
		Count:       100,
		Offset:      offset,
	})

}

type DataGroupQuery struct {
	Raw         bool   `json:"raw,omitempty"`
	QueryString string `json:"query_string,omitempty"`
}

func (c *Controller) QueryDataGroup(uclaim *claim.Session, source, group string, query DataGroupQuery) (any, error) {
	return c.dynHub.GetSource(source, uclaim.TenantId).SqlQuery(0, store.SqlQueryReq{
		NoTransform: false,
		Raw:         query.Raw,
		Group:       group,
		QStr:        query.QueryString,
	})
}

func (c *Controller) LiveSeed(uclaim *claim.Session, source, group, table string, max int) error {
	src := c.dynHub.GetSource(source, uclaim.TenantId)
	return src.LiveSeed(group, table, uclaim.UserID, max)
}
