package pageform

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"gopkg.in/yaml.v2"
)

type PfBuilder struct {
	app xtypes.App
	dev bool
}

var loaderJs []byte
var loaderCss []byte

func init() {
	loaderJs, _ = os.ReadFile("code/frontend/public/build/executor_pageform.js")
	loaderCss, _ = os.ReadFile("code/frontend/public/build/executor_pageform.css")
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	return &PfBuilder{
		app: app.(xtypes.App),
		dev: true,
	}, nil
}

func (pf *PfBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	out, _, err := opts.Binder.GetFileWithMeta("form1.yaml") // fixme => get_this from somewhere
	if err != nil {
		return nil, err
	}

	form := &FormModel{}
	err = yaml.Unmarshal(out, form)
	if err != nil {
		return nil, err
	}

	return &Pageform{
		builder: pf,
		model:   form,
	}, nil
}

func (pf *PfBuilder) ExecFile(file string) ([]byte, error) {
	pp.Println("@file", file)

	switch file {
	case "loader.css":
		if pf.dev {
			return os.ReadFile("code/frontend/public/build/executor_pageform.css")
		}

		return loaderCss, nil
	case "loader.js":
		if pf.dev {
			return os.ReadFile("code/frontend/public/build/executor_pageform.js")
		}

		return loaderJs, nil

	case "executor_pageform.js.map":
		return os.ReadFile("code/frontend/public/build/executor_pageform.js.map")

	default:
		return []byte(``), nil
	}
}
