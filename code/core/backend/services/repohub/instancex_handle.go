package repohub

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox/xinstance"
)

type Handle struct {
	instanced map[string]*xinstance.Response
	opts      repox.InstanceOptions
	pacman    *PacMan
}

func (h *Handle) GetFile(file string) ([]byte, error) {
	return h.pacman.RepoSourceGetBlob(h.opts.UserSession.TenentId, "", h.opts.BprintId, h.opts.RepoId, file)
}

func (h *Handle) LoadFile(file string, target any) error { // loads json/yaml
	return h.pacman.loadFile(h.opts.UserSession.TenentId, h.opts.BprintId, file, h.opts.RepoId, target)
}

func (h *Handle) GetPrevObject(name string) *xinstance.Response {
	return h.instanced[name]
}
