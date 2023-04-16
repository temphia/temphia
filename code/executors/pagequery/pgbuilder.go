package pagequery

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/elib/helper"
	"gopkg.in/yaml.v2"

	gojaExec "github.com/temphia/temphia/code/backend/engine/executors/javascript1/goja"
)

type PgBuilder struct {
	app    xtypes.App
	helper *helper.ExecutorHelper
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	h := helper.New("executor_pagequery", true)

	return &PgBuilder{
		app:    app.(xtypes.App),
		helper: h,
	}, nil
}

func (pg *PgBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	ffile := opts.EnvVars["model_file"]
	if ffile == "" {
		ffile = "model1.yaml"
	}

	out, _, err := opts.Binder.GetFileWithMeta(ffile)
	if err != nil {
		return nil, err
	}

	pp.Println("@model", string(out))

	model := &PgModel{}
	err = yaml.Unmarshal(out, model)
	if err != nil {
		return nil, err
	}

	rt := goja.New()

	if opts.File != "" {
		out, _, err := opts.Binder.GetFileWithMeta(opts.File)
		if err != nil {
			pp.Println(fmt.Sprintf("server.js could not load: %s", err.Error()))
			return nil, err
		}

		_, err = gojaExec.New(opts.Binder, rt)
		if err != nil {
			return nil, err
		}

		pp.Println(rt.RunString(string(out)))
	}

	return &PageQuery{
		builder:   pg,
		model:     model,
		jsruntime: rt,
		binder:    opts.Binder,
	}, nil
}

func (pf *PgBuilder) ExecFile(file string) ([]byte, error) {
	return pf.helper.Serve(file)
}
