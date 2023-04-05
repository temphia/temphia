package hsql

import (
	"fmt"
	"strings"

	"github.com/rqlite/sql"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/tns"
	"github.com/thoas/go-funk"
)

var _ sql.Visitor = (*HsqlVisitor)(nil)

type HsqlVisitor struct {
	tenantId        string
	group           string
	tns             tns.TNS
	inverseAliasMap map[string]string
	allowedTables   []string
}

func (h *HsqlVisitor) Visit(node sql.Node) (w sql.Visitor, err error) {

	switch snode := node.(type) {

	case *sql.QualifiedTableName:
		name := snode.Name.Name

		if h.allowedTables != nil {
			if !funk.ContainsString(h.allowedTables, name) {
				return nil, easyerr.Error(fmt.Sprintf("Not allowed table access: %s", name))
			}
		}
		if snode.Alias != nil {
			h.inverseAliasMap[snode.Alias.Name] = name
		}

		snode.Name.Name = h.tns.Table(h.tenantId, h.group, name)

	case *sql.QualifiedRef:
		snode.Table.Name = h.tns.Table(h.tenantId, h.group, snode.Table.Name)
	case *sql.Call:
		name := snode.Name.Name
		if !funcWhiteList[strings.ToUpper(name)] {
			return nil, easyerr.Error(fmt.Sprintf("Not allowed func expr: %s", name))
		}

	}

	return h, nil
}

func (h *HsqlVisitor) VisitEnd(node sql.Node) error {
	return nil
}
