package admin

import (
	"sync"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities/flowmap"
	"github.com/temphia/temphia/code/backend/xtypes/models/vmodels"
	"github.com/temphia/temphia/code/backend/xtypes/scopes"
)

// plug
func (c *Controller) PlugNew(uclaim *claim.Session, pg *entities.Plug) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.PlugNew(uclaim.TenantId, pg)
}

func (c *Controller) PlugUpdate(uclaim *claim.Session, pid string, data map[string]any) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.PlugUpdate(uclaim.TenantId, pid, data)
}

func (c *Controller) PlugGet(uclaim *claim.Session, pid string) (*entities.Plug, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.PlugGet(uclaim.TenantId, pid)
}

func (c *Controller) PlugDel(uclaim *claim.Session, pid string) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.PlugDel(uclaim.TenantId, pid)
}

func (c *Controller) PlugList(uclaim *claim.Session) ([]*entities.Plug, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.PlugList(uclaim.TenantId, nil)
}

func (c *Controller) PlugListByBprint(uclaim *claim.Session, bid string) ([]*entities.Plug, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.PlugListByBprint(uclaim.TenantId, bid)
}

// agent

func (c *Controller) AgentNew(uclaim *claim.Session, data *entities.Agent) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.AgentNew(uclaim.TenantId, data)
}

func (c *Controller) AgentUpdate(uclaim *claim.Session, pid string, aid string, data map[string]any) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.AgentUpdate(uclaim.TenantId, pid, aid, data)
}

func (c *Controller) AgentGet(uclaim *claim.Session, pid, agentId string) (*entities.Agent, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.AgentGet(uclaim.TenantId, pid, agentId)
}

func (c *Controller) AgentDel(uclaim *claim.Session, pid, agentId string) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.AgentDel(uclaim.TenantId, pid, agentId)
}

func (c *Controller) AgentList(uclaim *claim.Session, pid string) ([]*entities.Agent, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.AgentList(uclaim.TenantId, pid)
}

// resource

func (c *Controller) ResourceNew(uclaim *claim.Session, data *entities.Resource) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.ResourceNew(uclaim.TenantId, data)
}

func (c *Controller) ResourceUpdate(uclaim *claim.Session, rid string, data map[string]any) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.ResourceUpdate(uclaim.TenantId, rid, data)
}

func (c *Controller) ResourceGet(uclaim *claim.Session, rid string) (*entities.Resource, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.ResourceGet(uclaim.TenantId, rid)
}

func (c *Controller) ResourceDel(uclaim *claim.Session, rid string) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.coredb.ResourceDel(uclaim.TenantId, rid)
}

func (c *Controller) ResourceList(uclaim *claim.Session) ([]*entities.Resource, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.ResourceList(uclaim.TenantId, nil)
}

type FlowmapData struct {
	Plug       *entities.Plug                      `json:"plug,omitempty"`
	Agents     []*entities.Agent                   `json:"agents,omitempty"`
	AgentLinks map[string]*entities.AgentLink      `json:"agent_links,omitempty"`
	AgentExts  map[string]*entities.AgentExtension `json:"agent_exts,omitempty"`

	TargetApps     []*entities.TargetApp         `json:"target_apps,omitempty"`
	TargetHooks    []*entities.TargetHook        `json:"target_hooks,omitempty"`
	AgentResources map[string]*entities.Resource `json:"agent_resources,omitempty"`
}

func (c *Controller) PlugFlowmap2(uclaim *claim.Session, plugId string) (*flowmap.Data, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.coredb.GetFlowMap(uclaim.TenantId)
}

func (c *Controller) PlugFlowmap(uclaim *claim.Session, plugId string) (*FlowmapData, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	data := FlowmapData{
		Plug:           nil,
		Agents:         nil,
		AgentLinks:     make(map[string]*entities.AgentLink),
		AgentExts:      make(map[string]*entities.AgentExtension),
		TargetApps:     nil,
		TargetHooks:    nil,
		AgentResources: make(map[string]*entities.Resource),
	}

	var wg sync.WaitGroup
	wg.Add(7)

	go func() {
		defer wg.Done()
		plug, _ := c.coredb.PlugGet(uclaim.TenantId, plugId)
		data.Plug = plug
	}()

	go func() {
		defer wg.Done()
		agents, _ := c.coredb.AgentList(uclaim.TenantId, plugId)
		data.Agents = agents
	}()

	go func() {
		defer wg.Done()
		links, err := c.coredb.AgentLinkListByPlug(uclaim.TenantId, plugId)
		if err != nil {
			return
		}

		for _, link := range links {
			data.AgentLinks[link.FromAgent] = link
		}
	}()

	go func() {
		defer wg.Done()
		exts, err := c.coredb.AgentExtensionListByPlug(uclaim.TenantId, plugId)
		if err != nil {
			return
		}

		for _, ext := range exts {
			data.AgentExts[ext.Agent] = ext
		}
	}()

	go func() {
		defer wg.Done()
		hooks, _ := c.coredb.ListTargetHookByPlug(uclaim.TenantId, plugId)
		data.TargetHooks = hooks
	}()

	go func() {
		defer wg.Done()
		apps, _ := c.coredb.ListTargetAppByPlug(uclaim.TenantId, plugId)
		data.TargetApps = apps
	}()

	go func() {
		defer wg.Done()
		data.AgentResources = map[string]*entities.Resource{} // fixme
	}()

	wg.Wait()

	pp.Println(data)

	return &data, nil

}

func (c *Controller) ResourceAgentList(uclaim *claim.Session, req *vmodels.ResourceQuery) ([]*entities.Resource, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	// agent, err := c.coredb.AgentGet(uclaim.TenantId, req.PlugId, req.AgentId)
	// if err != nil {
	// 	return nil, err
	// }

	// // vals := make([]string, 0, len(agent.Resources))
	// // for _, v := range agent.Resources {
	// // 	vals = append(vals, v)
	// // }

	// resources, err := c.coredb.ResourcesMulti(uclaim.TenantId, vals...)
	// return resources, err

	return nil, nil
}
