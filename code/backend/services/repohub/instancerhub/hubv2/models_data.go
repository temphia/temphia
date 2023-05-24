package hubv2

import "github.com/temphia/temphia/code/backend/xtypes/store/dyndb"

type DataSchemaV2 struct {
	Steps []dyndb.MigrationStep `json:"steps,omitempty" yaml:"steps,omitempty"`
}
