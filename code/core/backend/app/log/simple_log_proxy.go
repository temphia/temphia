package log

import (
	"bufio"
	"os"

	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
	"github.com/tidwall/gjson"
)

type SimpleLogProxy struct {
	Path string
}

func (s *SimpleLogProxy) QueryAppTenant(from, to, tenantId string, filters map[string]string) ([]logx.Log, error) {
	return s.query("app", from, to, tenantId, filters)
}

func (s *SimpleLogProxy) QueryEngine(from, to, tenantId string, filters map[string]string) ([]logx.Log, error) {
	return s.query("engine", from, to, tenantId, filters)
}

func (s *SimpleLogProxy) QuerySite(from, to, tenantId string, filters map[string]string) ([]logx.Log, error) {
	return s.query("site", from, to, tenantId, filters)
}

func (s *SimpleLogProxy) query(index, from, to, tenantId string, filters map[string]string) ([]logx.Log, error) {
	readFile, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	const max = 100

	result := make([]logx.Log, 0, max)
	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		out := gjson.GetBytes(line, "index")
		if out.String() != index {
			continue
		}
		result = append(result, logx.Log(line))

		if len(result) == max {
			break
		}

	}

	readFile.Close()

	return result, nil
}
