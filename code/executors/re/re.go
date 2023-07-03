package re

// re stands for Remote Executor.

import (
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/re/rtypes"
)

type Options struct {
	BootstrapFunc func(ctx rtypes.BootstrapContext) error
	Runcmd        string
}

type Runner struct {
	opts      Options
	token     string
	entryFile string

	rootBinding bindx.Core

	controlLine *controlLine
	clineLock   sync.Mutex
}

func New(opts Options, eopts etypes.ExecutorOption) *Runner {

	tok, err := xutils.GenerateRandomString(10)
	if err != nil {
		panic(err)
	}

	r := &Runner{
		opts:        opts,
		token:       tok,
		controlLine: nil,
		clineLock:   sync.Mutex{},
		rootBinding: eopts.Binder.Clone(),
		entryFile:   eopts.File,
	}

	return r
}

func (r *Runner) Init() error {

	// ExRunner

	tdir, err := os.MkdirTemp("", "temphia_runner_*")
	if err != nil {
		return err
	}

	r.opts.BootstrapFunc(rtypes.BootstrapContext{
		Folder:   tdir,
		TenantId: "",
		PlugId:   "",
		AgentId:  "",
		File:     r.entryFile,
		GetFile: func(name string) ([]byte, error) {
			out, _, err := r.rootBinding.GetFileWithMeta(name)
			return out, err
		},
	})

	actualcmd := strings.Split(r.opts.Runcmd, " ")

	runcmd := exec.Command(actualcmd[0], actualcmd[1:]...)

	err = runcmd.Run()
	if err != nil {
		return err
	}

	return r.startServer()
}

func (r *Runner) Process(*event.Request) (*event.Response, error) {

	return nil, nil
}
