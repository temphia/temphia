package embed

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
)

type EmbedRepo struct {
	folder fs.FS
	name   string
}

func NewEmbed(name string, folder fs.FS) func(opts *xpacman.BuilderOptions) (xpacman.Repository, error) {
	return func(opts *xpacman.BuilderOptions) (xpacman.Repository, error) {
		return &EmbedRepo{
			folder: folder,
			name:   name,
		}, nil
	}
}

func (lr *EmbedRepo) Name() string {
	return lr.name
}

// "embed"

func (er *EmbedRepo) Query(tenantId string, opts *xpacman.RepoQuery) ([]xpacman.BPrint, error) {
	dirEntries, err := fs.ReadDir(er.folder, ".")
	if err != nil {
		return nil, err
	}

	items := make([]xpacman.BPrint, 0, len(dirEntries))
	for _, entry := range dirEntries {
		if entry.IsDir() {
			bp, err := er.readIndex(entry.Name())
			if err != nil {
				continue
			}

			items = append(items, *bp)
		}
	}

	return items, nil
}

func (er *EmbedRepo) Get(tenantid, slug string) (*xpacman.BPrint, error) {
	return er.readIndex(slug)
}

func (er *EmbedRepo) GetZip(tenantid, slug, version string) (io.ReadCloser, error) {

	tfile, err := xutils.ZipFsFolder(er.folder, slug)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(tfile)
	if err != nil {
		return nil, err
	}

	return xutils.NewTempFile(os.TempDir(), file), nil
}

func (er *EmbedRepo) readIndex(folder string) (*xpacman.BPrint, error) {
	bout, err := er.tryRead(folder, "bprint.json")
	if err != nil {
		return nil, err
	}

	bprint := &xpacman.BPrint{}
	err = json.Unmarshal(bout, &bprint)
	if err != nil {
		return nil, err
	}

	bprint.Versions = []string{"current"}

	return bprint, nil
}

func (er *EmbedRepo) tryRead(folder, file string) ([]byte, error) {
	f, err := er.folder.Open(path.Join(folder, file))
	if err != nil {
		return nil, err
	}

	return io.ReadAll(f)

}
