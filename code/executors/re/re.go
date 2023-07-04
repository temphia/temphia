package re

// re stands for Remote Executor.

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

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

	tenantId string
	plugId   string
	agentId  string

	rootBinding bindx.Core

	controlLine *controlLine
	clineLock   sync.Mutex

	listener net.Listener
}

func New(opts Options, eopts etypes.ExecutorOption) *Runner {

	tok, err := xutils.GenerateRandomString(10)
	if err != nil {
		panic(err)
	}

	r := &Runner{
		opts:        opts,
		token:       tok,
		tenantId:    eopts.TenantId,
		plugId:      eopts.PlugId,
		agentId:     eopts.AgentId,
		controlLine: nil,
		clineLock:   sync.Mutex{},
		rootBinding: eopts.Binder.Clone(),
		entryFile:   eopts.File,
	}

	return r
}

func (r *Runner) Init() error {
	tdir, err := os.MkdirTemp("", "temphia_runner_*")
	if err != nil {
		return err
	}

	err = r.opts.BootstrapFunc(rtypes.BootstrapContext{
		Folder:   tdir,
		TenantId: r.tenantId,
		PlugId:   r.plugId,
		AgentId:  r.agentId,
		File:     r.entryFile,
		GetFile: func(name string) ([]byte, error) {
			out, _, err := r.rootBinding.GetFileWithMeta(name)
			return out, err
		},
	})
	if err != nil {
		return err
	}

	err = r.startServer()
	if err != nil {
		return err
	}

	actualcmd := strings.Split(r.opts.Runcmd, " ")

	runcmd := exec.Command(actualcmd[0], actualcmd[1:]...)

	runcmd.Env = append(runcmd.Env,
		fmt.Sprintf("TEMPHIA_REMOTE_PORT=%d", 1234),
		fmt.Sprintf("TEMPHIA_TOKEN=%s", r.token),
		fmt.Sprintf("TEMPHIA_TENANT_ID=%s", r.tenantId),
		fmt.Sprintf("TEMPHIA_PLUG_ID=%s", r.plugId),
		fmt.Sprintf("TEMPHIA_AGENT_ID=%s", r.agentId),
	)

	err = runcmd.Run()
	if err != nil {
		return err
	}

	go r.acceptLoop()

	return nil
}

func (r *Runner) Close() error {
	return r.listener.Close()
}

func (r *Runner) Process(ev *event.Request) (*event.Response, error) {

	if r.controlLine == nil {
		fmt.Print("SLEEPING #")
		for {
			fmt.Print("#")
			time.Sleep(time.Second * 2)
			if r.controlLine != nil {
				break
			}
		}

	}

	return r.controlLine.process(ev)
}
