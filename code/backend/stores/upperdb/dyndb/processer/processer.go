package processer

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Processer interface {
	FromRowsDBType(rows []map[string]interface{}) error
	FromRowDBType(row map[string]interface{}) error
	ToRowDBType(row map[string]interface{}) error
	ToRowsDBType(rows []map[string]interface{}) error
}

func New(vendor string, columns map[string]*entities.Column) Processer {

	switch vendor {
	case store.VendorPostgres:
		return &PGCtypeProcesser{
			columns: columns,
		}

	case store.VendorSqlite:
		return &SqliteCtypeProcesser{
			columns: columns,
		}
	default:
		panic("Invalid Vendor" + vendor)
	}
}
