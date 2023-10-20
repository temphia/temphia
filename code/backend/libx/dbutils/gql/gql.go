package gql

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type Table struct {
	Columns map[string]*entities.Column
	Inner   *entities.Table
}

type GQL struct {
	Objects map[string]*graphql.Object
	Group   entities.TableGroup
	Tables  map[string]*Table
	schema  *graphql.Schema
}

func New(group entities.TableGroup, tables map[string]*Table) (*GQL, error) {

	g := &GQL{
		Objects: make(map[string]*graphql.Object),
		Group:   group,
		Tables:  tables,
		schema:  nil,
	}

	err := g.build()
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *GQL) build() error {

	ffs := graphql.Fields{}

	for _, t := range g.Tables {

		sfields := graphql.Fields{}

		for k, c := range t.Columns {
			var ctype graphql.Output

			switch c.Ctype {

			case dyndb.CtypeShortText, dyndb.CtypeEmail,
				dyndb.CtypeLongText, dyndb.CtypeColor,
				dyndb.CtypePhone, dyndb.CtypeSelect,
				dyndb.CtypeFile, dyndb.CtypeMultiFile,
				dyndb.CtypeMultSelect, dyndb.CtypeSingleUser,
				dyndb.CtypeMultiUser, dyndb.CtypeJSON:

				ctype = graphql.String
			case dyndb.CtypeCurrency, dyndb.CtypeNumber, dyndb.CtypeRangeNumber:
				ctype = graphql.Float
			case dyndb.CtypeLocation:
				ctype = graphql.String
			case dyndb.CtypeDateTime:
				ctype = graphql.DateTime
			case dyndb.CtypeCheckBox:
				ctype = graphql.Boolean
			default:
				panic("ctype to graphql type notimplemented")
			}

			sfields[k] = &graphql.Field{
				Name:    k,
				Type:    ctype,
				Resolve: nil,
			}
		}

		tfield := &graphql.Field{
			Name: t.Inner.Name,
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name:   t.Inner.Name,
				Fields: sfields,
			}),
		}

		ffs[t.Inner.Name] = tfield

	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: ffs,
		}),
	})
	if err != nil {
		return nil
	}

	g.schema = &schema

	return nil
}

func (g *GQL) Query(ctx context.Context, q string) (any, error) {

	return nil, nil
}
