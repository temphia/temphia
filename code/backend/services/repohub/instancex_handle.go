package repohub

import (
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

type Handle struct {
	instanced map[string]*xinstance.Response
	opts      repox.InstanceOptions
	pacman    *PacMan
}

func (h *Handle) GetFile(file string) ([]byte, error) {
	return h.pacman.BprintGetBlob(h.opts.UserSession.TenentId, h.opts.BprintId, file)
}

func (h *Handle) LoadFile(file string, target any) error { // loads json/yaml
	return h.pacman.loadFile(h.opts.UserSession.TenentId, h.opts.BprintId, file, target)
}

func (h *Handle) GetPrevObject(name string) *xinstance.Response {
	return h.instanced[name]
}
