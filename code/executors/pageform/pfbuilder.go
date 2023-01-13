package pageform

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type PfBuilder struct {
	app xtypes.App
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	return &PfBuilder{
		app: app.(xtypes.App),
	}, nil
}

func (pf *PfBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return &Pageform{
		builder: pf,
	}, nil
}

func (pf *PfBuilder) ExecFile(file string) ([]byte, error) {
	pp.Println("@file", file)

	return nil, nil
}
