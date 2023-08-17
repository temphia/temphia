package xpackage

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type DataSchema struct {
	Steps []MigrationStep `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type MigrationStep struct {
	Name string          `json:"name,omitempty" yaml:"name,omitempty"`
	Type string          `json:"type,omitempty" yaml:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty" yaml:"data,omitempty"`
}

type NewTableGroup struct {
	Name          string      `json:"name,omitempty" yaml:"name,omitempty"`
	Slug          string      `json:"slug,omitempty" yaml:"slug,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Tables        []*NewTable `json:"tables,omitempty" yaml:"tables,omitempty"`
	ExecOrder     []string    `json:"exec_order,omitempty" yaml:"exec_order,omitempty"`
	Renderer      string      `json:"renderer,omitempty" yaml:"renderer,omitempty"`
	CabinetSource string      `json:"-" yaml:"-"`
	CabinetFolder string      `json:"-" yaml:"-"`
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
	Views         []entities.View         `json:"views,omitempty" yaml:"views,omitempty"`
}

type RemoveTable struct {
	Slug string `json:"slug,omitempty" yaml:"slug,omitempty"`
}

type NewColumn struct {
	Name          string   `json:"name,omitempty" yaml:"name,omitempty"`
	Slug          string   `json:"slug,omitempty" yaml:"slug,omitempty"`
	Table         string   `json:"table,omitempty" yaml:"table,omitempty"`
	Ctype         string   `json:"ctype,omitempty" yaml:"ctype,omitempty"`
	Description   string   `json:"description,omitempty" yaml:"description,omitempty"`
	Icon          string   `json:"icon,omitempty" yaml:"icon,omitempty"`
	Options       []string `json:"options,omitempty" yaml:"options,omitempty"`
	NotNullable   bool     `json:"not_nullable,omitempty" yaml:"not_nullable,omitempty"`
	Pattern       string   `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	StrictPattern bool     `json:"strict_pattern,omitempty" yaml:"strict_pattern,omitempty"`
}

type RemoveColumn struct {
	Slug  string `json:"slug,omitempty" yaml:"slug,omitempty"`
	Table string `json:"table,omitempty" yaml:"table,omitempty"`
}

type NewTargetApp struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Agent       string `json:"agent,omitempty" yaml:"agent,omitempty"`
	ContextType string `json:"context_type,omitempty" yaml:"context_type,omitempty"`
}
