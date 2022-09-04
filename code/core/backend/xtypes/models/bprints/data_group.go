package bprints

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

type NewTableGroup struct {
	Name          string      `json:"name,omitempty" yaml:"name,omitempty"`
	Slug          string      `json:"slug,omitempty" yaml:"slug,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Tables        []*NewTable `json:"tables,omitempty" yaml:"tables,omitempty"`
	ExecOrder     []string    `json:"exec_order,omitempty" yaml:"exec_order,omitempty"`
	Source        string      `json:"-"`
	CabinetSource string      `json:"-"`
	CabinetFolder string      `json:"-"`
}

type NewTable struct {
	Name          string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Slug          string                  `json:"slug,omitempty" yaml:"slug,omitempty"`
	Description   string                  `json:"description,omitempty" yaml:"description,omitempty"`
	Icon          string                  `json:"icon,omitempty" yaml:"icon,omitempty"`
	MainColumn    string                  `json:"main_column,omitempty" yaml:"main_column,omitempty"`
	ActivityType  string                  `json:"activity_type,omitempty" yaml:"activity_type,omitempty"`
	SyncType      string                  `json:"sync_type,omitempty" yaml:"sync_type,omitempty"`
	Columns       []*NewColumn            `json:"columns,omitempty" yaml:"columns,omitempty"`
	Indexes       []entities.Index        `json:"indexes,omitempty" yaml:"indexes,omitempty"`
	UniqueIndexes []entities.Index        `json:"unique_indexes,omitempty" yaml:"unique_indexes,omitempty"`
	FTSIndex      *entities.FTSIndex      `json:"fts_index,omitempty" yaml:"fts_index,omitempty"`
	ColumnRef     []*entities.ColumnFKRef `json:"column_refs,omitempty" yaml:"column_refs,omitempty"`
	DeletedAt     bool                    `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
	Views         []entities.View         `json:"views,omitempty" yaml:"views,omitempty"`
	SeedData      SeedData                `json:"seed_data,omitempty" yaml:"seed_data,omitempty"`
}

type NewColumn struct {
	Name        string `json:"name,omitempty" db:"name,omitempty" yaml:"name,omitempty"`
	Slug        string `json:"slug,omitempty" db:"slug,omitempty" yaml:"slug,omitempty"`
	Ctype       string `json:"ctype,omitempty" db:"ctype,omitempty" yaml:"ctype,omitempty"`
	Description string `json:"description,omitempty" db:"description,omitempty" yaml:"description,omitempty"`
	Icon        string `json:"icon,omitempty" db:"icon,omitempty" yaml:"icon,omitempty"`

	Options       []string `json:"options,omitempty" db:"options,omitempty" yaml:"options,omitempty"`
	NotNullable   bool     `json:"not_nullable,omitempty" db:"not_nullable,omitempty" yaml:"not_nullable,omitempty"`
	Pattern       string   `json:"pattern,omitempty" db:"pattern,omitempty" yaml:"pattern,omitempty"`
	StrictPattern bool     `json:"strict_pattern,omitempty" db:"strict_pattern,omitempty" yaml:"strict_pattern,omitempty"`
}

type SeedData struct {
	Data         []map[string]any `json:"data,omitempty" yaml:"data,omitempty"`
	LinkedImages []string         `json:"linked_images,omitempty" yaml:"linked_images,omitempty"`
}

func (m *NewTableGroup) To(tenantId string) *entities.TableGroup {
	return &entities.TableGroup{
		Name:          m.Name,
		Slug:          m.Slug,
		Description:   m.Description,
		SourceDb:      m.Source,
		TenantID:      tenantId,
		CabinetSource: m.CabinetSource,
		CabinetFolder: m.CabinetFolder,
		Active:        false,
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
		TableID:       tslug,
		TenantID:      tenantId,
		RefId:         "",
		RefType:       "",
		RefTarget:     "",
		RefObject:     "",
		RefCopy:       "",
		ExtraMeta:     nil,
	}
}
