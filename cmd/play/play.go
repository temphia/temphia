package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {

	err := ZipIt("test1.zip", map[string]string{
		"testdata/pp.txt": "pp.txt",
		"testdata/mnop/":  "mnop/",
	})

	pp.Println("@err", err.Error())

}

func ZipIt(outFile string, files map[string]string) error {

	archive, err := os.Create(outFile)
	if err != nil {
		return err
	}

	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	defer zipWriter.Close()

	log.Println("creating zip")

	for kfile, vfile := range files {

		pp.Println("@K/V", kfile, vfile)

		if !strings.HasSuffix(kfile, "/") {
			rfile, err := os.Open(kfile)
			if err != nil {
				return err
			}

			defer rfile.Close()

			wfile, err := zipWriter.Create(vfile)
			if err != nil {
				return err
			}

			if _, err := io.Copy(wfile, rfile); err != nil {
				return err
			}

			continue
		}

		err = filepath.WalkDir(kfile, func(fpath string, d fs.DirEntry, err error) error {
			pp.Println("@fpath", fpath)

			if err != nil {
				pp.Println("@err", err.Error())
				return err
			}

			if d.IsDir() {
				pp.Println("@skipping dir")
				return nil
			}

			zipEntry, err := zipWriter.Create(path.Join(vfile, d.Name()))
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

	log.Println("end zip")

	return nil

}
