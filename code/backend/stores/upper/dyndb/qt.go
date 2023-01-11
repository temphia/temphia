package dyndb

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/store"
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

func transformFilters(fcs []*store.FilterCond) (interface{}, error) {

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

type LData struct {
	lat      float32
	long     float32
	distance float32
}

func locationData(rawstr string) LData {
	points := strings.Split(rawstr, " ")

	return LData{
		lat:      getFloat(points[0]),
		long:     getFloat(points[1]),
		distance: getFloat(points[2]),
	}
}

func getFloat(fstr string) float32 {
	f, err := strconv.ParseFloat(fstr, 32)
	if err != nil {
		panic(err)
	}
	return float32(f)
}
