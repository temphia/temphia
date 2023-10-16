package bunjs

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Builder struct {
	chub  store.CabinetHub
	waits map[string]chan *etypes.RemoteOptions
	wlock sync.Mutex
}

type Config struct {
	SubFolder string   `json:"sub_folder,omitempty"`
	Entry     string   `json:"entry,omitempty"`
	Files     []string `json:"files,omitempty"`
}

func (b *Builder) New(opts etypes.ExecutorOption) (etypes.Executor, error) {

	err := b.initData(opts)
	if err != nil {
		return nil, err
	}

	wchan := make(chan *etypes.RemoteOptions)
	key := opts.PlugId + opts.AgentId

	b.wlock.Lock()
	b.waits[key] = wchan
	b.wlock.Unlock()

	switch opts.DefaultRunner {
	case "":
		cmd := exec.Command("bun", "run", "main.js")
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
		proxy:     httputil.NewSingleHostReverseProxy(u),
		rPXPrefix: ropts.RPXPrefix,
		addr:      ropts.Addr,
	}, nil
}

func (b *Builder) initData(opts etypes.ExecutorOption) error {

	if doesFolderExist(opts.RunFolder) {
		return nil
	}

	err := os.MkdirAll(opts.RunFolder, os.ModePerm)
	if err != nil {
		return err
	}

	conf := Config{}
	fout, _, err := opts.Binder.GetFileWithMeta(opts.File)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fout, &conf)
	if err != nil {
		return err
	}

	zfile, err := b.chub.GetFolderAsZip(context.TODO(), opts.TenantId, path.Join(xtypes.BprintBlobFolder, opts.BprintId))
	if err != nil {
		return err
	}

	defer func() {
		os.Remove(zfile)
	}()

	return xutils.ExtractZip(zfile, opts.RunFolder)
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
