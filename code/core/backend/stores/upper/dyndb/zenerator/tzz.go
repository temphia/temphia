package zenerator

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/stores/upper/dyndb/tns"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type tzz struct {
	tenantId        string
	gslug           string
	allSiblings     []string
	tns             tns.TNS
	gzz             *zenerator
	model           *bprints.NewTable
	tableSlug       string
	referecedTables []string
}

func (g *zenerator) TZZ(tenantId, gslug string, model *bprints.NewTable, sibling []string) *tzz {
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

	wctx.Write(TableHead(t.tableSlug))

	for _, col := range t.model.Columns {
		colstr := t.gzz._innerColumn(col.Slug, col.Ctype, col.NotNullable, "")
		wctx.Seperator()
		wctx.Write(colstr)
	}

	wctx.CondWriteCol(t.model.DeletedAt, t.gzz._innerColumn("deleted_at", store.CtypeDateTime, false, ""))

	// unique index
	for _, idx := range t.model.UniqueIndexes {
		wctx.Seperator()
		wctx.Write(Unique(idx.Spans))
	}

	// foreign key
	for _, fk := range t.model.ColumnRef {
		if len(fk.ToCols) == 0 {
			fk.ToCols = []string{store.KeyPrimary}
		}

		switch fk.Type {
		case store.RefHardPriId, store.RefHardText, store.RefHardMulti:
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

	tname := t.tableSlug

	if t.gzz.vendor == store.VendorPostgres {

		activityTable := t.tns.ActivityTable(t.tenantId, t.gslug, t.model.Slug)

		wctx.Write(
			fmt.Sprintf(`

			CREATE TABLE %s (
				id serial primary key,
				type TEXT NOT NULL,
				row_id integer not null,
				row_version integer not null,
				user_id text not null DEFAULT '',
				user_sign text not null DEFAULT '',
				init_sign text not null DEFAULT '',
				payload text not null DEFAULT '',
				message text not null DEFAULT '',
				created_at timestamptz not null default now()
			);
			
			`, activityTable,
			),
		)

		if t.model.ActivityType == store.DynActivityTypeStrict {
			wctx.Write(fmt.Sprintf(`CREATE TRIGGER 
			data_tg_%s AFTER INSERT OR UPDATE ON 
			%s FOR EACH ROW EXECUTE 
			FUNCTION data_activity_tg();`, tname, tname))
		}

	}

	return wctx.buffer.String(), nil
}

func (t *tzz) GetIndexes() []string {
	indexes := make([]string, 0)

	if t.model.FTSIndex != nil {
		// primary fts index
		istr, err := t.gzz.AddIndex(t.tenantId, t.gslug, t.model.Slug, "fts", store.IndexFTS, t.model.FTSIndex.ColumnSpans)
		if err != nil {
			panic(err.Error() + "fixme")
		}

		indexes = append(indexes, istr)
	}

	for _, idx := range t.model.Indexes {
		istr, err := t.gzz.AddIndex(t.tenantId, t.gslug, t.model.Slug, idx.Slug, store.IndexNormal, idx.Spans)
		if err != nil {
			panic(err.Error() + "fixme")
		}

		indexes = append(indexes, istr)

	}

	return indexes
}
