package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

// func (g *gitlab) getIndex() map[string]map[string]*entities.BPrint {
// 	g.mu.Lock()
// 	defer g.mu.Unlock()

// 	if g.cacheIndex != nil {
// 		return g.cacheIndex
// 	}
// 	bps, err := fetchBlueprints(g.base + glIndexFile)
// 	if err != nil {
// 		return nil
// 	}

// 	ci := make(map[string]map[string]*entities.BPrint)

// 	for _, bprint := range bps {
// 		gci, ok := ci[bprint.Type]
// 		if !ok {
// 			gci = make(map[string]*entities.BPrint)
// 			ci[bprint.Type] = gci
// 		}
// 		gci[bprint.Slug] = &bprint
// 	}
// 	g.cacheIndex = ci
// 	return g.cacheIndex
// }

func (g *gitlab) list(tenantid, group string, tags ...string) ([]entities.BPrint, error) {
	return fetchBlueprints(g.base + "/" + glIndexFile)
}

func (g *gitlab) get(tenantid, group, slug string) (*entities.BPrint, error) {
	url := fmt.Sprintf("%s/data/%s/%s/index.json", g.base, group, slug)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	newbprnt := &entities.BPrint{}

	err = json.Unmarshal(bytes, newbprnt)
	if err != nil {
		return nil, err
	}

	return newbprnt, nil
}

func (g *gitlab) getBlob(tenantid, group, slug, file string) ([]byte, error) {

	return fetchBytes(g.fileUrl(group, slug, file))
}

func (g *gitlab) GetBlobURL(tenantid, group, slug, file string) (string, error) {

	return g.fileUrl(group, slug, file), nil
}

func fetchBlueprints(url string) ([]entities.BPrint, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var final []entities.BPrint

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(out, &final)
	return final, err
}

func fetchBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("count not fetch file http bad status")
	}

	return ioutil.ReadAll(resp.Body)
}

func (g *gitlab) fileUrl(group, slug, file string) string {
	return fmt.Sprintf("%s/data/%s/%s/%s", g.base, group, slug, file)
}
