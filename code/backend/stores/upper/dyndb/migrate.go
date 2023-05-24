package dyndb

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (d *DynDB) migrateSchema(tenantId string, opts dyndb.MigrateOptions) error {
	if opts.New {
		return d.newMigrateSchema(tenantId, opts)
	}

	for _, mstep := range opts.Steps {

		switch mstep.Type {
		case dyndb.MigTypeNewGroup:
			schema := &xbprint.NewTableGroup{}

			stmt, err := d.dyngen.NewGroup(tenantId, schema)
			if err != nil {
				return err
			}

			pp.Println(stmt.String())

		default:

		}

	}

	return nil
}

func (d *DynDB) newMigrateSchema(tenantId string, opts dyndb.MigrateOptions) error {

	return nil
}
