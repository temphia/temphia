package cabinet

import (
	"context"
	"path"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/fdatautil"
)

type Controller struct {
	hub    store.CabinetHub
	signer service.Signer
}

func New(cabinet store.CabinetHub, signer service.Signer) *Controller {
	return &Controller{
		hub:    cabinet,
		signer: signer,
	}
}

func (c *Controller) ListRoot(uclaim *claim.Session, source string) ([]string, error) {

	resp, err := c.hub.ListFolder(context.Background(), uclaim.TenantId, "")
	if err != nil {
		return nil, err
	}

	rstrs := make([]string, 0, len(resp))

	for _, bi := range resp {
		rstrs = append(rstrs, bi.Name)
	}

	return rstrs, nil
}

func (c *Controller) AddFolder(uclaim *claim.Session, source, folder string) error {
	dir := ""
	name := strings.TrimPrefix(folder, "/")
	if strings.Contains(name, "/") {
		dir, name = path.Split(folder)
	}

	return c.hub.NewFolder(context.Background(), uclaim.TenantId, dir, name)

}

func (c *Controller) AddBlob(uclaim *claim.Session, source, folder, file string, contents []byte) error {
	return c.hub.NewFile(context.Background(), uclaim.TenantId, folder, file, fdatautil.NewFromBytes(contents))
}

func (c *Controller) ListFolder(uclaim *claim.Session, source, folder string) ([]*store.BlobInfo, error) {
	return c.hub.ListFolder(context.Background(), uclaim.TenantId, folder)
}

func (c *Controller) GetBlob(uclaim *claim.Session, source, fpath string) ([]byte, error) {

	data, err := c.hub.GetFile(context.Background(), uclaim.TenantId, fpath)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	out, err := data.AsBytes()
	return out, err
}

func (c *Controller) DeleteBlob(uclaim *claim.Session, source, folder, file string) error {
	return c.hub.DeleteFile(context.Background(), uclaim.TenantId, folder, file)
}
