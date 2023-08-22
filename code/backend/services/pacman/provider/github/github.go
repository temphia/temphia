package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/tools/repobuild/index"
)

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

	cache map[string]xpacman.BPrint
	// cacheFillWip bool
	cLock sync.Mutex
}

func (g *Github) Name() string { return "github" }

func (g *Github) Query(tenantId string, opts *xpacman.RepoQuery) ([]xpacman.BPrint, error) {

	err := g.fillCache()
	if err != nil {
		return nil, err
	}

	vals := make([]xpacman.BPrint, 0, len(g.cache))
	for _, v := range g.cache {
		vals = append(vals, v)
	}

	return vals, nil

}

func (g *Github) Get(tenantid, slug string) (*xpacman.BPrint, error) {

	err := g.fillCache()
	if err != nil {
		return nil, err
	}

	bp, ok := g.cache[slug]
	if !ok {
		return nil, easyerr.NotFound("cache bprint")
	}

	return &bp, nil
}

func (g *Github) GetZip(tenantid, slug, version string) (io.ReadCloser, error) {

	zipurl := g.fileURL("data/"+slug, version+".zip")

	pp.Println(zipurl)

	resp, err := http.Get(zipurl)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (g *Github) fillCache() error {

	if g.cache != nil {
		// fixme => do lastcheckTime here and when it timeouts refill cache
		return nil
	}

	resp, err := http.Get(g.mainDB())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	db := index.DB{}

	dc := json.NewDecoder(resp.Body)
	err = dc.Decode(&db)
	if err != nil {
		return err
	}

	g.cLock.Lock()
	if g.cache == nil {
		g.cache = db.Items
	}
	g.cLock.Unlock()

	return nil
}

func (g *Github) fileURL(folder, file string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s/%s", g.user, g.repo, g.branch, folder, file)
}

func (g *Github) mainDB() string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/data/db.json", g.user, g.repo, g.branch)
}
