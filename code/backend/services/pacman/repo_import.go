package pacman

import (
	"archive/zip"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
)

func (p *PacMan) RepoImport(tenantid string, opts *xpacman.RepoImportOpts) (string, error) {

	pp.Println("@repo_import")

	reader, err := p.RepoGetZip(tenantid, opts.Source, opts.Slug, opts.Version)
	if err != nil {
		pp.Println("@could_not_get_zip", err.Error())
		return "", err
	}
	pp.Println("@after_get_zip")

	return p.BprintCreateFromZip(tenantid, reader)
}

func (p *PacMan) RepoList(tenantid, group string, source int64, tags ...string) ([]xpacman.BPrint, error) {
	repo := p.getRepoSource(tenantid, source)

	if repo == nil {
		return nil, easyerr.NotFound("repo")
	}

	return repo.Query(tenantid, &xpacman.RepoQuery{
		Group: group,
		Tags:  tags,
		Page:  0,
	})

}

func (p *PacMan) RepoGet(tenantid, slug string, source int64) (*xpacman.BPrint, error) {
	repo := p.getRepoSource(tenantid, source)
	return repo.Get(tenantid, slug)
}

func (p *PacMan) RepoGetZip(tenantid string, source int64, slug, version string) (io.ReadCloser, error) {
	repo := p.getRepoSource(tenantid, source)
	return repo.GetZip(tenantid, slug, version)
}

func (p *PacMan) BprintCreateFromZip(tenantId string, rawreader io.ReadCloser) (string, error) {

	defer rawreader.Close()

	pp.Println("@bprint_from_zip")

	file, err := os.CreateTemp(os.TempDir(), "import_bprint*.zip")
	if err != nil {
		return "", err
	}
	defer func() {
		name := file.Name()
		file.Close()
		os.Remove(path.Join(os.TempDir(), name))
	}()

	bn, err := io.Copy(file, rawreader)
	if err != nil {
		pp.Println("@copy_err", err.Error())
		return "", err
	}

	pp.Println("@copy_bytes", bn)

	err = file.Sync()
	if err != nil {
		return "", err
	}

	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	reader, err := zip.NewReader(file, info.Size())
	if err != nil {
		pp.Println("@couldnot open zip", err.Error())
		pp.Println("@reader", file.Name())

		return "", err
	}

	ifile, err := reader.Open("bprint.json")
	if err != nil {
		return "", err
	}

	bprint := &entities.BPrint{}
	err = json.NewDecoder(ifile).Decode(bprint)
	if err != nil {
		return "", err
	}

	bprint.TenantID = tenantId
	bid, err := p.BprintCreate(tenantId, bprint)
	if err != nil {
		return "", err
	}

	cleanUpFunc := func() {
		// try to clean up
		p.corehub.BprintDel(tenantId, bid)
		p.bstore.DeleteRoot(tenantId, bid)
	}

	err = p.bstore.NewRoot(tenantId, bid)
	if err != nil {
		cleanUpFunc()
		return "", err
	}

	for _, file := range reader.File {
		if !file.FileInfo().IsDir() {
			continue
		}

		err = p.bstore.NewFolder(tenantId, bid, file.Name)
		if err != nil {
			cleanUpFunc()
			return "", err
		}
	}

	for _, file := range reader.File {
		if file.Name == "bprint.json" || file.FileInfo().IsDir() {
			continue
		}

		name := file.Name
		folder := ""
		frags := strings.Split(file.Name, "/")
		if len(frags) > 1 {
			name = frags[len(frags)-1]
			folder = strings.TrimRight(file.Name, name)
		}

		rfile, err := file.Open()
		if err != nil {
			cleanUpFunc()
			return "", err
		}

		out, err := ioutil.ReadAll(rfile)
		if err != nil {
			cleanUpFunc()
			return "", err
		}

		err = p.bstore.NewBlob(tenantId, bid, folder, name, out)
		if err != nil {
			cleanUpFunc()
			rfile.Close()
			return "", err
		}
		rfile.Close()
	}

	return bid, err

}
