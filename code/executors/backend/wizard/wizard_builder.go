package wizard

import (
	"io/ioutil"
	"os"

	"path"
	"strings"

	"github.com/dop251/goja"
	"github.com/goccy/go-yaml"

	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/libx/xutils"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"
	"github.com/ztrue/tracerr"
)

func init() {
	registry.SetExecutor("simple.wizard", func(app interface{}) (etypes.ExecutorBuilder, error) {
		return &SWBuilder{}, nil
	})
}

type SWBuilder struct{}

var DevMode = true

func New(opts etypes.ExecutorOption) (etypes.Executor, error) {

	if DevMode {
		return newDev(opts)
	}

	out, _, err := opts.Binder.GetFileWithMeta("wizard.yaml")
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	model := wmodels.Wizard{}

	err = yaml.Unmarshal(out, &model)
	if err != nil {
		return nil, err
	}

	return &SimpleWizard{
		model:         model,
		binding:       opts.Binder,
		jsRuntime:     nil,
		nativeScripts: nil,
	}, nil
}

const fpath = "../executors/backend/wizard/sample/"

func newDev(opts etypes.ExecutorOption) (etypes.Executor, error) {

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	out, err := os.ReadFile(path.Join(pwd, fpath, "_two.yaml"))
	if err != nil {
		return nil, err
	}

	hookjs, err := os.ReadFile(path.Join(pwd, fpath, "_one.js"))
	if err != nil {
		hookjs = []byte(``)
	}

	model := wmodels.Wizard{}

	err = yaml.Unmarshal(out, &model)
	if err != nil {
		return nil, err
	}

	rt := goja.New()

	_, err = rt.RunString(string(hookjs))
	if err != nil {
		return nil, err
	}

	for skey, s := range model.Stages {
		s.Name = skey
	}

	for k, s := range model.Sources {
		s.Name = k
	}

	return &SimpleWizard{
		model:         model,
		binding:       opts.Binder,
		jsRuntime:     rt,
		nativeScripts: nil,
	}, nil

}

func (sd *SWBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return New(opts)
}

func (sd *SWBuilder) ExecFile(file string) ([]byte, error) {

	if strings.HasSuffix(file, ".css") {
		if xutils.FileExists(DevPath, "wizard.css") {
			return sd.serveFile(DevPath, "wizard.css")
		}
		return loaderCSS, nil
	}

	if strings.HasSuffix(file, ".js") {
		if xutils.FileExists(DevPath, "wizard.js") {
			return sd.serveFile(DevPath, "wizard.js")
		}
		return loaderJS, nil
	}

	if strings.HasSuffix(file, ".js.map") {
		if xutils.FileExists(DevPath, "wizard.js.map") {
			return sd.serveFile(DevPath, "wizard.js.map")
		}
	}

	return nil, easyerr.NotFound()
}

func (sd *SWBuilder) Init(app interface{}) error {
	return nil
}

func (sd *SWBuilder) serveFile(dpath, file string) ([]byte, error) {
	return ioutil.ReadFile(path.Join("executors/frontend/public/build/", file))
}
