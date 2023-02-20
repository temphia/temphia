package bdev

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func ZipIt(bprint *xbprint.LocalBprint, outFile string) error {

	// precheck if all files exists

	files := make([]string, 0)

	for fk, fpath := range bprint.Files {
		if !xutils.FileExists(fpath, "") {
			return easyerr.Error(fmt.Sprintf("[%s] file not found in path [%s]", fk, fpath))
		}

		files = append(files, fk)
	}

	archive, err := os.Create(outFile)
	if err != nil {
		return easyerr.Wrap("could not create zip file", err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	// write index.json

	newBprint := &entities.BPrint{
		Name:        bprint.Name,
		Slug:        bprint.Slug,
		Type:        bprint.Type,
		Icon:        bprint.Icon,
		Description: bprint.Description,
		Tags:        bprint.Tags,
		ExtraMeta:   bprint.ExtraMeta,
		Files:       files,
	}

	bout, err := json.Marshal(newBprint)
	if err != nil {
		return err
	}

	iwriter, err := zipWriter.Create("index.json")
	if err != nil {
		return err
	}

	_, err = iwriter.Write(bout)
	if err != nil {
		return err
	}

	// write all remaining files

	for fk, fpath := range bprint.Files {

		rfile, err := os.Open(fpath)
		if err != nil {
			return err
		}

		defer rfile.Close()

		wfile, err := zipWriter.Create(fk)
		if err != nil {
			return err
		}

		if _, err := io.Copy(wfile, rfile); err != nil {
			return err
		}
	}
	return zipWriter.Close()
}
