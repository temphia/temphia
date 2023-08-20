package xinstancer

import (
	"encoding/json"
)

// agent

const (
	PlugStepNewAgent      = "new_agent"
	PlugStepUpdateAgent   = "update_agent"
	PlugStepRemoveAgent   = "remove_agent"
	PlugStepAddInnerLink  = "add_inner_link"
	PlugStepAddRemoveLink = "remove_inner_link"
)

// resource

const (
	PlugStepNewResourceModule    = "new_resource"
	PlugStepUpdateResourceModule = "update_resource"
	PlugStepRemoveResourceModule = "remove_resource"

	PlugStepAddResourceLink    = "add_resource_link"
	PlugStepRemoveResourceLink = "remove_resource_link"
)

// target

const (
	PlugStepAddTargetApp    = "add_target_app"
	PlugStepUpdateTargetApp = "update_target_app"
	PlugStepDeleteTargetApp = "delete_target_app"

	PlugStepAddTargetHook    = "add_target_hook"
	PlugStepUpdateTargetHook = "update_target_hook"
	PlugStepDeleteTargetHook = "delete_target_hook"
)

// data

type MigrateOptions struct {
	Steps  []Step `json:"steps,omitempty" yaml:"steps,omitempty"`
	New    bool   `json:"new,omitempty" yaml:"new,omitempty"`
	Gslug  string `json:"gslug,omitempty" yaml:"gslug,omitempty"`
	PlugId string `json:"plug_id,omitempty" yaml:"plug_id,omitempty"`
	DryRun bool   `json:"-"`
}

type Step struct {
	Name string          `json:"name,omitempty" yaml:"name,omitempty"`
	Type string          `json:"type,omitempty" yaml:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty" yaml:"data,omitempty"`
}

const (
	PlugStepRunDataMigration = "run_data_migration"
)

const (
	MigTypeNewGroup              = "new_group"
	MigTypeAddTable              = "add_table"
	MigTypeRemoveTable           = "remove_table"
	MigTypeAddColumn             = "add_column"
	MigTypeRemoveColumn          = "remove_column"
	MigTypeAddUniqueConstrain    = "add_unique_constrain"
	MigTypeRemoveUniqueConstrain = "remove_unique_constrain"
	MigTypeAddConstrain          = "add_constrain"
	MigTypeRemoveConstrain       = "remove_constrain"
	MigTypeAddForeignKey         = "add_foreign_key"
	MigTypeRemoveForeignKey      = "remove_foreign_key"
	MigTypeAddIndex              = "add_index"
	MigTypeRemoveIndex           = "remove_index"
	MigTypeAddFTSIndex           = "add_fts_index"
	MigTypeRemoveFTSIndex        = "remove_fts_index"
	MigTypeAddEmbedIndex         = "add_embed_index"
	MigTypeRemoveEmbedIndex      = "remove_embed_index"

	MigTypeSoftColTypeChange   = "soft_coltype_change"
	MigTypeSoftColOptsUpdate   = "soft_colopts_update"
	MigTypeSoftTableOptsChange = "soft_tableopts_update"

	MigTypeAddView    = "add_view"
	MigTypeRemoveView = "remove_view"
	MigTypeUpdateView = "update_view"

	MigTypeTransformData = "transform_data"
)
