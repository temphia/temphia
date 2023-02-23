package filter

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

const (
	filterEqual     = "equal"
	filterNotEqual  = "not_equal"
	filterIn        = "in"
	filterNotIn     = "not_in"
	filterLT        = "less_than"
	filterGT        = "greater_than"
	filterLTE       = "less_than_or_equal"
	filterGTE       = "greater_than_or_equal"
	filterAround    = "around"
	filterNotAround = "not_around"

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
		filterEqual:    OptrEqual,
		filterNotEqual: OptrNotEqual,
		filterIn:       OptrIn,
		filterNotIn:    OptrNotIn,
		filterLT:       OptrLT,
		filterGT:       OptrGT,
		filterLTE:      OptrLTE,
		filterGTE:      OptrGTE,
	}
)

func Transform(fcs []dyndb.FilterCond) (interface{}, error) {

	conds := make(db.Cond)
	resp := []interface{}{
		conds,
	}

	for _, filter := range fcs {

		switch filter.Cond {
		case filterAround, filterNotAround:
			// fixme => sql_escape column ?
			data := locationData(filter.Value.(string))
			comOp := "<"
			if filter.Cond == filterNotAround {
				comOp = ">"
			}

			resp = append(resp, db.Raw(fmt.Sprintf("ST_Distance(%s, ST_MakePoint(?, ?)::geography) %s", filter.Column, comOp), data.lat, data.long, data.distance))
		default:
			optr, ok := OptrMap[filter.Cond]
			if !ok {
				return conds, nil
			}
			conds[filter.Column+optr] = filter.Value
		}

	}

	return conds, nil
}
