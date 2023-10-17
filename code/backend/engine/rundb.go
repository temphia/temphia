package engine

import (
	"encoding/json"
	"io"
	"os"

	"github.com/temphia/temphia/code/backend/libx/xutils"
)

type entry struct {
	BprintId string `json:"bprint_id,omitempty"`
	RunId    string `json:"run_id,omitempty"`
	PlugId   string `json:"plug_id,omitempty"`
	AgentId  string `json:"agent_id,omitempty"`
}

type runDB struct {
	file        *os.File
	entrieCache map[string]*entry // <plug_id + agent_id, entry>
}

func newRunDB(f string) runDB {

	dbexist := xutils.FileExists("", f)

	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	entrieCache := make(map[string]*entry)

	if dbexist {
		out, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(out, &entrieCache)
		if err != nil {
			panic(err)
		}
	} else {
		_, err = file.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
	}

	return runDB{
		file:        file,
		entrieCache: entrieCache,
	}
}

func (r *runDB) set(e *entry) {
	r.entrieCache[e.PlugId+e.AgentId] = e
}

func (r *runDB) get(plugId, agentId string) *entry {
	return r.entrieCache[plugId+agentId]
}

func (r *runDB) flush() error {

	out, err := json.Marshal(r.entrieCache)
	if err != nil {
		return err
	}

	_, err = r.file.Write(out)

	return err
}
