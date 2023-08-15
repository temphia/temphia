package dev

import (
	"io"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/thoas/go-funk"
)

type Controller struct {
	pacman  repox.Pacman
	bstore  repox.BStore
	corehub store.CoreHub
}

func New(pacman repox.Pacman, corehub store.CoreHub) *Controller {
	return &Controller{
		pacman:  pacman,
		corehub: corehub,
		bstore:  pacman.GetBprintFileStore(),
	}
}

func (c *Controller) DevBprintFileList(tkt *claim.PlugDevTkt) (map[string]string, error) {
	return c.pacman.BprintListBlobs(tkt.TenantId, tkt.BprintId)
}

func (c *Controller) DevBprintFileDel(tkt *claim.PlugDevTkt, file string) error {

	return c.bstore.DeleteBlob(tkt.TenantId, tkt.BprintId, "", file)
}

func (c *Controller) DevBprintFileGet(tkt *claim.PlugDevTkt, file string) ([]byte, error) {
	return c.bstore.GetBlob(tkt.TenantId, tkt.BprintId, "", file)
}

func (c *Controller) DevPushFiles(tkt *claim.PlugDevTkt, files map[string]io.Reader) error {

	// fixme => folder support

	ffs, err := c.bstore.ListBlob(tkt.TenantId, tkt.BprintId, "")
	if err != nil {
		return err
	}

	bffs := make(map[string]*store.BlobInfo)

	for _, file := range ffs {
		bffs[file.Name] = file
	}

	for filekey, filerc := range files {
		out, err := io.ReadAll(filerc)
		if err != nil {
			return err
		}

		if _, ok := bffs[filekey]; ok {
			err = c.bstore.NewBlob(tkt.TenantId, tkt.BprintId, "", filekey, out)
			if err != nil {
				return err
			}

		} else {
			err = c.bstore.UpdateBlob(tkt.TenantId, tkt.BprintId, "", filekey, out)
			if err != nil {
				return err
			}
		}

	}

	return nil

}

func (c *Controller) DevModifyPlug(tkt *claim.PlugDevTkt, pid string, data map[string]any) error {
	if !only(data, "name", "executor", "live", "dev", "handlers", "extra_meta") {
		return easyerr.Error("Not allowed field")
	}

	return c.corehub.PlugUpdate(tkt.TenantId, pid, data)
}

func (c *Controller) DevModifyAgent(tkt *claim.PlugDevTkt, pid string, aid string, data map[string]any) error {
	if !only(data, "name", "type", "executor", "iface_file", "web_files", "web_entry", "web_script", "web_style", "web_loader", "extra_meta", "env_vars") {
		return easyerr.Error("Not allowed field")
	}

	return c.corehub.AgentUpdate(tkt.TenantId, pid, aid, data)
}

func only(data map[string]any, keys ...string) bool {
	if len(data) > len(keys) {
		return false
	}

	for hkey := range data {
		if !funk.ContainsString(keys, hkey) {
			return false
		}
	}
	return true
}
