package engine

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type getFileReq struct {
	File string `json:"file,omitempty"`
}

func (e *Engine) remotePerform(opts etypes.Remote) ([]byte, error) {
	b := e.getBinding(opts.TenantId, opts.PlugId, opts.AgentId)

	var err error
	var resp any

	switch opts.Action {
	case "get_self_file":
		req := &getFileReq{}
		err = json.Unmarshal(opts.Data, req)
		if err != nil {
			break
		}
		resp, _, err = b.GetFileWithMeta(req.File)
	}

	if err != nil {
		return nil, err
	}

	return json.Marshal(resp)

}

/*

	Log(eid, msg string)
	LazyLog(eid string, msgs []string)

	GetSelfFile(file string) (data []byte, err error)
	ListResources() ([]*Resource, error)
	GetResource(name string) (*Resource, error)
	InLinks() ([]Link, error)
	OutLinks() ([]Link, error)

	LinkExec(name, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	LinkExecEmit(name, method string, data xtypes.LazyData) error

	NewModule(name string, data xtypes.LazyData) (int32, error)
	ModuleTicket(name string, opts xtypes.LazyData) (string, error)
	ModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error)

	UserContext(eid string) *claim.UserContext

*/
