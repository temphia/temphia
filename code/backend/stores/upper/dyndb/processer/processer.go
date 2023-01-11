package processer

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type Processer interface {
	FromRowsDBType(rows []map[string]interface{}) error
	FromRowDBType(row map[string]interface{}) error
	ToRowDBType(row map[string]interface{}) error
	ToRowsDBType(rows []map[string]interface{}) error
}

func New(vendor string, columns map[string]*entities.Column) Processer {
	pg := &PGCtypeProcesser{
		columns: columns,
	}

	return pg
}
