package xpackage

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (m *NewTableGroup) To(tenantId string) *entities.TableGroup {

	return &entities.TableGroup{
		Name:          m.Name,
		Slug:          m.Slug,
		Description:   m.Description,
		TenantID:      tenantId,
		Active:        false,
		CabinetSource: m.CabinetSource,
		CabinetFolder: m.CabinetFolder,
		Renderer:      m.Renderer,
	}
}

func (m *NewTable) To(tenantId, gslug string) *entities.Table {
	return &entities.Table{
		Name:        m.Name,
		Slug:        m.Slug,
		Description: m.Description,
		Icon:        m.Icon,
		GroupID:     gslug,
		TenantID:    tenantId,
		MainColumn:  "",
	}
}

func (m *NewColumn) To(tenantId, gslug, tslug string) *entities.Column {
	return &entities.Column{
		Name:          m.Name,
		Slug:          m.Slug,
		Ctype:         m.Ctype,
		Description:   m.Description,
		GroupID:       gslug,
		Icon:          m.Icon,
		Options:       m.Options,
		OrderID:       0,
		Pattern:       m.Pattern,
		StrictPattern: m.StrictPattern,
		NotNullable:   m.NotNullable,
		TableID:       tslug,
		TenantID:      tenantId,
		RefId:         "",
		RefType:       "",
		RefTarget:     "",
		RefObject:     "",
		ExtraMeta:     nil,
	}
}
