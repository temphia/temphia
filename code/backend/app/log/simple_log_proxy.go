package log

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/tidwall/gjson"
)

type SimpleLogProxy struct {
	Path string
}

func (s *SimpleLogProxy) Query(from, to, cursor, tenantId string, filters map[string]string) ([]logx.Log, error) {
	return s.query(from, to, "", tenantId, filters)
}

func (s *SimpleLogProxy) query(from, to, cursor, tenantId string, filters map[string]string) ([]logx.Log, error) {
	readFile, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	const max = 1000

	// parse to,from time
	var (
		fromTime *time.Time
		toTime   *time.Time
	)

	err = parseTime(from, to, fromTime, toTime)
	if err != nil {
		return nil, err
	}

	// parse cursor
	var cxid xid.ID
	if cursor != "" {
		cxid, err = xid.FromString(cursor)
		if err != nil {
			return nil, easyerr.Wrap("could not parse cursor", err)
		}
	}

	result := make([]logx.Log, 0, max)

lineLoop:
	for fileScanner.Scan() {
		line := fileScanner.Bytes()

		curXid, err := xid.FromString(gjson.GetBytes(line, "log_event_id").String())
		if err != nil {
			return nil, easyerr.Wrap("err extracting log_event_id for checking cursor", err)
		}

		// cursor check
		if cursor != "" {

			if cxid.Compare(curXid) == -1 {
				continue
			}
		}

		// to/from time check here
		ltime := curXid.Time()
		if fromTime != nil {
			if fromTime.Before(ltime) {
				continue
			}
		}

		if toTime != nil {
			if toTime.After(ltime) {
				break
			}
		}

		if tenantId != "" {
			if gjson.GetBytes(line, "tenant_id").String() != tenantId {
				continue
			}
		}
	fiterLoop:
		for filterKey, filterValue := range filters {
			out := gjson.GetBytes(line, filterKey)

			switch out.Type {
			case gjson.Null:
				continue lineLoop
			case gjson.String:
				if out.String() == filterValue {
					continue fiterLoop
				}

				continue lineLoop

			case gjson.True:
				if filterValue == "true" {
					continue fiterLoop
				}

				continue lineLoop
			case gjson.False:
				if filterValue == "true" {
					continue fiterLoop
				}
				continue lineLoop
			case gjson.Number:
				fval, err := strconv.ParseFloat(filterValue, 64)
				if err != nil {
					continue lineLoop
				}

				if out.Float() == fval {
					continue fiterLoop
				}
				continue lineLoop
			default:
				continue lineLoop
			}

		}

		result = append(result, logx.Log(line))

		if len(result) == max {
			break
		}

	}

	readFile.Close()

	return result, nil
}

func parseTime(from, to string, fromTime *time.Time, toTime *time.Time) error {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"

	if from != "" {
		_fttime, err := time.Parse(layout, from)
		if err != nil {
			return easyerr.Wrap("could not parse from time error", err)
		}

		*fromTime = _fttime
	}

	if to != "" {
		_ttime, err := time.Parse(layout, to)
		if err != nil {
			return easyerr.Wrap("could not parse error", err)
		}

		*toTime = _ttime
	}

	return nil

}
