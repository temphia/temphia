package pageform

import (
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/helper"
	"gopkg.in/yaml.v2"

	gojaExec "github.com/temphia/temphia/code/backend/engine/executors/javascript1/goja"
)

type PfBuilder struct {
	app    xtypes.App
	helper *helper.ExecutorHelper
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	h := helper.New("executor_pageform", true)

	return &PfBuilder{
		app:    app.(xtypes.App),
		helper: h,
	}, nil
}

func (pf *PfBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	ffile := opts.EnvVars["form_file"]
	if ffile == "" {
		ffile = "form1.yaml"
	}

	out, _, err := opts.Binder.GetFileWithMeta(ffile)
	if err != nil {
		return nil, err
	}

	pp.Println("@form1", string(out))

	form := &FormModel{}
	err = yaml.Unmarshal(out, form)
	if err != nil {
		return nil, err
	}

	rt := goja.New()

	if opts.File != "" {
		out, _, err := opts.Binder.GetFileWithMeta(opts.File)
		if err != nil {
			pp.Println("server.js not found")
			return nil, err
		}

		_, err = gojaExec.New(opts.Binder, rt)
		if err != nil {
			return nil, err
		}

		pp.Println(rt.RunString(string(out)))
	}

	return &Pageform{
		builder: pf,
		model:   form,
		runtime: rt,
		binder:  opts.Binder,
	}, nil
}

func (pf *PfBuilder) ExecFile(file string) ([]byte, error) {
	return pf.helper.Serve(file)
}
