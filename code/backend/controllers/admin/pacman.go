package admin

import (
	"archive/zip"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (c *Controller) BprintList(uclaim *claim.Session, group string) ([]*entities.BPrint, error) {

	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintList(uclaim.TenantId, group)
}

func (c *Controller) BprintCreate(uclaim *claim.Session, bp *entities.BPrint) (string, error) {
	if !uclaim.IsSuperAdmin() {
		return "", easyerr.NotImpl()
	}

	bp.TenantID = uclaim.TenantId
	return c.pacman.BprintCreate(uclaim.TenantId, bp)
}

func (c *Controller) BprintUpdate(uclaim *claim.Session, bp *entities.BPrint) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintUpdate(uclaim.TenantId, bp)
}

func (c *Controller) BprintGet(uclaim *claim.Session, bid string) (*entities.BPrint, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintGet(uclaim.TenantId, bid)
}

func (c *Controller) BprintRemove(uclaim *claim.Session, bid string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintRemove(uclaim.TenantId, bid)
}

func (c *Controller) BprintListBlobs(uclaim *claim.Session, bid string) (any, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintListBlobs(uclaim.TenantId, bid)
}

func (c *Controller) BprintNewBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintNewBlob(uclaim.TenantId, bid, file, payload, true)
}

func (c *Controller) BprintUpdateBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}
	return c.pacman.BprintUpdateBlob(uclaim.TenantId, bid, file, payload)
}

func (c *Controller) BprintGetBlob(uclaim *claim.Session, bid, file string) ([]byte, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintGetBlob(uclaim.TenantId, bid, file)
}

func (c *Controller) BprintDeleteBlob(uclaim *claim.Session, bid, file string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintDeleteBlob(uclaim.TenantId, bid, file)
}

func (c *Controller) BprintCreateFromZip(uclaim *claim.Session, form *multipart.Form) error {
	zfile := form.File["file"][0]
	zipfile, err := zfile.Open()
	if err != nil {
		return err
	}

	defer zipfile.Close()

	reader, err := zip.NewReader(zipfile, zfile.Size)
	if err != nil {
		return err
	}

	ifile, err := reader.Open("index.json")
	if err != nil {
		return err
	}

	bprint := &entities.BPrint{}
	err = json.NewDecoder(ifile).Decode(bprint)
	if err != nil {
		return err
	}

	files := make([]string, 0)
	bprint.Files = entities.JsonArray{}

	bprint.TenantID = uclaim.TenantId
	bid, err := c.pacman.BprintCreate(uclaim.TenantId, bprint)
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		if file.Name == "index.json" {
			continue
		}

		rfile, err := file.Open()
		if err != nil {
			return err
		}

		out, err := ioutil.ReadAll(rfile)
		if err != nil {
			return err
		}

		err = c.pacman.BprintNewBlob(uclaim.TenantId, bid, file.Name, out, false)
		if err != nil {
			rfile.Close()
			return err
		}

		files = append(files, file.Name)

		rfile.Close()
	}

	pp.Println("@files", files)

	return c.pacman.BprintUpdateFilesList(bprint.TenantID, bid, files...)
}

// repo

func (c *Controller) BprintImport(uclaim *claim.Session, opts *repox.RepoImportOpts) (string, error) {
	return c.pacman.RepoSourceImport(uclaim.TenantId, opts)
}

type InstanceOptions struct {
	InstancerType  string          `json:"instancer_type,omitempty"`
	File           string          `json:"file,omitempty"`
	UserConfigData json.RawMessage `json:"data,omitempty"`
	Auto           bool            `json:"auto,omitempty"`
}

func (c *Controller) BprintInstance(uclaim *claim.Session, bid string, opts *InstanceOptions) (any, error) {

	pp.Println(" ||>>", opts)
	pp.Println(" ||>>", string(opts.UserConfigData))

	instancer := c.pacman.GetInstanceHub()

	fopt := repox.InstanceOptions{
		BprintId:       bid,
		InstancerType:  opts.InstancerType,
		File:           opts.File,
		UserConfigData: opts.UserConfigData,
		Auto:           opts.Auto,
		UserSession:    uclaim,
	}

	switch opts.InstancerType {
	case xbprint.TypeBundle:
		if opts.Auto {
			return instancer.AutomaticBundle(fopt)
		}
		return instancer.ManualBundleItem(fopt)

	default:
		if opts.Auto {
			return instancer.AutomaticSingle(fopt)
		}
		return instancer.ManualSingle(fopt)
	}
}
