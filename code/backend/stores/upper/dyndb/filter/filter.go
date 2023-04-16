package filter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

const (
	FilterEqual     = "equal"
	FilterNotEqual  = "not_equal"
	FilterIn        = "in"
	FilterNotIn     = "not_in"
	FilterNumIn     = "num_in"
	FilterNumNotIn  = "num_not_in"
	FilterLT        = "less_than"
	FilterGT        = "greater_than"
	FilterLTE       = "less_than_or_equal"
	FilterGTE       = "greater_than_or_equal"
	FilterAround    = "around"
	FilterNotAround = "not_around"
	FilterLike      = "like"
	FilterNotLike   = "not_like"
	FilterILike     = "insensitive_like"
	FilterNotILike  = "not_insensitive_like"
	FilterRgex      = "regex"
	FilterNotRgex   = "not_regex"
	FilterBefore    = "before"
	FilterAfter     = "after"

	FilterContains   = "contains"
	FilterHasPrefix  = "has_prefix"
	FilterHasSuffix  = "has_suffix"
	FilterIsNull     = "is_null"
	FilterIsNotNull  = "is_not_null"
	FilterBetween    = "between"
	FilterNotBetween = "not_between"

	OptrEqual      = ""
	OptrNotEqual   = " !="
	OptrIn         = " IN"
	OptrNotIn      = " NOT IN"
	OptrLT         = " <"
	OptrGT         = " >"
	OptrLTE        = " <="
	OptrGTE        = " >="
	OptrLIKE       = " LIKE"
	OptrNotLIKE    = " NOT LIKE"
	OptrILIKE      = " ILIKE"
	OptrNotILIKE   = " NOT ILIKE"
	OptrRegexp     = " REGEXP"
	OptrNotRegexp  = " NOT REGEXP"
	OptrAfter      = " >"
	OptrBefore     = " <"
	OptrIsNull     = " IS"
	OptrIsNotNull  = " IS NOT"
	OptrBetween    = " BETWEEN" // convert to appoprate type ? https://github.com/upper/db/blob/6d34eff2084ed3b148d5d7df13630a122180d347/comparison_test.go#L84
	OptrNotBetween = " NOT BETWEEN"
)

var (
	OptrMap = map[string]string{
		FilterEqual:      OptrEqual,
		FilterNotEqual:   OptrNotEqual,
		FilterIn:         OptrIn,
		FilterNotIn:      OptrNotIn,
		FilterLT:         OptrLT,
		FilterGT:         OptrGT,
		FilterLTE:        OptrLTE,
		FilterGTE:        OptrGTE,
		FilterLike:       OptrLIKE,
		FilterNotLike:    OptrNotLIKE,
		FilterILike:      OptrILIKE,
		FilterNotILike:   OptrNotILIKE,
		FilterRgex:       OptrRegexp,
		FilterNotRgex:    OptrNotRegexp,
		FilterBefore:     OptrBefore,
		FilterAfter:      OptrAfter,
		FilterIsNull:     OptrIsNull,
		FilterIsNotNull:  OptrIsNotNull,
		FilterBetween:    OptrBetween,
		FilterNotBetween: OptrNotBetween,
	}
)

func Transform(fcs []dyndb.FilterCond) (db.Cond, error) {
	return TransformWithPrefix(fcs, "")
}

func TransformWithPrefix(fcs []dyndb.FilterCond, prefix string) (db.Cond, error) {

	conds := make(db.Cond)

	for _, filter := range fcs {

		normalTransform := func() {
			optr, ok := OptrMap[filter.Cond]
			if !ok {
				return
			}
			conds[fmt.Sprintf("%s%s%s", prefix, filter.Column, optr)] = filter.Value
		}

		switch filter.Cond {
		case FilterAround, FilterNotAround:
			panic("location filter not implemented")
		case FilterIn, FilterNotIn:
			switch v := filter.Value.(type) {
			case string:
				filter.Value = strings.Split(v, ",")
			default:
				panic(fmt.Sprintf("In filter should be string, but found %t", v))
			}

			normalTransform()

		case FilterNumIn, FilterNumNotIn:
			if FilterNumIn == filter.Cond {
				filter.Cond = FilterIn
			} else {
				filter.Cond = FilterNotIn
			}

			switch v := filter.Value.(type) {
			case string:
				vals := strings.Split(v, ",")
				numvals := make([]float64, 0, len(vals))
				for _, _val := range vals {
					fnum, _ := strconv.ParseFloat(_val, 64)
					numvals = append(numvals, fnum)
				}

				filter.Value = numvals
			default:
				panic(fmt.Sprintf("Num In filter should be string, but found %t", v))
			}

			normalTransform()
		case FilterContains:
			filter.Cond = FilterILike
			filter.Value = fmt.Sprintf("%%%s%%", filter.Value)
			normalTransform()
		case FilterIsNull:
			filter.Value = nil
			normalTransform()
		case FilterHasPrefix:
			filter.Cond = FilterLike
			filter.Value = fmt.Sprintf("%%%s", filter.Value)
			normalTransform()
		case FilterHasSuffix:
			filter.Cond = FilterLike
			filter.Value = fmt.Sprintf("%s%%", filter.Value)
			normalTransform()
		default:
			normalTransform()
		}

	}

	return conds, nil

}

// location old code

/*

	fixme => sql_escape column ?

	data := locationData(filter.Value.(string))
	comOp := "<"
	if filter.Cond == FilterNotAround {
		comOp = ">"
	}

	resp = append(resp, db.Raw(fmt.Sprintf("ST_Distance(%s, ST_MakePoint(?, ?)::geography) %s", filter.Column, comOp), data.lat, data.long, data.distance))
*/
