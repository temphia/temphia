package dyndb

import (
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dynddl2"
)

func (d *DynDB) performNewMigrate(tenantId string, migctx dynddl2.MigrateContext) error {

	// applyer := applyNewGroup{
	// 	session:    d.session, // fixme use txt
	// 	sharedLock: d.sharedLock,
	// 	logger:     zerolog.New(os.Stdout),
	// }

	return nil

}
