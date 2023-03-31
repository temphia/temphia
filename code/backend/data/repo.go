package data

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

type EmbedRepo struct {
	assetStore *DataBox
}

func NewEmbed(opts *repox.BuilderOptions) (repox.Repository, error) {
	return &EmbedRepo{
		assetStore: defaultDataBox,
	}, nil
}

func (lr *EmbedRepo) Name() string {
	return "embed"
}

func (er *EmbedRepo) Query(tenantId string, opts *repox.RepoQuery) ([]repox.BPrint, error) {
	out, err := er.assetStore.tryRead("", "repo", "index.json")
	if err != nil {
		return nil, err
	}

	index := map[string]string{}
	err = json.Unmarshal(out, &index)
	if err != nil {
		return nil, err
	}

	items := make([]repox.BPrint, 0, len(index))
	for _, idxfile := range index {

		bp, err := er.readIndex(idxfile)
		if err != nil {
			continue
		}

		items = append(items, *bp)
	}

	return items, nil
}

func (er *EmbedRepo) Get(tenantid, slug string) (*repox.BPrint, error) {
	return er.readIndex(fmt.Sprintf("%s_index.json", slug))
}

func (er *EmbedRepo) GetZip(tenantid, slug, version string) (io.ReadCloser, error) {

	bout, err := er.assetStore.tryRead("", "repo", fmt.Sprintf("%s_index.json", slug))
	if err != nil {
		return nil, err
	}

	bprint := entities.BPrint{}
	err = json.Unmarshal(bout, &bprint)
	if err != nil {
		return nil, err
	}

	file, err := os.CreateTemp(os.TempDir(), "embed_*.zip")
	if err != nil {
		return nil, err
	}

	writer := zip.NewWriter(file)

	for _, file := range bprint.Files {
		bout, err := er.assetStore.tryRead("", "repo", fmt.Sprintf("%s_%s", slug, file))
		if err != nil {
			return nil, err
		}

		fw, err := writer.Create(file)
		if err != nil {
			pp.Println(err)
			continue
		}

		_, err = fw.Write(bout)
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

func (er *EmbedRepo) readIndex(file string) (*repox.BPrint, error) {
	bout, err := er.assetStore.tryRead("", "repo", file)
	if err != nil {
		return nil, err
	}

	bprint := repox.BPrint{}
	err = json.Unmarshal(bout, &bprint)
	if err != nil {
		return nil, err
	}

	bprint.Versions = []string{"current"}

	return &bprint, nil
}
