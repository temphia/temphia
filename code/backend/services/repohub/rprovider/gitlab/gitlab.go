package gitlab

import (
	"net/url"
	"path"
	"sync"

	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

const (
	glIndexFile = "index.json"
)

func init() {
	registry.SetRepoBuilder("gitlab", NewProvider)
}

func NewProvider(opts *repox.BuilderOptions) (repox.Repository, error) {
	// https://gitlab.com/temphia/blueprint_store/-/raw/master/

	url, err := url.Parse(opts.BasePath)
	if err != nil {
		return nil, err
	}
	url.Path = path.Join(url.Path, "-/raw/master/")
	provider := New(url.String())
	return provider, nil
}

type gitlab struct {
	mu         sync.Mutex
	cacheIndex map[string]map[string]*entities.BPrint
	base       string
}

func New(url string) *gitlab {

	return &gitlab{
		cacheIndex: nil,
		base:       url,
		mu:         sync.Mutex{},
	}
}

func (g *gitlab) Name() string {
	return "gitlab"
}

func (g *gitlab) Query(tenantId string, opts *repox.RepoQuery) ([]entities.BPrint, error) {
	return g.list(tenantId, opts.Group, opts.Tags...)
}

func (g *gitlab) GetItem(tenantid, group, slug string) (*entities.BPrint, error) {
	return g.get(tenantid, group, slug)
}

func (g *gitlab) GetFile(tenantid, group, slug, file string) ([]byte, error) {
	return g.getBlob(tenantid, group, slug, file)
}

func (g *gitlab) GetFileURL(tenantid, group, slug, file string) (string, error) {
	return "", nil
}
