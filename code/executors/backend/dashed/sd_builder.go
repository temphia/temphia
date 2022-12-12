package dashed

import (
	"io/ioutil"
	"strings"

	"github.com/goccy/go-yaml"

	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/backend/dashed/dashmodels"

	"github.com/ztrue/tracerr"
)

func init() {
	registry.SetExecutor("simple.dash", func(app interface{}) (etypes.ExecutorBuilder, error) {
		return &SDBuilder{}, nil
	})
}

type SDBuilder struct{}

func (sd *SDBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return New(opts)
}

func (sd *SDBuilder) ExecFile(file string) ([]byte, error) {
	if strings.HasSuffix(file, ".js") {
		return ioutil.ReadFile("frontend/public/build/dashed.js")
	}

	if strings.HasSuffix(file, ".css") {
		return ioutil.ReadFile("frontend/public/build/dashed.css")
	}

	if strings.HasSuffix(file, ".js.map") {
		return ioutil.ReadFile("frontend/public/build/dashed.js.map")
	}

	return nil, easyerr.NotFound()
}

func (sd *SDBuilder) Init(app interface{}) error {
	return nil
}

func New(opts etypes.ExecutorOption) (*SimpleDash, error) {

	out, _, err := opts.Binder.GetFileWithMeta("dash.yaml")
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	model := dashmodels.Dashboard{}

	err = yaml.Unmarshal(out, &model)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &SimpleDash{
		model:    model,
		bindings: opts.Binder,
	}, nil
}
