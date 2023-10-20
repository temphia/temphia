package gql

import (
	"context"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func TestGGL(t *testing.T) {

	g, err := New(entities.TableGroup{
		Name:        "Test",
		Slug:        "test",
		Description: "this is test",
	}, map[string]*Table{
		"products": {
			Inner: &entities.Table{
				Name: "Products",
				Slug: "product",
			},
			Columns: map[string]*entities.Column{
				"name": {
					Name:        "Name",
					Slug:        "name",
					Ctype:       "shorttext",
					Description: "Name of product",
				},

				"count": {
					Name:        "Count",
					Slug:        "count",
					Ctype:       "number",
					Description: "No of product in inventry",
				},
				"price": {
					Name:        "Price",
					Slug:        "price",
					Ctype:       "number",
					Description: "Prince of product",
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	pp.Println(g.Query(context.Background(), `
	
		
	
	`))

}
