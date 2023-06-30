package runner

import (
	"testing"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/executors/runner/python"
)

func TestRunner(t *testing.T) {

	runner := New(&Options{
		BootstrapFile:   string(python.Bootstrap),
		ExecutorLibData: string(python.Lib),
		ExecutorLibName: "temphia_python_executor",
		RootFilesFunc: func() ([]byte, error) {
			return nil, nil
		},
	})

	pp.Println(runner.init())

}
