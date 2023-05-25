package step

type MigrateOptions struct {
	Steps []Step `json:"steps,omitempty" yaml:"steps,omitempty"`
	New   bool   `json:"new,omitempty" yaml:"new,omitempty"`
	Slug  string `json:"slug,omitempty" yaml:"slug,omitempty"`
}

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
