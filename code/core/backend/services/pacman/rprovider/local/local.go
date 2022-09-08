package local

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/core/backend/app/registry"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
)

func init() {
	registry.SetRepoBuilder("local", New)
}

type Local struct {
	rootFolder string
}

func New(opts *repox.BuilderOptions) (repox.Repository, error) {

	return &Local{
		rootFolder: opts.BaseURL,
	}, nil
}

func (l *Local) Name() string { return "local" }

func (l *Local) Query(tenantId string, opts *repox.RepoQuery) ([]entities.BPrint, error) {

	dirs, err := ioutil.ReadDir(l.rootFolder)
	if err != nil {
		return nil, err
	}

	resp := make([]entities.BPrint, 0, len(dirs))

	for _, dir := range dirs {
		if !dir.IsDir() {
			pp.Println("Skipping Dir", dir.Name())
			continue
		}

		out, err := ioutil.ReadFile(path.Join(l.rootFolder, dir.Name(), "index.json"))
		if err != nil {
			pp.Println(err)
			continue
		}

		bprint := entities.BPrint{}
		err = json.Unmarshal(out, &bprint)
		if err != nil {
			pp.Println(err)
			continue
		}

		if opts.Group != "" {
			if opts.Group != bprint.Type {
				pp.Println("Skipping Dir", dir.Name())
				continue
			}
		}

		resp = append(resp, bprint)
	}

	return resp, nil
}

func (l *Local) GetItem(tenantid, group, slug string) (*entities.BPrint, error) {
	out, err := ioutil.ReadFile(path.Join(l.rootFolder, slug, "index.json"))
	if err != nil {
		return nil, err
	}

	bprint := &entities.BPrint{}
	err = json.Unmarshal(out, bprint)
	if err != nil {
		return nil, err
	}

	return bprint, nil
}

func (l *Local) GetFile(tenantid, group, slug, file string) ([]byte, error) {
	return ioutil.ReadFile(path.Join(l.rootFolder, slug, file))
}

func (l *Local) GetFileURL(tenantid, group, slug, file string) (string, error) {
	return "", nil
}
