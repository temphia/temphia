package local

import (
	"encoding/json"
	"os"
	"path"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/app/registry"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

func init() {
	registry.SetRepoBuilder("local", New)
	registry.SetRepoBuilder("dev", DevNew)
}

type Local struct {
	rootFolder string
	name       string
}

func DevNew(opts *repox.BuilderOptions) (repox.Repository, error) {

	return &Local{
		rootFolder: "./cmd/dev/repo",
		name:       "dev",
	}, nil
}

func New(opts *repox.BuilderOptions) (repox.Repository, error) {

	return &Local{
		rootFolder: opts.BasePath,
		name:       "local",
	}, nil
}

func (l *Local) Name() string { return l.name }

func (l *Local) Query(tenantId string, opts *repox.RepoQuery) ([]entities.BPrint, error) {

	dirs, err := os.ReadDir(l.rootFolder)
	if err != nil {
		return nil, err
	}

	resp := make([]entities.BPrint, 0, len(dirs))

	for _, dir := range dirs {
		if !dir.IsDir() {
			pp.Println("Skipping Dir", dir.Name())
			continue
		}

		out, err := os.ReadFile(path.Join(l.rootFolder, dir.Name(), "index.json"))
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
	out, err := os.ReadFile(path.Join(l.rootFolder, slug, "index.json"))
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
	return os.ReadFile(path.Join(l.rootFolder, slug, file))
}

func (l *Local) GetFileURL(tenantid, group, slug, file string) (string, error) {
	return "", nil
}
