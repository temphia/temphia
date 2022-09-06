package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
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
func (c *Controller) AddTable(uclaim *claim.Session, model *bprints.NewTable) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.AddTable(uclaim.Path[2], model)
}

func (c *Controller) EditTable(uclaim *claim.Session, tslug string, model *entities.TablePartial) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.EditTable(uclaim.Path[2], tslug, model)
}

func (c *Controller) GetTable(uclaim *claim.Session, tslug string) (*entities.Table, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.GetTable(uclaim.Path[2], tslug)
}

func (c *Controller) ListTables(uclaim *claim.Session) ([]*entities.Table, error) {

	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.ListTables(uclaim.Path[2])
}

func (c *Controller) DeleteTable(uclaim *claim.Session, tslug string) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.DeleteTable(uclaim.Path[2], tslug)
}

// dyn_table_column
func (c *Controller) AddColumn(uclaim *claim.Session, tslug string, model *bprints.NewColumn) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.AddColumn(uclaim.Path[2], tslug, model)
}

func (c *Controller) GetColumn(uclaim *claim.Session, tslug, cslug string) (*entities.Column, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.GetColumn(uclaim.Path[2], tslug, cslug)
}

func (c *Controller) EditColumn(uclaim *claim.Session, tslug, cslug string, model *entities.ColumnPartial) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.EditColumn(uclaim.Path[2], tslug, cslug, model)
}

func (c *Controller) ListColumns(uclaim *claim.Session, tslug string) ([]*entities.Column, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.ListColumns(uclaim.Path[2], tslug)
}

func (c *Controller) DeleteColumn(uclaim *claim.Session, tslug, cslug string) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.DeleteColumn(uclaim.Path[2], tslug, cslug)
}

func (c *Controller) AddIndex(uclaim *claim.Session, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.AddIndex(uclaim.Path[2], tslug, model)
}

// dyn_table_meta
func (c *Controller) AddUniqueIndex(uclaim *claim.Session, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.AddUniqueIndex(uclaim.Path[2], tslug, model)
}

func (c *Controller) AddFTSIndex(uclaim *claim.Session, tslug string, model *entities.FTSIndex) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.AddFTSIndex(uclaim.Path[2], tslug, model)
}

func (c *Controller) AddColumnFRef(uclaim *claim.Session, tslug string, model *entities.ColumnFKRef) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.AddColumnFRef(uclaim.Path[2], tslug, model)
}

func (c *Controller) ListIndex(uclaim *claim.Session, tslug string) ([]*entities.Index, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.ListIndex(uclaim.Path[2], tslug)
}

func (c *Controller) RemoveIndex(uclaim *claim.Session, tslug, slug string) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.RemoveIndex(uclaim.Path[2], tslug, slug)
}
