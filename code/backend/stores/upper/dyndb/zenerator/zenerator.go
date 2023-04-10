package zenerator

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/tns"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type zenerator struct {
	vendor       string
	tns          tns.TNS
	_innerColumn func(cslug, ctype string, notnull bool, defval string) string
	_index       func(tblname, iname, itype string, spans []string) string
}

var _ ucore.Zenerator = (*zenerator)(nil)

func New(vendor string, tns tns.TNS) *zenerator {
	z := &zenerator{
		vendor:       vendor,
		tns:          tns,
		_innerColumn: nil,
		_index:       nil,
	}

	switch vendor {
	case store.VendorPostgres:
		z._index = z.indexPg
		z._innerColumn = z.innerColumnPg(pgCtypeMap)
	case store.VendorSqlite:
		z._index = z.indexSqlite
		z._innerColumn = z.innerColumnPg(sqliteCtypeMap)

	default:
		panic("not supported vendor:" + vendor)
	}

	return z
}

func (g *zenerator) NewGroup(tenantId string, model *xbprint.NewTableGroup) (*ucore.DDLGroupStmt, error) {

	if err := g.tns.CheckGroupSlug(model.Slug); err != nil {
		return nil, err
	}

	gstmt := &ucore.DDLGroupStmt{
		TableStmts:      make(map[string]string),
		FRefs:           make(map[string][]string),
		TableIndexStmts: make(map[string][]string),
		GroupSlug:       model.Slug,
	}

	siblings := make([]string, 0, len(model.Tables))

	for _, tbl := range model.Tables {
		siblings = append(siblings, tbl.Slug)
	}

	for _, table := range model.Tables {
		stmt, err := g.NewTable(tenantId, model.Slug, table, siblings)
		if err != nil {
			return nil, err
		}

		gstmt.TableStmts[table.Slug] = stmt.Stmt
		gstmt.FRefs[table.Slug] = stmt.FRefs
		gstmt.TableIndexStmts[table.Slug] = stmt.TableIndexStmts

	}

	return gstmt, nil
}

func (g *zenerator) NewTable(tenantId, gslug string, model *xbprint.NewTable, siblings []string) (*ucore.DDLStmt, error) {

	tgen := g.newTZZ(tenantId, gslug, model, siblings) // table generator
	str, err := tgen.CreateTable()
	if err != nil {
		return nil, err
	}

	stmt := &ucore.DDLStmt{
		Stmt:            str,
		FRefs:           tgen.referecedTables,
		TableIndexStmts: tgen.GetIndexes(),
	}

	return stmt, nil
}

func (g *zenerator) DropTable(tenantId, gslug, tslug string) (string, error) {
	tbl := g.tns.Table(tenantId, gslug, tslug)
	atbl := g.tns.ActivityTable(tenantId, gslug, tslug)

	return fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
		DROP TABLE IF EXISTS %s CASCADE;
	`, tbl, atbl), nil
}

func (g *zenerator) AddColumn(tenantId, gslug, tslug, cslug string, model *xbprint.NewColumn) (string, error) {
	return g.addColumn(tenantId, gslug, tslug, cslug, model.Ctype, model.NotNullable)
}

func (g *zenerator) addColumn(tenantId, gslug, tslug, cname, ctype string, notnull bool) (string, error) {
	tbl := g.tns.Table(tenantId, gslug, tslug)
	return fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s;", tbl, g._innerColumn(cname, ctype, notnull, "")), nil
}

func (g *zenerator) DropColumn(tenantId, gslug, tslug, cname string) (string, error) {
	tbl := g.tns.Table(tenantId, gslug, tslug)
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s;", tbl, cname), nil
}

func (g *zenerator) AddIndex(tenantId, gslug, tslug, iname, itype string, spans []string) (string, error) {
	tbl := g.tns.Table(tenantId, gslug, tslug)
	return g._index(tbl, iname, itype, spans), nil
}

func (g *zenerator) AddFKRef(tenantId, gslug, tslug, target string, from []string, to []string) (string, error) {
	tbl := g.tns.Table(tenantId, gslug, tslug)
	return fmt.Sprintf(`ALTER TABLE ADD CONSTRAINT %s %s;`, tbl, innerFKRef(target, from, to)), nil
}

func (g *zenerator) RemoveFKRef(tenantId, gslug, tslug string, fkslug string) (string, error) {
	tbl := g.tns.Table(tenantId, gslug, tslug)
	return fmt.Sprintf("ALTER TABLE %s DROP FOREIGN KEY %s;", tbl, fkslug), nil
}

func (g *zenerator) RemoveIndex(tenantId, gslug, tslug, islug string) (string, error) {
	return fmt.Sprintf("DROP INDEX %s;", islug), nil
}

func (g *zenerator) GetIndexs(tenantId, gslug, tslug string, fn func(query string) (map[string]interface{}, error)) ([]*entities.Index, error) {
	return nil, easyerr.NotImpl()
}

func (g *zenerator) GetFKRefs(tenantId, gslug, tslug string, fn func(query string) (map[string]interface{}, error)) ([]*entities.Index, error) {
	return nil, easyerr.NotImpl()
}
