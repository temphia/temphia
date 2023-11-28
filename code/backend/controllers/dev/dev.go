package dev

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Controller struct {
	pacman  xpacman.Pacman
	bstore  xpacman.BStore
	corehub store.CoreHub
	ecache  etypes.Ecache
}

func New(pacman xpacman.Pacman, corehub store.CoreHub) *Controller {
	return &Controller{
		pacman:  pacman,
		corehub: corehub,
		bstore:  pacman.GetBprintFileStore(),
	}
}

func (c *Controller) ArtifactList(dclaim *claim.PlugDevTkt) ([]*store.BlobInfo, error) {
	plug := c.ecache.GetPlug(dclaim.TenantId, dclaim.PlugId)
	if plug == nil {
		return nil, easyerr.NotFound("plug")
	}

	// plug.BprintId

	return c.bstore.ListBlob(dclaim.TenantId, plug.BprintId, "")
}

func (c *Controller) ArtifactPush(dclaim *claim.PlugDevTkt) {

}

func (c *Controller) ArtifactGet(dclaim *claim.PlugDevTkt) {

}

func (c *Controller) ArtifactDelete(dclaim *claim.PlugDevTkt) {

}

func (c *Controller) ArtifactBulkPush(dclaim *claim.PlugDevTkt) {

}
