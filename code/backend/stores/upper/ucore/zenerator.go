package ucore

import (
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

type Zenerator interface {
	NewGroup(tenantId string, model *xbprint.NewTableGroup) (*DDLGroupStmt, error)
	NewTable(tenantId, gslug string, model *xbprint.NewTable, siblings []string) (*DDLStmt, error)
	DropTable(tenantId, gslug, tslug string) (string, error)

	AddColumn(tenantId, gslug, tslug, cslug string, model *xbprint.NewColumn) (string, error)
	DropColumn(tenantId, gslug, tslug, cslug string) (string, error)

	AddIndex(tenantId, gslug, tslug, iname, itype string, spans []string) (string, error)
	AddFKRef(tenantId, gslug, tslug, target string, to []string, from []string) (string, error)
	RemoveFKRef(tenantId, gslug, tslug, fkslug string) (string, error)
	RemoveIndex(tenantId, gslug, tslug, islug string) (string, error)

	GetIndexs(tenantId, gslug, tslug string, fn func(query string) (map[string]interface{}, error)) ([]*entities.Index, error)
	GetFKRefs(tenantId, gslug, tslug string, fn func(query string) (map[string]interface{}, error)) ([]*entities.Index, error)
}

type DDLGroupStmt struct {
	TableStmts      map[string]string
	FRefs           map[string][]string
	TableIndexStmts map[string][]string
	GroupSlug       string
}

// this concats ddl stmt of table using proper ordering
// bashed on which table has foreign reference to another
// t2 references t1 then first t1 should be created
func (d *DDLGroupStmt) String() string {
	remaining := make(map[string]struct{}, len(d.TableStmts))
	for k := range d.TableStmts {
		remaining[k] = struct{}{}
	}

	buf := strings.Builder{}
	write := func(tbl string) {
		stmt := d.TableStmts[tbl]
		buf.WriteString(stmt)

		for _, indexstmt := range d.TableIndexStmts[tbl] {
			buf.WriteString(indexstmt)
		}
		delete(remaining, tbl)
	}

	i := 0
	for {

		if len(remaining) == 0 {
			break
		}
		i = i + 1
		if i == 50 {
			break
		}

		for k := range d.TableStmts {
			if _, ok := remaining[k]; !ok {
				continue
			}

			frefs, ok := d.FRefs[k]
			if !ok {
				write(k)
				continue
			}

			allRefDone := true
			for _, i := range frefs {
				if _, remain := remaining[i]; remain {
					allRefDone = false
				}
			}

			if allRefDone {
				write(k)
				continue
			}

		}

	}

	return buf.String()

}

type DDLStmt struct {
	Stmt            string
	FRefs           []string
	TableIndexStmts []string
}

func (d *DDLStmt) String() string {
	buf := strings.Builder{}
	buf.WriteString(d.Stmt)
	for _, indexstmt := range d.TableIndexStmts {
		buf.WriteString(indexstmt)
	}
	return buf.String()
}
