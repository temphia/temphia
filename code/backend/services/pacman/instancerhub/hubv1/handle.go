package hubv1

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

// shared handle

type Handle struct {
	instanced map[string]*xinstance.Response
	opts      repox.InstanceOptionsV1
	pacman    repox.BStore
}

func (h *Handle) GetFile(file string) ([]byte, error) {

	return h.pacman.GetBlob(h.opts.UserContext.TenantId, h.opts.BprintId, "", file)
}

func (h *Handle) LoadFile(file string, target any) error {
	return h.loadFile(h.opts.UserContext.TenantId, h.opts.BprintId, file, target)
}

func (h *Handle) GetPrevObject(name string) *xinstance.Response {
	return h.instanced[name]
}

func (h *Handle) loadFile(tenantId, bid string, file string, target any) error {
	out, err := h.pacman.GetBlob(tenantId, bid, "", file)
	if err != nil {
		return err
	}

	return json.Unmarshal(out, target)
}
