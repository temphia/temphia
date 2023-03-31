package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

// syncme => repobuild code https://github.com/temphia/repo

type DB struct {
	GroupIndex map[string][]string     `json:"group_index" yaml:"group_index"`
	TagIndex   map[string][]string     `json:"tag_index" yaml:"tag_index"`
	Items      map[string]repox.BPrint `json:"items" yaml:"items"`
}

/*


	structure

	data
		index.json
		example1/
			index.json
			version.zip


*/

type Github struct {
	user   string
	repo   string
	branch string

	cache map[string]repox.BPrint
	cLock sync.Mutex
}

func (g *Github) Name() string { return "github" }

func (g *Github) Query(tenantId string, opts *repox.RepoQuery) ([]repox.BPrint, error) {
	if g.cache == nil {
		err := g.fillCache()
		if err != nil {
			return nil, err
		}
	}

	vals := make([]repox.BPrint, 0, len(g.cache))
	for _, v := range g.cache {
		vals = append(vals, v)
	}

	return vals, nil

}

func (g *Github) Get(tenantid, slug string) (*repox.BPrint, error) {
	if g.cache == nil {
		err := g.fillCache()
		if err != nil {
			return nil, err
		}
	}

	bp, ok := g.cache[slug]
	if ok {
		return nil, easyerr.NotFound()
	}

	return &bp, nil
}

func (g *Github) GetZip(tenantid, slug, version string) (io.ReadCloser, error) {

	resp, err := http.Get(g.fileURL("data/"+slug, version+".zip"))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

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

func (g *Github) fileURL(folder, file string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s/%s", g.user, g.repo, g.branch, folder, file)
}

func (g *Github) mainDB() string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/data/db.json", g.user, g.repo, g.branch)
}
