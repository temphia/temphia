package lreader

import (
	"encoding/json"
	"testing"

	"github.com/temphia/temphia/code/backend/xtypes/logx"
)

type Log struct {
	Level  string  `json:"level,omitempty"`
	MetaId float64 `json:"meta_id,omitempty"`
}

func TestLreader(t *testing.T) {

	sl := Lreader{
		path: "testdata.log",
	}

	sl.Query("", logx.QueryRequest{
		Filters: map[string]string{
			"level": "debug",
		},
	})

	logs, err := sl.Query("", logx.QueryRequest{
		Filters: map[string]string{
			"level":   "debug",
			"meta_id": "42",
		},
	})

	if err != nil {
		t.Error(err)
	}

	if len(logs) != 1 {
		t.Fatal("expected one result found ", len(logs))
	}

	rLog := &Log{}

	err = json.Unmarshal([]byte(logs[0]), rLog)
	if err != nil {
		t.Error(err)
	}

	if rLog.Level != "debug" {
		t.Fatal("expected log level debug from output")
	}

	if rLog.MetaId != 42 {
		t.Fatal("expected 42 meta_id")
	}

}
