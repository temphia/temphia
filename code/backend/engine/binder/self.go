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

	activeModules map[string]etypes.Module
}

func NewSelfBindings(handle *handle.Handle, root *Binder) SelfBindings {

	return SelfBindings{
		root:          root,
		handle:        handle,
		pacman:        handle.Deps.Pacman,
		cabhub:        handle.Deps.CabinetHub,
		db:            handle.Deps.Corehub,
		runtime:       handle.Deps.Runtime,
		activeModules: make(map[string]etypes.Module),
	}
}

func (b *SelfBindings) SelfGetFile(file string) ([]byte, error) {
	return b.selfGetFile(file)
}

func (b *SelfBindings) SelfAddFile(file string, data []byte) error {
	return b.selfAddFile(file, data)
}

func (b *SelfBindings) SelfUpdateFile(file string, data []byte) error {
	return b.selfUpdateFile(file, data)
}

func (b *SelfBindings) SelfAddDataFile(file string, data []byte) error {
	return b.selfAddDataFile(file, data)
}

func (b *SelfBindings) SelfUpdateDataFile(file string, data []byte) error {
	return b.selfUpdateDataFile(file, data)
}

func (b *SelfBindings) SelfGetDataFile(file string) ([]byte, error) {
	return b.selfGetDataFile(file)
}

func (b *SelfBindings) SelfListDataFiles() (map[string]string, error) {
	return b.selfListDataFiles()
}

func (b *SelfBindings) SelfDeleteDataFile(file string) error {
	return b.selfDeleteDataFile(file)
}

func (b *SelfBindings) SelfModuleExec(name, method, path string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return b.selfModuleExec(name, method, path, data)
}

func (b *SelfBindings) SelfInLinks() ([]bindx.Link, error) {
	return b.selfInLinks()
}

func (b *SelfBindings) SelfOutLinks() ([]bindx.Link, error) {
	return b.selfOutLinks()
}

func (b *SelfBindings) SelfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error) {
	return b.selfLinkExec(name, method, data, async, detached)
}

func (b *SelfBindings) SelfForkExec(method string, data []byte) error {
	return b.selfForkExec(method, data)
}
