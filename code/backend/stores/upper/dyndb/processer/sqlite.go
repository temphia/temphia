package processer

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type SqliteCtypeProcesser struct {
	columns map[string]*entities.Column
}

func (scp *SqliteCtypeProcesser) FromRowsDBType(rows []map[string]interface{}) error { return nil }
func (scp *SqliteCtypeProcesser) FromRowDBType(row map[string]interface{}) error     { return nil }
func (scp *SqliteCtypeProcesser) ToRowDBType(row map[string]interface{}) error       { return nil }
func (scp *SqliteCtypeProcesser) ToRowsDBType(rows []map[string]interface{}) error   { return nil }
