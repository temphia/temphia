package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) AgentLinkNew(uclaim *claim.Session, data *entities.AgentLink) error {
	data.TenantId = uclaim.TenantId

	return c.coredb.AgentLinkNew(uclaim.TenantId, data)
}

func (c *Controller) AgentLinkUpdate(uclaim *claim.Session, pid, aid string, id int64, data map[string]any) error {
	return c.coredb.AgentLinkUpdate(uclaim.TenantId, pid, aid, id, data)
}

func (c *Controller) AgentLinkGet(uclaim *claim.Session, pid, aid string, id int64) (*entities.AgentLink, error) {
	return c.coredb.AgentLinkGet(uclaim.TenantId, pid, aid, id)
}

func (c *Controller) AgentLinkDel(uclaim *claim.Session, pid, aid string, id int64) error {
	return c.coredb.AgentLinkDel(uclaim.TenantId, pid, aid, id)
}

func (c *Controller) AgentLinkList(uclaim *claim.Session, pid, aid string) ([]*entities.AgentLink, error) {
	return c.coredb.AgentLinkList(uclaim.TenantId, pid, aid)
}

func (c *Controller) AgentExtensionNew(uclaim *claim.Session, data *entities.AgentExtension) error {
	data.TenantId = uclaim.TenantId

	return c.coredb.AgentExtensionNew(uclaim.TenantId, data)
}

func (c *Controller) AgentExtensionUpdate(uclaim *claim.Session, pid, aid string, id int64, data map[string]any) error {
	return c.coredb.AgentExtensionUpdate(uclaim.TenantId, pid, aid, id, data)
}

func (c *Controller) AgentExtensionGet(uclaim *claim.Session, pid, aid string, id int64) (*entities.AgentExtension, error) {
	return c.coredb.AgentExtensionGet(uclaim.TenantId, pid, aid, id)
}

func (c *Controller) AgentExtensionDel(uclaim *claim.Session, pid, aid string, id int64) error {
	return c.coredb.AgentExtensionDel(uclaim.TenantId, pid, aid, id)
}

func (c *Controller) AgentExtensionList(uclaim *claim.Session, pid, aid string) ([]*entities.AgentExtension, error) {
	return c.coredb.AgentExtensionList(uclaim.TenantId, pid, aid)
}

func (c *Controller) AgentResourceNew(uclaim *claim.Session, data *entities.AgentResource) error {
	data.TenantId = uclaim.TenantId

	return c.coredb.AgentResourceNew(uclaim.TenantId, data)
}

func (c *Controller) AgentResourceUpdate(uclaim *claim.Session, pid, aid, slug string, data map[string]any) error {
	return c.coredb.AgentResourceUpdate(uclaim.TenantId, pid, aid, slug, data)
}

func (c *Controller) AgentResourceGet(uclaim *claim.Session, pid, aid, slug string) (*entities.AgentResource, error) {
	return c.coredb.AgentResourceGet(uclaim.TenantId, pid, aid, slug)
}

func (c *Controller) AgentResourceDel(uclaim *claim.Session, pid, aid, slug string) error {
	return c.coredb.AgentResourceDel(uclaim.TenantId, pid, aid, slug)
}

func (c *Controller) AgentResourceList(uclaim *claim.Session, pid, aid string) ([]*entities.AgentResource, error) {
	return c.coredb.AgentResourceList(uclaim.TenantId, pid, aid)
}
