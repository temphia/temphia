package dynddl2

import (
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

type (
	MigrateContext struct {
		BaseSchema  *xbprint.NewTableGroup
		StmtString  string
		PostItems   []PostDDLItem
		Siblings    map[string]map[string]string
		Options     step.MigrateOptions
		LastMigHead string
		NextMigHead string
	}

	PostDDLItem struct {
		Mtype string
		Data  any
	}
)
