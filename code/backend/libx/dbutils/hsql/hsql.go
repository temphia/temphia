package hsql

import (
	"regexp"
	"strings"

	"github.com/rqlite/sql"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/tns"
)

type Result struct {
	InverseAlias     map[string]string
	TransformedQuery string
}

type Hsql struct {
	tns tns.TNS
}

func New(tns tns.TNS) *Hsql {
	return &Hsql{
		tns: tns,
	}
}

func (h *Hsql) Transform(tenantId, group string, allowedTables []string, query string) (*Result, error) {
	query = removeSQLComments(query)

	parser := sql.NewParser(strings.NewReader(query))
	stmt, err := parser.ParseStatement()
	if err != nil {
		return nil, err
	}

	switch stmt.(type) {
	case *sql.SelectStatement:
	default:

		return nil, easyerr.Error("invalid statement type")
	}

	v := &HsqlVisitor{
		tenantId:        tenantId,
		group:           group,
		tns:             h.tns,
		allowedTables:   allowedTables,
		inverseAliasMap: make(map[string]string),
	}

	err = (sql.Walk(v, stmt))
	if err != nil {
		return nil, err
	}

	return &Result{
		InverseAlias:     v.inverseAliasMap,
		TransformedQuery: stmt.String(),
	}, nil
}

func removeSQLComments(query string) string {
	commentRegex := regexp.MustCompile(`(--|#|/\*).*?(\*/|\n)`)
	return commentRegex.ReplaceAllString(query, "")
}
