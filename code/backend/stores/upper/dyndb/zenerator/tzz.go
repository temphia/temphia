package zenerator

import (
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/tns"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type tzz struct {
	tenantId        string
	gslug           string
	allSiblings     []string
	tns             tns.TNS
	gzz             *zenerator
	model           *xbprint.NewTable
	tableSlug       string
	referecedTables []string
}

func (g *zenerator) TZZ(tenantId, gslug string, model *xbprint.NewTable, sibling []string) *tzz {
	return &tzz{
		tenantId:        tenantId,
		gslug:           gslug,
		allSiblings:     sibling,
		tns:             g.tns,
		gzz:             g,
		model:           model,
		tableSlug:       g.tns.Table(tenantId, gslug, model.Slug),
		referecedTables: []string{},
	}
}

func (t *tzz) CreateTable() (string, error) {

	wctx := WriterCtx{}

	switch t.gzz.vendor {
	case store.VendorSqlite:
		wctx.Write(TableHeadSqlite(t.tableSlug))
	default:
		wctx.Write(TableHead(t.tableSlug))
	}

	for _, col := range t.model.Columns {
		colstr := t.gzz._innerColumn(col.Slug, col.Ctype, col.NotNullable, "")
		wctx.Seperator()
		wctx.Write(colstr)
	}

	// unique index
	for _, idx := range t.model.UniqueIndexes {
		wctx.Seperator()
		wctx.Write(Unique(idx.Spans))
	}

	// foreign key
	for _, fk := range t.model.ColumnRef {
		if len(fk.ToCols) == 0 {
			fk.ToCols = []string{dyndb.KeyPrimary}
		}

		switch fk.Type {
		case dyndb.RefHardPriId, dyndb.RefHardText, dyndb.RefHardMulti:
			t.referecedTables = append(t.referecedTables, fk.Target)
			wctx.Seperator()
			wctx.Write(
				InnerFKRef(t.tns.Table(t.tenantId, t.gslug, fk.Target), fk.FromCols, fk.ToCols),
			)
		default:
			continue
		}
	}

	wctx.Terminate()

	switch t.gzz.vendor {
	case store.VendorPostgres:
		t.activityTablePg(&wctx)
	case store.VendorSqlite:
		t.activityTableSqlite(&wctx)
	default:
		panic("Invalid verndor " + t.gzz.vendor)
	}

	return wctx.buffer.String(), nil
}

func (t *tzz) GetIndexes() []string {
	indexes := make([]string, 0)

	if t.model.FTSIndex != nil {
		// primary fts index
		istr, err := t.gzz.AddIndex(t.tenantId, t.gslug, t.model.Slug, "fts", dyndb.IndexFTS, t.model.FTSIndex.ColumnSpans)
		if err != nil {
			panic(err.Error() + "fixme")
		}

		indexes = append(indexes, istr)
	}

	for _, idx := range t.model.Indexes {
		istr, err := t.gzz.AddIndex(t.tenantId, t.gslug, t.model.Slug, idx.Slug, dyndb.IndexNormal, idx.Spans)
		if err != nil {
			panic(err.Error() + "fixme")
		}

		indexes = append(indexes, istr)

	}

	return indexes
}
