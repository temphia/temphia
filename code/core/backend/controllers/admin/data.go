package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (c *Controller) ListSources(uclaim *claim.Session) ([]string, error) {
	return c.dynHub.ListSources(uclaim.TenentId)
}

// dyn_table_group
func (c *Controller) NewGroup(uclaim *claim.Session, source string, model *bprints.NewTableGroup) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.NewGroup(model)
}

func (c *Controller) EditGroup(uclaim *claim.Session, source, gslug string, model *entities.TableGroupPartial) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.EditGroup(gslug, model)
}

func (c *Controller) GetGroup(uclaim *claim.Session, source, gslug string) (*entities.TableGroup, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.GetGroup(gslug)
}

func (c *Controller) ListGroup(uclaim *claim.Session, source string) ([]*entities.TableGroup, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.ListGroup()
}

func (c *Controller) DeleteGroup(uclaim *claim.Session, source, gslug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.DeleteGroup(gslug)
}

// dyn_table
func (c *Controller) AddTable(uclaim *claim.Session, source, group string, model *bprints.NewTable) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.AddTable(group, model)
}

func (c *Controller) EditTable(uclaim *claim.Session, source, group, tslug string, model *entities.TablePartial) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.EditTable(group, tslug, model)
}

func (c *Controller) GetTable(uclaim *claim.Session, source, group, tslug string) (*entities.Table, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.GetTable(group, tslug)
}

func (c *Controller) ListTables(uclaim *claim.Session, source, group string) ([]*entities.Table, error) {

	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.ListTables(group)
}

func (c *Controller) DeleteTable(uclaim *claim.Session, source, group, tslug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.DeleteTable(group, tslug)
}

// dyn_table_column
func (c *Controller) AddColumn(uclaim *claim.Session, source, group, tslug string, model *bprints.NewColumn) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.AddColumn(group, tslug, model)
}

func (c *Controller) GetColumn(uclaim *claim.Session, source, group, tslug string, cslug string) (*entities.Column, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.GetColumn(group, tslug, cslug)
}

func (c *Controller) EditColumn(uclaim *claim.Session, source, group, tslug string, cslug string, model *entities.ColumnPartial) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.EditColumn(group, tslug, cslug, model)
}

func (c *Controller) ListColumns(uclaim *claim.Session, source, group, tslug string) ([]*entities.Column, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.ListColumns(group, tslug)
}

func (c *Controller) DeleteColumn(uclaim *claim.Session, source, group, tslug string, cslug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.DeleteColumn(group, tslug, cslug)
}

func (c *Controller) AddIndex(uclaim *claim.Session, source, group, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.AddIndex(group, tslug, model)
}

// dyn_table_meta
func (c *Controller) AddUniqueIndex(uclaim *claim.Session, source, group, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.AddUniqueIndex(group, tslug, model)
}

func (c *Controller) AddFTSIndex(uclaim *claim.Session, source, group, tslug string, model *entities.FTSIndex) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.AddFTSIndex(group, tslug, model)
}

func (c *Controller) AddColumnFRef(uclaim *claim.Session, source, group, tslug string, model *entities.ColumnFKRef) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.AddColumnFRef(group, tslug, model)
}

func (c *Controller) ListIndex(uclaim *claim.Session, source, group, tslug string) ([]*entities.Index, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.ListIndex(group, tslug)
}

func (c *Controller) RemoveIndex(uclaim *claim.Session, source, group, tslug, slug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.RemoveIndex(group, tslug, slug)
}

func (c *Controller) DataActivityQuery(uclaim *claim.Session, source, group, tslug string, offset int64) ([]*entities.DynActivity, error) {

	return c.dynHub.GetSource(source, uclaim.TenentId).QueryActivity(group, tslug, &entities.ActivityQuery{
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
	return c.dynHub.GetSource(source, uclaim.TenentId).SqlQuery(0, store.SqlQueryReq{
		NoTransform: false,
		Raw:         query.Raw,
		Group:       group,
		QStr:        query.QueryString,
	})
}
