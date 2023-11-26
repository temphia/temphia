package builder

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	ppath "path"
	"path/filepath"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

type Zipper struct {
	bprint *xpackage.Manifest
	file   *os.File
	zfile  *zip.Writer
}

func NewZipper(bprint *xpackage.Manifest, outFile string) (*Zipper, error) {

	for fk, fpath := range bprint.Files {
		if !xutils.FileExists(fpath, "") {
			return nil, easyerr.Error(fmt.Sprintf("[%s] file not found in path [%s]", fk, fpath))
		}
	}

	archive, err := os.Create(outFile)
	if err != nil {
		return nil, easyerr.Wrap("could not create zip file", err)
	}

	zipWriter := zip.NewWriter(archive)

	return &Zipper{
		bprint: bprint,
		file:   archive,
		zfile:  zipWriter,
	}, nil
}

func (z *Zipper) Build() error {

	// write index.json

	newBprint := &entities.BPrint{
		Name:        z.bprint.Name,
		Slug:        z.bprint.Slug,
		Type:        z.bprint.Type,
		Icon:        z.bprint.Icon,
		Description: z.bprint.Description,
		Tags:        z.bprint.Tags,
		ExtraMeta:   z.bprint.ExtraMeta,
	}

	bout, err := json.Marshal(newBprint)
	if err != nil {
		return err
	}

	err = z.writeFile("bprint.json", bout)
	if err != nil {
		return err
	}

	appj := &xpackage.AppSchema{
		Name:    z.bprint.Name,
		Objects: z.bprint.Objects,
		Steps:   z.bprint.Steps,
	}

	bout, err = json.Marshal(appj)
	if err != nil {
		return err
	}

	err = z.writeFile("app.json", bout)
	if err != nil {
		return err
	}

	log.Println("creating zip")

	// write all remaining files

	for fk, fpath := range z.bprint.Files {
		log.Println("addng file: ", fk, fpath)

		// its a file
		if !strings.HasSuffix(fk, "/") {
			rfile, err := os.Open(fpath)
			if err != nil {
				return err
			}

			defer rfile.Close()

			wfile, err := z.zfile.Create(fk)
			if err != nil {
				return err
			}

			if _, err := io.Copy(wfile, rfile); err != nil {
				return err
			}

			continue
		}

		// its a folder, now zip all folders contents
		err = filepath.WalkDir(fk, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			pp.Println("@zipping_file", path)

			zipEntry, err := z.zfile.Create(ppath.Join(fk, d.Name()))
			if err != nil {
				return err
			}

			file, err := os.Open(fpath)
			if err != nil {
				return err
			}

			defer file.Close()

			_, err = io.Copy(zipEntry, file)
			return err
		})

		if err != nil {
			return err
		}

	}

	log.Println("create zip ok")

	return nil
}

func (z *Zipper) writeFile(file string, data []byte) error {

	iwriter, err := z.zfile.Create(file)
	if err != nil {
		return err
	}

	_, err = iwriter.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (z *Zipper) Close() {
	z.zfile.Close()
	z.file.Close()
}
