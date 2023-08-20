package dyndb

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

type DCache interface {
	CachedColumns(tenantId, group, table string) (map[string]*entities.Column, error)
	EvictColumns(tenantId, group, table string)
}

func ExtractColumns(m *xpackage.NewTable, tenantId, gslug string) []*entities.Column {
	indexedCol := make(map[string]*entities.Column)

	for _, nc := range m.Columns {
		newcol := nc.To(tenantId, gslug, m.Slug)
		indexedCol[nc.Slug] = newcol

	}

	for _, colref := range m.ColumnRef {
		if colref.Slug == "" {
			colref.Slug = xid.New().String()
		}

		if colref.Type == RefHardPriId || colref.Type == RefSoftPriId {
			if len(colref.ToCols) == 0 {
				colref.ToCols = []string{KeyPrimary}
			}
		}

		for idx, colId := range colref.FromCols {
			col := indexedCol[colId]
			col.RefId = colref.Slug
			col.RefType = colref.Type
			col.RefTarget = colref.Target
			col.RefObject = colref.ToCols[idx]
			col.RefCopy = colref.RefCopy
		}
	}

	cols := make([]*entities.Column, 0, len(m.Columns)+len(m.ColumnRef))
	for _, v := range indexedCol {
		cols = append(cols, v)
	}

	return cols

}

type (
	Schema struct {
		Group   *entities.TableGroup
		Tables  map[string]*entities.Table
		Columns map[string]*entities.Column
	}
)

func (s *Schema) AddColumn(cols ...*entities.Column) {
	for _, col := range cols {
		s.Columns[col.TableID+col.Slug] = col
	}
}

func (s *Schema) GetColumn(table, column string) *entities.Column {
	return s.Columns[table+column]
}
