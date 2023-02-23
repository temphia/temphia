package main

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/filter"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func main() {

	pp.Println(goqu.From("test").Where(goqu.Ex{
		"d": []string{"a", "b", "c"},
	}).ToSQL())

	pp.Println("@t")

	transform([][]dyndb.FilterCond{
		{
			dyndb.FilterCond{
				Column: "aaa",
				Cond:   "eq",
				Value:  "111",
			},

			dyndb.FilterCond{
				Column: "bb",
				Cond:   "eq",
				Value:  "111",
			},
		},

		{
			dyndb.FilterCond{
				Column: "ccc",
				Cond:   "eq",
				Value:  "43333",
			},
		},
	})

}

func transform(filters [][]dyndb.FilterCond) {

	exprs := make([]goqu.Expression, 0)

	for _, qry := range filters {
		inner := goqu.ExOr{}

		for _, fc := range qry {

			switch fc.Cond {
			case filter.OptrGT:
				inner[fc.Column] = goqu.Op{"gt": fc.Value}
			default:
				inner[fc.Column] = goqu.Op{"eq": fc.Value}
			}
		}
		exprs = append(exprs, inner)
	}

	pp.Println(goqu.Select("c1").Where(exprs...).ToSQL())

}
