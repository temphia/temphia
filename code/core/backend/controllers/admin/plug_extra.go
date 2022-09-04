package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) AgentLinkNew(uclaim *claim.Session, data *entities.AgentLink) error {
	return c.coredb.AgentLinkNew(uclaim.TenentId, data)
}

func (c *Controller) AgentLinkUpdate(uclaim *claim.Session, pid, aid string, id int64, data map[string]any) error {
	return c.coredb.AgentLinkUpdate(uclaim.TenentId, pid, aid, id, data)
}

func (c *Controller) AgentLinkGet(uclaim *claim.Session, pid, aid string, id int64) (*entities.AgentLink, error) {
	return c.coredb.AgentLinkGet(uclaim.TenentId, pid, aid, id)
}

func (c *Controller) AgentLinkDel(uclaim *claim.Session, pid, aid string, id int64) error {
	return c.coredb.AgentLinkDel(uclaim.TenentId, pid, aid, id)
}

func (c *Controller) AgentLinkList(uclaim *claim.Session, pid, aid string) ([]*entities.AgentLink, error) {
	return c.coredb.AgentLinkList(uclaim.TenentId, pid, aid)
}

func (c *Controller) AgentExtensionNew(uclaim *claim.Session, data *entities.AgentExtension) error {
	return c.coredb.AgentExtensionNew(uclaim.TenentId, data)
}

func (c *Controller) AgentExtensionUpdate(uclaim *claim.Session, pid, aid string, id int64, data map[string]any) error {
	return c.coredb.AgentExtensionUpdate(uclaim.TenentId, pid, aid, id, data)
}

func (c *Controller) AgentExtensionGet(uclaim *claim.Session, pid, aid string, id int64) (*entities.AgentExtension, error) {
	return c.coredb.AgentExtensionGet(uclaim.TenentId, pid, aid, id)
}

func (c *Controller) AgentExtensionDel(uclaim *claim.Session, pid, aid string, id int64) error {
	return c.coredb.AgentExtensionDel(uclaim.TenentId, pid, aid, id)
}

func (c *Controller) AgentExtensionList(uclaim *claim.Session, pid, aid string) ([]*entities.AgentExtension, error) {
	return c.coredb.AgentExtensionList(uclaim.TenentId, pid, aid)
}

func (c *Controller) AgentResourceNew(uclaim *claim.Session, data *entities.AgentResource) error {
	return c.coredb.AgentResourceNew(uclaim.TenentId, data)
}

func (c *Controller) AgentResourceUpdate(uclaim *claim.Session, pid, aid, slug string, data map[string]any) error {
	return c.coredb.AgentResourceUpdate(uclaim.TenentId, pid, aid, slug, data)
}

func (c *Controller) AgentResourceGet(uclaim *claim.Session, pid, aid, slug string) (*entities.AgentResource, error) {
	return c.coredb.AgentResourceGet(uclaim.TenentId, pid, aid, slug)
}

func (c *Controller) AgentResourceDel(uclaim *claim.Session, pid, aid, slug string) error {
	return c.coredb.AgentResourceDel(uclaim.TenentId, pid, aid, slug)
}

func (c *Controller) AgentResourceList(uclaim *claim.Session, pid, aid string) ([]*entities.AgentResource, error) {
	return c.coredb.AgentResourceList(uclaim.TenentId, pid, aid)
}
