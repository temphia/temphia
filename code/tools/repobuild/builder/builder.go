package builder

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/tools/repobuild/index"
	"github.com/temphia/temphia/code/tools/repobuild/models"
)

type RepoBuilder struct {
	config *models.BuildConfig

	// build stage states
	erroredItems map[string]error
	outputs      map[string]string

	// index stage states
	indexer *index.Indexer

	repoCache      map[string]*git.Repository
	bprintFileCace map[string]*xbprint.LocalBprint
}

func New(conf *models.BuildConfig) *RepoBuilder {

	return &RepoBuilder{
		config:         conf,
		indexer:        index.New(path.Join(conf.OutputFolder, "db.json")),
		erroredItems:   make(map[string]error),
		outputs:        make(map[string]string),
		repoCache:      make(map[string]*git.Repository),
		bprintFileCace: make(map[string]*xbprint.LocalBprint),
	}
}

func (rb *RepoBuilder) Build() error {

	os.RemoveAll(rb.config.BuildFolder)

	for k := range rb.config.Items {

		ofolder, err := rb.buildItem(k)
		if err != nil {
			rb.erroredItems[k] = err
			continue
		}
		rb.outputs[k] = ofolder
	}

	return rb.indexer.Save()
}

func (rb *RepoBuilder) PrintResult() {
	for k, err := range rb.erroredItems {
		pp.Println("@err", k, err)
	}

	for k, v := range rb.outputs {
		pp.Println("@build_ok", k, v)
	}

}
