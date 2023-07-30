package dtable

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upperdb/dyndb/tns"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func Validate(schema *xbprint.NewTableGroup) error {
	tns := tns.New("shared")
	err := tns.CheckGroupSlug(schema.Slug)
	if err != nil {
		return err
	}

	tables := make(map[string]bool)

	for _, table := range schema.Tables {
		err = tns.CheckTableSlug(table.Slug)
		if err != nil {
			return err
		}

		if tables[table.Slug] {
			return easyerr.Error(fmt.Sprintf("%s deuplicate table: %s", table.Name, table.Slug))
		} else {
			tables[table.Slug] = true
		}

		columns := make(map[string]bool)

		for _, col := range table.Columns {
			if columns[col.Slug] {
				return easyerr.Error(fmt.Sprintf("%s/%s deuplicate column slug: %s", table.Name, col.Name, col.Slug))
			} else {
				columns[col.Slug] = true
			}

			err = tns.CheckColumnSlug(col.Slug)
			if err != nil {
				return err
			}

			_, ok := dyndb.AllColumns[col.Ctype]
			if !ok {
				return easyerr.Error(fmt.Sprintf("%s has unknown column type %s", col.Name, col.Ctype))
			}

		}

	}

	return nil
}
