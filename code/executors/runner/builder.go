package runner

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/runner/rtypes"
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

	runner := New(&Options{
		BootstrapFunc: b.bootstrapFunc,
		Runcmd:        b.runcmd,
		EntryFile:     opts.File,
		GetFile: func(name string) ([]byte, error) {
			out, _, err := opts.Binder.GetFileWithMeta(name)
			return out, err
		},
	})

	return runner, nil
}

func (b *Builder) ExecFile(file string) ([]byte, error) {
	return nil, nil
}

func (b *Builder) IfaceFile() (*etypes.ExecutorIface, error) {
	return &etypes.ExecutorIface{}, nil
}
