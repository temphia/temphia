package re

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/re/rtypes"
)

type Builder struct {
	name          string
	runcmd        string
	bootstrapFunc rtypes.BootstrapFunc
}

func NewBuilder(name, runcmd string, bootstrapFunc rtypes.BootstrapFunc) *Builder {
	return &Builder{
		name:          name,
		runcmd:        runcmd,
		bootstrapFunc: bootstrapFunc,
	}

}

func (b *Builder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	runner := New(Options{
		BootstrapFunc: b.bootstrapFunc,
		Runcmd:        b.runcmd,
	}, opts)

	return runner, nil
}

func (b *Builder) ExecFile(file string) ([]byte, error) {
	return []byte(""), nil
}

func (b *Builder) IfaceFile() (*etypes.ExecutorIface, error) {
	return &etypes.ExecutorIface{}, nil
}
