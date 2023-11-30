package bunjs

import (
	"context"
	"errors"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/fdatautil"
)

type Builder struct {
	chub   store.CabinetHub
	signer service.Signer
	waits  map[string]chan *etypes.RemoteOptions
	wlock  sync.Mutex
	confd  config.Confd
}

type Config struct {
	SubFolder  string `json:"sub_folder,omitempty"`
	RunCommand string `json:"run_command,omitempty"`
}

func (b *Builder) New(opts etypes.ExecutorOption) (etypes.Executor, error) {
	conf := Config{}
	err := b.initData(opts, &conf)
	if err != nil {
		return nil, err
	}

	wchan := make(chan *etypes.RemoteOptions)
	key := opts.PlugId + opts.AgentId

	b.wlock.Lock()
	b.waits[key] = wchan
	b.wlock.Unlock()

	var cmd *exec.Cmd

	switch opts.DefaultRunner {
	case "":

		token, err := b.signer.SignRemoteExec(opts.TenantId, &claim.RemoteExec{
			TenantId: opts.TenantId,
			Plug:     opts.PlugId,
			Agent:    opts.AgentId,
		})
		if err != nil {
			return nil, err
		}

		cmd = exec.Command("bun", "run", conf.RunCommand)
		cmd.Dir = opts.RunFolder
		cmd.Env = b.confd.GetRemoteExecEnvs(opts.PlugId, opts.AgentId, opts.BprintId, token)

		err = cmd.Start()
		if err != nil {
			return nil, err
		}

	default:
		panic("not implemented runner")
	}

	ropts := <-wchan

	b.wlock.Lock()
	delete(b.waits, key)
	b.wlock.Unlock()

	u, err := url.Parse(ropts.Addr)
	if err != nil {
		return nil, err
	}

	return &BunJS{
		tenantId:  opts.TenantId,
		plugId:    opts.PlugId,
		agentId:   opts.AgentId,
		cmd:       cmd,
		proxy:     httputil.NewSingleHostReverseProxy(u),
		rPXPrefix: ropts.RPXPrefix,
		addr:      ropts.Addr,
	}, nil
}

func (b *Builder) initData(opts etypes.ExecutorOption, conf *Config) error {

	if doesFolderExist(opts.RunFolder) {
		return nil
	}

	err := os.MkdirAll(opts.RunFolder, os.ModePerm)
	if err != nil {
		return err
	}

	conf.RunCommand = opts.WebOptions["run_command"]
	conf.SubFolder = opts.WebOptions["run_sub_folder"]

	data, err := b.chub.CompressFolder(context.TODO(), opts.TenantId, path.Join(xtypes.BprintBlobFolder, opts.BprintId, conf.SubFolder))
	if err != nil {
		return err
	}

	return fdatautil.ExtractZipAndClose(data, opts.RunFolder)
}

func (b *Builder) ServeFile(file string) (xtypes.BeBytes, error) {

	pp.Println("@serving file", file)

	return nil, nil
}

func (b *Builder) SetRemoteOptions(opts etypes.RemoteOptions) error {

	key := opts.PlugId + opts.AgentId

	b.wlock.Lock()
	wchan := b.waits[key]
	b.wlock.Unlock()

	if wchan == nil {
		return easyerr.NotFound("wait chan")
	}

	wchan <- &opts

	return nil
}

func doesFolderExist(folder string) bool {
	// check folder contents ??
	_, err := os.Stat(folder)
	if err == nil {
		return true
	}

	return !errors.Is(err, os.ErrNotExist)
}
