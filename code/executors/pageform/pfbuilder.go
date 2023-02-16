package pageform

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
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
	return &Pageform{
		builder: pf,
		model: &FormModel{
			Name: "Test 1",
			Items: []FormItem{
				{
					Name: "name",
					Info: "Name of product",
					Type: "shorttext",
				},
				{
					Name: "Info",
					Info: "Product information",
					Type: "longtext",
				},
			},
			Data: make(map[string]any),
		},
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
	default:
		return []byte(``), nil
	}
}
