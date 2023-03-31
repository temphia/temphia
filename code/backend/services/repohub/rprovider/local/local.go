package local

import (
	"archive/zip"
	"encoding/json"
	"io"
	"os"
	"path"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/libx/xutils"
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

func (l *Local) Query(tenantId string, opts *repox.RepoQuery) ([]repox.BPrint, error) {

	dirs, err := os.ReadDir(l.rootFolder)
	if err != nil {
		return nil, err
	}

	resp := make([]repox.BPrint, 0, len(dirs))

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

		bprint := repox.BPrint{}
		err = json.Unmarshal(out, &bprint)
		if err != nil {
			pp.Println(err)
			continue
		}

		bprint.Versions = []string{"current"}

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

func (l *Local) Get(tenantid, slug string) (*repox.BPrint, error) {

	out, err := os.ReadFile(path.Join(l.rootFolder, slug, "index.json"))
	if err != nil {
		return nil, err
	}

	bprint := repox.BPrint{}
	err = json.Unmarshal(out, &bprint)
	if err != nil {
		return nil, err
	}

	bprint.Versions = []string{"current"}

	return &bprint, nil
}

func (l *Local) GetZip(tenantid, slug, version string) (io.ReadCloser, error) {

	out, err := os.ReadFile(path.Join(l.rootFolder, slug, "index.json"))
	if err != nil {
		return nil, err
	}

	bprint := entities.BPrint{}
	err = json.Unmarshal(out, &bprint)
	if err != nil {
		return nil, err
	}

	file, err := os.CreateTemp(os.TempDir(), "embed_*.zip")
	if err != nil {
		return nil, err
	}

	writer := zip.NewWriter(file)

	for _, file := range bprint.Files {

		fw, err := writer.Create(file)
		if err != nil {
			pp.Println(err)
			continue
		}

		out, err := os.ReadFile(path.Join(l.rootFolder, slug, file))
		if err != nil {
			pp.Println(err)
			continue
		}

		_, err = fw.Write(out)
		if err != nil {
			pp.Println(err)
			continue
		}
	}

	err = writer.Close()
	if err != nil {
		os.ReadFile(path.Join(os.TempDir(), file.Name()))
		return nil, err
	}

	return xutils.NewTempFile(os.TempDir(), file), nil
}
