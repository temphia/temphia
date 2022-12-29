package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
)

// syncme => repobuild code https://github.com/temphia/repo

type DB struct {
	GroupIndex map[string][]string         `json:"group_index" yaml:"group_index"`
	TagIndex   map[string][]string         `json:"tag_index" yaml:"tag_index"`
	Items      map[string]*entities.BPrint `json:"items" yaml:"items"`
}

/*

// fixme =>
	- auth for private repos ?
	- shard by group and cache by group
	- expire cache and fill on expire
	- if cache fill expire then take time before retrying
*/

type Github struct {
	user   string
	repo   string
	branch string

	cache map[string]*entities.BPrint
	cLock sync.Mutex
}

func (g *Github) Name() string { return "github" }

func (g *Github) Query(tenantId string, opts *repox.RepoQuery) ([]entities.BPrint, error) {
	if g.cache == nil {
		err := g.fillCache()
		if err != nil {
			return nil, err
		}
	}

	vals := make([]entities.BPrint, 0, len(g.cache))
	for _, v := range g.cache {
		vals = append(vals, *v)
	}

	return vals, nil
}

func (g *Github) GetItem(tenantid, group, slug string) (*entities.BPrint, error) {

	data := &entities.BPrint{}

	out, err := g.GetFile(tenantid, group, slug, "index.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(out, data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (g *Github) GetFile(tenantid, group, slug, file string) ([]byte, error) {
	resp, err := http.Get(g.fileURL(("data/" + slug), file))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (g *Github) GetFileURL(tenantid, group, slug, file string) (string, error) {
	return g.fileURL(("data/" + slug), file), nil
}

// private

func (g *Github) fillCache() error {

	resp, err := http.Get(g.mainDB())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	db := DB{}

	dc := json.NewDecoder(resp.Body)
	err = dc.Decode(&db)
	if err != nil {
		return err
	}

	if g.cache == nil {
		g.cLock.Lock()
		g.cache = db.Items
		g.cLock.Unlock()
	}

	return nil
}

func (g *Github) mainDB() string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/data/db.json", g.user, g.repo, g.branch)
}

func (g *Github) fileURL(folder, file string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s/%s", g.user, g.repo, g.branch, folder, file)
}
