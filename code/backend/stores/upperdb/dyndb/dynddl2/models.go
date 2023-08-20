package dynddl2

import (
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

type (
	MigrateContext struct {
		BaseSchema  *xpackage.NewTableGroup
		StmtString  string
		PostItems   []PostDDLItem
		Siblings    map[string]map[string]string
		Options     xpackage.MigrateOptions
		Gslug       string
		LastMigHead string
		NextMigHead string
	}

	PostDDLItem struct {
		Name  string
		Mtype string
		Data  any
		Stmt  string
	}
)
