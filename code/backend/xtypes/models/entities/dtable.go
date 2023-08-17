package entities

import (
	"github.com/lib/pq"
)

type TableGroup struct {
	Name          string `json:"name,omitempty" db:"name"`
	Slug          string `json:"slug,omitempty" db:"slug"`
	Description   string `json:"description,omitempty" db:"description"`
	SourceDb      string `json:"source_db,omitempty" db:"source_db"`
	TenantID      string `json:"tenant_id,omitempty" db:"tenant_id"`
	CabinetSource string `json:"cabinet_source,omitempty" db:"cabinet_source"`
	CabinetFolder string `json:"cabinet_folder,omitempty" db:"cabinet_folder"`
	Renderer      string `json:"renderer,omitempty" db:"renderer"`

	OwnedByPlug   string `json:"owned_by_plug,omitempty"  db:"owned_by_plug,omitempty"`
	MigrationHead string `json:"migration_head,omitempty"  db:"migration_head,omitempty"`

	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	Active    bool       `json:"active,omitempty" db:"active"`
}

type Table struct {
	Name         string     `json:"name,omitempty" db:"name"`
	Slug         string     `json:"slug,omitempty" db:"slug"`
	GroupID      string     `json:"group_id,omitempty" db:"group_id"`
	Description  string     `json:"description,omitempty" db:"description"`
	Icon         string     `json:"icon,omitempty" db:"icon"`
	MainColumn   string     `json:"main_column,omitempty" db:"main_column"`
	MainView     string     `json:"main_view,omitempty" db:"main_view"`
	ActivityType string     `json:"activity_type,omitempty" db:"activity_type"`
	SyncType     string     `json:"sync_type,omitempty" db:"sync_type"`
	TenantID     string     `json:"tenant_id,omitempty" db:"tenant_id"`
	ExtraMeta    JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type Column struct {
	Name          string         `json:"name,omitempty" db:"name"`
	Slug          string         `json:"slug,omitempty" db:"slug"`
	Ctype         string         `json:"ctype,omitempty" db:"ctype"`
	Description   string         `json:"description,omitempty" db:"description"`
	Icon          string         `json:"icon,omitempty" db:"icon"`
	Options       pq.StringArray `json:"options" db:"options"`
	OrderID       int64          `json:"order_id,omitempty" db:"order_id"`
	Pattern       string         `json:"pattern,omitempty" db:"pattern"`
	StrictPattern bool           `json:"strict_pattern,omitempty" db:"strict_pattern"`
	RefId         string         `json:"ref_id,omitempty" db:"ref_id"`
	RefType       string         `json:"ref_type,omitempty" db:"ref_type"`
	RefTarget     string         `json:"ref_target,omitempty" db:"ref_target"`
	RefObject     string         `json:"ref_object,omitempty" db:"ref_object"`
	RefCopy       string         `json:"ref_copy,omitempty" db:"ref_copy"`
	TableID       string         `json:"table_id,omitempty" db:"table_id"`
	GroupID       string         `json:"group_id,omitempty" db:"group_id"`
	TenantID      string         `json:"tenant_id,omitempty" db:"tenant_id"`
	ExtraMeta     JsonStrMap     `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type DTableCache struct {
	TableSlug string
	GroupSlug string
	Model     *Table
	Columns   map[string]*Column
	Hooks     []*TargetHook
	Views     []*DataView
}

type DCache interface {
	GetTableCache(tenantId, group, table string) (*DTableCache, error)
	EvictTable(tenantId, group, table string)
	GetColumnCache(tenantId, group, table string) (map[string]*Column, error)
	EvictColumns(tenantId, group, table string)
}

type Index struct {
	Mtype string   `json:"mtype,omitempty" yaml:"mtype,omitempty"`
	Slug  string   `json:"slug,omitempty" yaml:"slug,omitempty"`
	Spans []string `json:"spans" yaml:"spans"`
}

type FTSIndex struct {
	Type        string         `json:"type,omitempty" yaml:"type,omitempty"`
	Slug        string         `json:"slug,omitempty" yaml:"slug,omitempty"`
	ColumnSpans []string       `json:"spans" yaml:"spans"`
	Options     map[string]any `json:"options" yaml:"options"`
}

type ColumnFKRef struct {
	Slug     string   `json:"slug,omitempty" yaml:"slug,omitempty"`
	Type     string   `json:"type,omitempty" yaml:"type,omitempty"`
	Target   string   `json:"target,omitempty" yaml:"target,omitempty"`
	FromCols []string `json:"from_cols,omitempty" yaml:"from_cols,omitempty"`
	ToCols   []string `json:"to_cols,omitempty" yaml:"to_cols,omitempty"`
	RefCopy  string   `json:"ref_copy,omitempty" yaml:"ref_copy,omitempty"`
}
