package bdev

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

func ZipIt(bprint *xpackage.Manifest, outFile string) error {

	for fk, fpath := range bprint.Files {
		if !xutils.FileExists(fpath, "") {
			return easyerr.Error(fmt.Sprintf("[%s] file not found in path [%s]", fk, fpath))
		}

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
	}

	bout, err := json.Marshal(newBprint)
	if err != nil {
		return err
	}

	iwriter, err := zipWriter.Create("bprint.json")
	if err != nil {
		return err
	}

	_, err = iwriter.Write(bout)
	if err != nil {
		return err
	}

	log.Println("creating zip")

	// write all remaining files

	for fk, fpath := range bprint.Files {
		log.Println("addng file: ", fk, fpath)

		rfile, err := os.Open(fpath)
		if err != nil {
			return err
		}

		defer rfile.Close()

		wfile, err := zipWriter.Create(fk)
		if err != nil {
			return err
		}

		// fixme => add subfolder

		if _, err := io.Copy(wfile, rfile); err != nil {
			return err
		}
	}

	log.Println("create zip ok")
	pp.Println(outFile)

	return zipWriter.Close()
}
