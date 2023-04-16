package filter

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

const (
	FilterEqual     = "equal"
	FilterNotEqual  = "not_equal"
	FilterIn        = "in"
	FilterNotIn     = "not_in"
	FilterLT        = "less_than"
	FilterGT        = "greater_than"
	FilterLTE       = "less_than_or_equal"
	FilterGTE       = "greater_than_or_equal"
	FilterAround    = "around"
	FilterNotAround = "not_around"

	OptrEqual    = ""
	OptrNotEqual = " !="
	OptrIn       = " IN"
	OptrNotIn    = " NOT IN"
	OptrLT       = " <"
	OptrGT       = " >"
	OptrLTE      = " <="
	OptrGTE      = " >="
)

/*
	filterIsNULL = "is_null"
	filterLIKE   = "like"
	ref(target) => join
*/

var (
	OptrMap = map[string]string{
		FilterEqual:    OptrEqual,
		FilterNotEqual: OptrNotEqual,
		FilterIn:       OptrIn,
		FilterNotIn:    OptrNotIn,
		FilterLT:       OptrLT,
		FilterGT:       OptrGT,
		FilterLTE:      OptrLTE,
		FilterGTE:      OptrGTE,
	}
)

func Transform(fcs []dyndb.FilterCond) (db.Cond, error) {
	return TransformWithPrefix(fcs, "")
}

func TransformWithPrefix(fcs []dyndb.FilterCond, prefix string) (db.Cond, error) {

	conds := make(db.Cond)

	for _, filter := range fcs {

		switch filter.Cond {
		case FilterAround, FilterNotAround:
			// fixme => sql_escape column ?

			/*
				data := locationData(filter.Value.(string))
				comOp := "<"
				if filter.Cond == FilterNotAround {
					comOp = ">"
				}

				resp = append(resp, db.Raw(fmt.Sprintf("ST_Distance(%s, ST_MakePoint(?, ?)::geography) %s", filter.Column, comOp), data.lat, data.long, data.distance))
			*/
		default:
			optr, ok := OptrMap[filter.Cond]
			if !ok {
				continue
			}

			conds[fmt.Sprintf("%s%s%s", prefix, filter.Column, optr)] = filter.Value
		}

	}

	return conds, nil

}
