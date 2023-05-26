package dyncore

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/upper/db/v4"
)

func GroupTable(sess db.Session) db.Collection {
	return sess.Collection("data_table_groups")
}

func Table(sess db.Session) db.Collection {
	return sess.Collection("data_tables")
}

func TableColumn(sess db.Session) db.Collection {
	return sess.Collection("data_table_columns")
}

type DynMeta interface {
	NewGroupMeta(tenantId string, model *xbprint.NewTableGroup) error
	RollbackGroupMeta(tenantId, gslug string)

	NewTableMeta(tenantId, gslug string, model *xbprint.NewTable) error
	RollbackTableMeta(tenantId, gslug, tslug string)

	NewColumnMeta(tenantId, gslug, tslug string, model *entities.Column) error
	RollbackColumnMeta(tenantId, gslug, tslug, cslug string)
}
