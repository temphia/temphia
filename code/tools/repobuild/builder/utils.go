package builder

import (
	"archive/zip"
	"crypto/sha1"
	"encoding/base64"
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

func (rb *RepoBuilder) hashedBuidlPath(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return ppath.Join(rb.config.BuildFolder, sha)
}

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

	defer zipWriter.Close()

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

		// its a file
		if !strings.HasSuffix(fk, "/") {
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

			zipEntry, err := zipWriter.Create(ppath.Join(fk, d.Name()))
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
	pp.Println(outFile)

	return nil
}
