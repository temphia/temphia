package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/scopes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (c *Controller) ListSources(uclaim *claim.Session) ([]string, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	return c.dynHub.ListSources(uclaim.TenantId)
}

// dyn_table_group
func (c *Controller) NewGroup(uclaim *claim.Session, source string, model *xbprint.NewTableGroup) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.NewGroup(uclaim.TenantId, model)
}

func (c *Controller) EditGroup(uclaim *claim.Session, source, gslug string, model *entities.TableGroupPartial) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.EditGroup(uclaim.TenantId, gslug, model)
}

func (c *Controller) GetGroup(uclaim *claim.Session, source, gslug string) (*entities.TableGroup, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.GetGroup(uclaim.TenantId, gslug)
}

func (c *Controller) GetGroupSheets(uclaim *claim.Session, source, gslug string) (any, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	sheet := dynDB.GetDataSheetHub(uclaim.TenantId, gslug)
	return sheet.ListSheet(0)
}

func (c *Controller) ListGroup(uclaim *claim.Session, source string) ([]*entities.TableGroup, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.ListGroup(uclaim.TenantId)
}

func (c *Controller) DeleteGroup(uclaim *claim.Session, source, gslug string) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.DeleteGroup(uclaim.TenantId, gslug)
}

// dyn_table

func (c *Controller) EditTable(uclaim *claim.Session, source, group, tslug string, model *entities.TablePartial) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.EditTable(uclaim.TenantId, group, tslug, model)
}

func (c *Controller) GetTable(uclaim *claim.Session, source, group, tslug string) (*entities.Table, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.GetTable(uclaim.TenantId, group, tslug)
}

func (c *Controller) ListTables(uclaim *claim.Session, source, group string) ([]*entities.Table, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.ListTables(uclaim.TenantId, group)
}

func (c *Controller) DeleteTable(uclaim *claim.Session, source, group, tslug string) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.DeleteTable(uclaim.TenantId, group, tslug)
}

// dyn_table_column

func (c *Controller) GetColumn(uclaim *claim.Session, source, group, tslug string, cslug string) (*entities.Column, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.GetColumn(uclaim.TenantId, group, tslug, cslug)
}

func (c *Controller) EditColumn(uclaim *claim.Session, source, group, tslug string, cslug string, model *entities.ColumnPartial) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.EditColumn(uclaim.TenantId, group, tslug, cslug, model)
}

func (c *Controller) ListColumns(uclaim *claim.Session, source, group, tslug string) ([]*entities.Column, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.ListColumns(uclaim.TenantId, group, tslug)
}

func (c *Controller) DeleteColumn(uclaim *claim.Session, source, group, tslug string, cslug string) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.DeleteColumn(uclaim.TenantId, group, tslug, cslug)
}

func (c *Controller) DataActivityQuery(uclaim *claim.Session, source, group, tslug string, offset int64) ([]*entities.DynActivity, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	return c.dynHub.GetSource(source, uclaim.TenantId).QueryActivity(uclaim.TenantId, group, tslug, &entities.ActivityQuery{
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
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	src := c.dynHub.GetSource(source, uclaim.TenantId)

	return src.GetDataTableHub(uclaim.TenantId, group).SqlQuery(0, dyndb.SqlQueryReq{
		NoTransform: false,
		Raw:         query.Raw,
		Group:       group,
		QStr:        query.QueryString,
	})

}

func (c *Controller) LiveSeed(uclaim *claim.Session, source, group, table string, max int) error {
	src := c.dynHub.GetSource(source, uclaim.TenantId)

	return src.GetDataTableHub(uclaim.TenantId, group).LiveSeed(table, uclaim.UserID, max)
}
