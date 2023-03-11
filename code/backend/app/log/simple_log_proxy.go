package log

import (
	"bufio"
	"os"
	"strconv"

	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/tidwall/gjson"
)

type SimpleLogProxy struct {
	Path string
}

func (s *SimpleLogProxy) Query(from, to, tenantId string, filters map[string]string) ([]logx.Log, error) {
	return s.query("", from, to, tenantId, filters)
}

func (s *SimpleLogProxy) query(index, from, to, tenantId string, filters map[string]string) ([]logx.Log, error) {
	readFile, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	const max = 1000

	result := make([]logx.Log, 0, max)

lineLoop:
	for fileScanner.Scan() {
		line := fileScanner.Bytes()

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
