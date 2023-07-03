package re

// re stands for Remote Executor.

import (
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/re/rtypes"
)

type Options struct {
	BootstrapFunc func(ctx rtypes.BootstrapContext) error
	Runcmd        string
	EntryFile     string
	GetFile       func(name string) ([]byte, error)
}

type Runner struct {
	opts  *Options
	token string

	blines     map[string]*bindingsLine
	blinesLock sync.Mutex

	controlLine *controlLine
	clineLock   sync.Mutex
}

func New(opts *Options) *Runner {
	r := &Runner{
		opts:        opts,
		token:       "",
		controlLine: nil,
		clineLock:   sync.Mutex{},
		blines:      make(map[string]*bindingsLine),
		blinesLock:  sync.Mutex{},
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
		File:     r.opts.EntryFile,
		GetFile:  r.opts.GetFile,
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
