package xbprint

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type NewTableGroup struct {
	Name          string      `json:"name,omitempty" yaml:"name,omitempty"`
	Slug          string      `json:"slug,omitempty" yaml:"slug,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Tables        []*NewTable `json:"tables,omitempty" yaml:"tables,omitempty"`
	ExecOrder     []string    `json:"exec_order,omitempty" yaml:"exec_order,omitempty"`
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
	SeedData      SeedData                `json:"seed_data,omitempty" yaml:"seed_data,omitempty"`
}

type NewColumn struct {
	Name          string   `json:"name,omitempty" yaml:"name,omitempty"`
	Slug          string   `json:"slug,omitempty" yaml:"slug,omitempty"`
	Ctype         string   `json:"ctype,omitempty" yaml:"ctype,omitempty"`
	Description   string   `json:"description,omitempty" yaml:"description,omitempty"`
	Icon          string   `json:"icon,omitempty" yaml:"icon,omitempty"`
	Options       []string `json:"options,omitempty" yaml:"options,omitempty"`
	NotNullable   bool     `json:"not_nullable,omitempty" yaml:"not_nullable,omitempty"`
	Pattern       string   `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	StrictPattern bool     `json:"strict_pattern,omitempty" yaml:"strict_pattern,omitempty"`

	// secondary objects
	TargetApps []NewTargetApp `json:"target_apps,omitempty" yaml:"target_apps,omitempty"`
}

type NewTargetApp struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Type  string `json:"type,omitempty" yaml:"type,omitempty"`
	Plug  string `json:"plug,omitempty" yaml:"plug,omitempty"`
	Agent string `json:"agent,omitempty" yaml:"agent,omitempty"`
}

type SeedData struct {
	Data         []map[string]any `json:"data,omitempty" yaml:"data,omitempty"`
	LinkedImages []string         `json:"linked_images,omitempty" yaml:"linked_images,omitempty"`
}
