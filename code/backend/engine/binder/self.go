package binder

import (
	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type SelfBindings struct {
	handle  *handle.Handle
	pacman  repox.Hub
	cabhub  store.CabinetHub
	db      store.CoreHub
	runtime etypes.Runtime
	root    *Binder

	activeModules    map[int32]etypes.Module
	activeModCounter int32
}

func NewSelfBindings(handle *handle.Handle, root *Binder) SelfBindings {

	return SelfBindings{
		root:          root,
		handle:        handle,
		pacman:        handle.Deps.Pacman,
		cabhub:        handle.Deps.CabinetHub,
		db:            handle.Deps.Corehub,
		runtime:       handle.Deps.Runtime,
		activeModules: make(map[int32]etypes.Module),
	}
}

func (b *SelfBindings) NewModule(name string, data xtypes.LazyData) (int32, error) {
	return b.selfNewModule(name, data)
}

func (b *SelfBindings) ModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return b.selfModuleExec(mid, method, data)
}

func (b *SelfBindings) InLinks() ([]bindx.Link, error) {
	return b.selfInLinks()
}

func (b *SelfBindings) OutLinks() ([]bindx.Link, error) {
	return b.selfOutLinks()
}

func (b *SelfBindings) LinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error) {
	return b.selfLinkExec(name, method, data, async, detached)
}

func (b *SelfBindings) ForkExec(method string, data []byte) error {
	return b.selfForkExec(method, data)
}
