package self

import (
	"github.com/temphia/temphia/code/core/backend/engine/binders/standard/handle"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Binding struct {
	handle  *handle.Handle
	pacman  service.Pacman
	cabhub  store.CabinetHub
	db      store.CoreHub
	runtime etypes.Runtime
}

func New(handle *handle.Handle) Binding {

	return Binding{
		handle:  handle,
		pacman:  handle.Deps.Pacman,
		cabhub:  handle.Deps.CabinetHub,
		db:      handle.Deps.Corehub,
		runtime: handle.Deps.Runtime,
	}
}

func (b *Binding) SelfGetFile(file string) ([]byte, error) {
	return b.selfGetFile(file)
}

func (b *Binding) SelfAddFile(file string, data []byte) error {
	return b.selfAddFile(file, data)
}

func (b *Binding) SelfUpdateFile(file string, data []byte) error {
	return b.selfUpdateFile(file, data)
}

func (b *Binding) SelfAddDataFile(file string, data []byte) error {
	return b.selfAddDataFile(file, data)
}

func (b *Binding) SelfUpdateDataFile(file string, data []byte) error {
	return b.selfUpdateDataFile(file, data)
}

func (b *Binding) SelfGetDataFile(file string) ([]byte, error) {
	return b.selfGetDataFile(file)
}

func (b *Binding) SelfListDataFiles() (map[string]string, error) {
	return b.selfListDataFiles()
}

func (b *Binding) SelfDeleteDataFile(file string) error {
	return b.selfDeleteDataFile(file)
}

func (b *Binding) SelfModuleExec(name, method, path string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return b.selfModuleExec(name, method, path, data)
}

func (b *Binding) SelfInLinks() ([]bindx.Link, error) {
	return b.selfInLinks()
}

func (b *Binding) SelfOutLinks() ([]bindx.Link, error) {
	return b.selfOutLinks()
}

func (b *Binding) SelfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error) {
	return b.selfLinkExec(name, method, data, async, detached)
}

func (b *Binding) SelfForkExec(method string, data []byte) error {
	return b.selfForkExec(method, data)
}
