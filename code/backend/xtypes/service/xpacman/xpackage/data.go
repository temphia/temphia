package xpackage

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type MigrateOptions struct {
	Steps  []MigrationStep `json:"steps,omitempty" toml:"steps,omitempty"`
	New    bool            `json:"new,omitempty" toml:"new,omitempty"`
	Gslug  string          `json:"gslug,omitempty" toml:"gslug,omitempty"`
	PlugId string          `json:"plug_id,omitempty" toml:"plug_id,omitempty"`
	DryRun bool            `json:"-"`
}

type MigrationStep struct {
	Name string          `json:"name,omitempty" toml:"name,omitempty"`
	Type string          `json:"type,omitempty" toml:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty" toml:"data,omitempty"`
}

type NewTableGroup struct {
	Name          string      `json:"name,omitempty" toml:"name,omitempty"`
	Slug          string      `json:"slug,omitempty" toml:"slug,omitempty"`
	Description   string      `json:"description,omitempty" toml:"description,omitempty"`
	Tables        []*NewTable `json:"tables,omitempty" toml:"tables,omitempty"`
	ExecOrder     []string    `json:"exec_order,omitempty" toml:"exec_order,omitempty"`
	Renderer      string      `json:"renderer,omitempty" toml:"renderer,omitempty"`
	CabinetSource string      `json:"-" toml:"-"`
	CabinetFolder string      `json:"-" toml:"-"`
}

type NewTable struct {
	Name          string                  `json:"name,omitempty" toml:"name,omitempty"`
	Slug          string                  `json:"slug,omitempty" toml:"slug,omitempty"`
	Description   string                  `json:"description,omitempty" toml:"description,omitempty"`
	Icon          string                  `json:"icon,omitempty" toml:"icon,omitempty"`
	MainColumn    string                  `json:"main_column,omitempty" toml:"main_column,omitempty"`
	ActivityType  string                  `json:"activity_type,omitempty" toml:"activity_type,omitempty"`
	SyncType      string                  `json:"sync_type,omitempty" toml:"sync_type,omitempty"`
	Columns       []*NewColumn            `json:"columns,omitempty" toml:"columns,omitempty"`
	Indexes       []entities.Index        `json:"indexes,omitempty" toml:"indexes,omitempty"`
	UniqueIndexes []entities.Index        `json:"unique_indexes,omitempty" toml:"unique_indexes,omitempty"`
	FTSIndex      *entities.FTSIndex      `json:"fts_index,omitempty" toml:"fts_index,omitempty"`
	ColumnRef     []*entities.ColumnFKRef `json:"column_refs,omitempty" toml:"column_refs,omitempty"`
	Views         []entities.View         `json:"views,omitempty" toml:"views,omitempty"`
}

type RemoveTable struct {
	Slug string `json:"slug,omitempty" toml:"slug,omitempty"`
}

type NewColumn struct {
	Name          string   `json:"name,omitempty" toml:"name,omitempty"`
	Slug          string   `json:"slug,omitempty" toml:"slug,omitempty"`
	Table         string   `json:"table,omitempty" toml:"table,omitempty"`
	Ctype         string   `json:"ctype,omitempty" toml:"ctype,omitempty"`
	Description   string   `json:"description,omitempty" toml:"description,omitempty"`
	Icon          string   `json:"icon,omitempty" toml:"icon,omitempty"`
	Options       []string `json:"options,omitempty" toml:"options,omitempty"`
	NotNullable   bool     `json:"not_nullable,omitempty" toml:"not_nullable,omitempty"`
	Pattern       string   `json:"pattern,omitempty" toml:"pattern,omitempty"`
	StrictPattern bool     `json:"strict_pattern,omitempty" toml:"strict_pattern,omitempty"`
}

type RemoveColumn struct {
	Slug  string `json:"slug,omitempty" toml:"slug,omitempty"`
	Table string `json:"table,omitempty" toml:"table,omitempty"`
}

type NewTargetApp struct {
	Name        string `json:"name,omitempty" toml:"name,omitempty"`
	Agent       string `json:"agent,omitempty" toml:"agent,omitempty"`
	ContextType string `json:"context_type,omitempty" toml:"context_type,omitempty"`
}
