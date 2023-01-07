package corehub

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

// repo

func (c *CoreHub) RepoNew(tenantId string, data *entities.Repo) error {
	return c.coredb.RepoNew(tenantId, data)
}

func (c *CoreHub) RepoUpdate(tenantId string, id int64, data map[string]any) error {
	return c.coredb.RepoUpdate(tenantId, id, data)
}

func (c *CoreHub) RepoGet(tenantId string, id int64) (*entities.Repo, error) {
	return c.coredb.RepoGet(tenantId, id)
}

func (c *CoreHub) RepoDel(tenantId string, id int64) error {
	return c.coredb.RepoDel(tenantId, id)
}

func (c *CoreHub) RepoList(tenantId string) ([]*entities.Repo, error) {
	return c.coredb.RepoList(tenantId)
}

// bprint

func (c *CoreHub) BprintNew(tenantId string, et *entities.BPrint) error {
	return c.coredb.BprintNew(tenantId, et)
}

func (c *CoreHub) BprintUpdate(tenantId, id string, data map[string]any) error {
	return c.coredb.BprintUpdate(tenantId, id, data)
}

func (c *CoreHub) BprintGet(tenantId, id string) (*entities.BPrint, error) {
	return c.coredb.BprintGet(tenantId, id)
}

func (c *CoreHub) BprintDel(tenantId, id string) error {
	return c.coredb.BprintDel(tenantId, id)
}

func (c *CoreHub) BprintList(tenantId, group string) ([]*entities.BPrint, error) {
	return c.coredb.BprintList(tenantId, group)
}

// plug

func (c *CoreHub) PlugNew(tenantId string, pg *entities.Plug) error {
	return c.coredb.PlugNew(tenantId, pg)
}

func (c *CoreHub) PlugUpdate(tenantId string, id string, data map[string]any) error {
	return c.coredb.PlugUpdate(tenantId, id, data)
}

func (c *CoreHub) PlugGet(tenantId, pid string) (*entities.Plug, error) {
	return c.coredb.PlugGet(tenantId, pid)
}

func (c *CoreHub) PlugDel(tenantId, pid string) error {
	return c.coredb.PlugDel(tenantId, pid)
}

func (c *CoreHub) PlugList(tenantId string) ([]*entities.Plug, error) {
	return c.coredb.PlugList(tenantId)
}

func (c *CoreHub) PlugListByBprint(tenantId, bid string) ([]*entities.Plug, error) {
	return c.coredb.PlugListByBprint(tenantId, bid)
}

func (c *CoreHub) AgentNew(tenantId string, data *entities.Agent) error {
	return c.coredb.AgentNew(tenantId, data)
}

func (c *CoreHub) AgentUpdate(tenantId, pid, id string, data map[string]any) error {
	return c.coredb.AgentUpdate(tenantId, pid, id, data)
}

func (c *CoreHub) AgentGet(tenantId, pid, id string) (*entities.Agent, error) {
	return c.coredb.AgentGet(tenantId, pid, id)
}

func (c *CoreHub) AgentDel(tenantId, pid, agentId string) error {
	return c.coredb.AgentDel(tenantId, pid, agentId)
}

func (c *CoreHub) AgentList(tenantId, pid string) ([]*entities.Agent, error) {
	return c.coredb.AgentList(tenantId, pid)
}

// resource
func (c *CoreHub) ResourceNew(tenantId string, obj *entities.Resource) error {
	return c.coredb.ResourceNew(tenantId, obj)
}

func (c *CoreHub) ResourceUpdate(tenantId string, id string, data map[string]any) error {
	return c.coredb.ResourceUpdate(tenantId, id, data)
}

func (c *CoreHub) ResourceGet(tenantId, rid string) (*entities.Resource, error) {
	return c.coredb.ResourceGet(tenantId, rid)
}

func (c *CoreHub) ResourceDel(tenantId, rid string) error {
	return c.coredb.ResourceDel(tenantId, rid)
}

func (c *CoreHub) ResourceList(tenantId string) ([]*entities.Resource, error) {
	return c.coredb.ResourceList(tenantId)
}

func (c *CoreHub) ResourcesMulti(tenantId string, rids ...string) ([]*entities.Resource, error) {
	return c.coredb.ResourcesMulti(tenantId, rids...)
}

func (c *CoreHub) ResourcesByTarget(tenantId string, target string) ([]*entities.Resource, error) {
	return c.coredb.ResourcesByTarget(tenantId, target)
}

func (c *CoreHub) ResourceListByPlug(tenantId string, plugId string) ([]*entities.Resource, error) {
	return c.coredb.ResourceListByPlug(tenantId, plugId)
}

// agent extra

// link
func (c *CoreHub) AgentLinkNew(tenantId string, data *entities.AgentLink) error {
	return c.coredb.AgentLinkNew(tenantId, data)
}

func (c *CoreHub) AgentLinkUpdate(tenantId, pid, aid string, id int64, data map[string]any) error {
	return c.coredb.AgentLinkUpdate(tenantId, pid, aid, id, data)
}

func (c *CoreHub) AgentLinkGet(tenantId, pid, aid string, id int64) (*entities.AgentLink, error) {
	return c.coredb.AgentLinkGet(tenantId, pid, aid, id)
}

func (c *CoreHub) AgentLinkDel(tenantId, pid, aid string, id int64) error {
	return c.coredb.AgentLinkDel(tenantId, pid, aid, id)
}

func (c *CoreHub) AgentLinkList(tenantId, pid, aid string) ([]*entities.AgentLink, error) {
	return c.coredb.AgentLinkList(tenantId, pid, aid)
}

func (c *CoreHub) AgentLinkListReverse(tenantId, pid, aid string) ([]*entities.AgentLink, error) {
	return c.coredb.AgentLinkListReverse(tenantId, pid, aid)
}

// extension
func (c *CoreHub) AgentExtensionNew(tenantId string, data *entities.AgentExtension) error {
	return c.coredb.AgentExtensionNew(tenantId, data)
}

func (c *CoreHub) AgentExtensionUpdate(tenantId, pid, aid string, id int64, data map[string]any) error {
	return c.coredb.AgentExtensionUpdate(tenantId, pid, aid, id, data)
}

func (c *CoreHub) AgentExtensionGet(tenantId, pid, aid string, id int64) (*entities.AgentExtension, error) {
	return c.coredb.AgentExtensionGet(tenantId, pid, aid, id)
}

func (c *CoreHub) AgentExtensionDel(tenantId, pid, aid string, id int64) error {
	return c.coredb.AgentExtensionDel(tenantId, pid, aid, id)
}

func (c *CoreHub) AgentExtensionList(tenantId, pid, aid string) ([]*entities.AgentExtension, error) {
	return c.coredb.AgentExtensionList(tenantId, pid, aid)
}

// resource
func (c *CoreHub) AgentResourceNew(tenantId string, data *entities.AgentResource) error {
	return c.coredb.AgentResourceNew(tenantId, data)
}

func (c *CoreHub) AgentResourceUpdate(tenantId, pid, aid, slug string, data map[string]any) error {
	return c.coredb.AgentResourceUpdate(tenantId, pid, aid, slug, data)
}

func (c *CoreHub) AgentResourceGet(tenantId, pid, aid, slug string) (*entities.AgentResource, error) {
	return c.coredb.AgentResourceGet(tenantId, pid, aid, slug)
}

func (c *CoreHub) AgentResourceDel(tenantId, pid, aid, slug string) error {
	return c.coredb.AgentResourceDel(tenantId, pid, aid, slug)
}

func (c *CoreHub) AgentResourceList(tenantId, pid, aid string) ([]*entities.AgentResource, error) {
	return c.coredb.AgentResourceList(tenantId, pid, aid)
}

func (c *CoreHub) ResourceListByAgent(tenantId string, pid, aid string) ([]*entities.Resource, error) {
	return c.coredb.ResourceListByAgent(tenantId, pid, aid)
}

func (c *CoreHub) ListTargetHookByPlug(tenantId, plug string) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHookByPlug(tenantId, plug)
}

func (c *CoreHub) AgentLinkListByPlug(tenantId, pid string) ([]*entities.AgentLink, error) {
	return c.coredb.AgentLinkListByPlug(tenantId, pid)
}

func (c *CoreHub) ListTargetAppByPlug(tenantId, plug string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByPlug(tenantId, plug)
}

func (c *CoreHub) AgentExtensionListByPlug(tenantId, pid string) ([]*entities.AgentExtension, error) {
	return c.coredb.AgentExtensionListByPlug(tenantId, pid)
}
