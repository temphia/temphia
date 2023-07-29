package dev

import (
	"io"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/thoas/go-funk"
)

type Controller struct {
	pacman  repox.Pacman
	corehub store.CoreHub
}

func New(pacman repox.Pacman, corehub store.CoreHub) *Controller {
	return &Controller{
		pacman:  pacman,
		corehub: corehub,
	}
}

func (c *Controller) DevBprintFileList(tkt *claim.PlugDevTkt) (map[string]string, error) {
	return c.pacman.BprintListBlobs(tkt.TenantId, tkt.BprintId)
}

func (c *Controller) DevBprintFileDel(tkt *claim.PlugDevTkt, file string) error {
	return c.pacman.BprintDeleteBlob(tkt.TenantId, tkt.BprintId, file)
}

func (c *Controller) DevBprintFileGet(tkt *claim.PlugDevTkt, file string) ([]byte, error) {
	return c.pacman.BprintGetBlob(tkt.TenantId, tkt.BprintId, file)
}

func (c *Controller) DevPushFiles(tkt *claim.PlugDevTkt, files map[string]io.Reader) error {

	bprint, err := c.pacman.BprintGet(tkt.TenantId, tkt.BprintId)
	if err != nil {
		return err
	}

	if bprint.Files == nil {
		bprint.Files = make(entities.JsonArray, 0)

	}

	for filekey, filerc := range files {
		out, err := io.ReadAll(filerc)
		if err != nil {
			return err
		}

		if !funk.ContainsString(bprint.Files, filekey) {
			err = c.pacman.BprintNewBlob(tkt.TenantId, tkt.BprintId, filekey, out, true)
			if err != nil {
				return err
			}
			bprint.Files = append(bprint.Files, filekey)
		} else {
			err = c.pacman.BprintUpdateBlob(tkt.TenantId, tkt.BprintId, filekey, out)
			if err != nil {
				return err
			}
		}

	}

	return c.pacman.BprintUpdateFilesList(tkt.TenantId, tkt.BprintId, bprint.Files...)
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
