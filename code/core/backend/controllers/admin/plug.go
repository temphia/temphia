package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"
)

// plug
func (c *Controller) PlugNew(uclaim *claim.Session, pg *entities.Plug) error {
	return c.coredb.PlugNew(uclaim.TenentId, pg)
}

func (c *Controller) PlugUpdate(uclaim *claim.Session, pid string, data map[string]any) error {
	return c.coredb.PlugUpdate(uclaim.TenentId, pid, data)
}

func (c *Controller) PlugGet(uclaim *claim.Session, pid string) (*entities.Plug, error) {
	return c.coredb.PlugGet(uclaim.TenentId, pid)
}

func (c *Controller) PlugDel(uclaim *claim.Session, pid string) error {
	return c.coredb.PlugDel(uclaim.TenentId, pid)
}

func (c *Controller) PlugList(uclaim *claim.Session) ([]*entities.Plug, error) {
	// fixme

	return c.coredb.PlugList(uclaim.TenentId)
}

func (c *Controller) PlugListByBprint(uclaim *claim.Session, bid string) ([]*entities.Plug, error) {
	return c.coredb.PlugListByBprint(uclaim.TenentId, bid)
}

// agent

func (c *Controller) AgentNew(uclaim *claim.Session, data *entities.Agent) error {
	return c.coredb.AgentNew(uclaim.TenentId, data)
}

func (c *Controller) AgentUpdate(uclaim *claim.Session, pid string, aid string, data map[string]any) error {
	return c.coredb.AgentUpdate(uclaim.TenentId, pid, aid, data)
}

func (c *Controller) AgentGet(uclaim *claim.Session, pid, agentId string) (*entities.Agent, error) {
	return c.coredb.AgentGet(uclaim.TenentId, pid, agentId)
}

func (c *Controller) AgentDel(uclaim *claim.Session, pid, agentId string) error {
	return c.coredb.AgentDel(uclaim.TenentId, pid, agentId)
}

func (c *Controller) AgentList(uclaim *claim.Session, pid string) ([]*entities.Agent, error) {
	// fixme =>

	return c.coredb.AgentList(uclaim.TenentId, pid)
}

// resource

func (c *Controller) ResourceNew(uclaim *claim.Session, data *entities.Resource) error {
	return c.coredb.ResourceNew(uclaim.TenentId, data)
}

func (c *Controller) ResourceUpdate(uclaim *claim.Session, rid string, data map[string]any) error {
	return c.coredb.ResourceUpdate(uclaim.TenentId, rid, data)
}

func (c *Controller) ResourceGet(uclaim *claim.Session, rid string) (*entities.Resource, error) {
	return c.coredb.ResourceGet(uclaim.TenentId, rid)
}

func (c *Controller) ResourceDel(uclaim *claim.Session, rid string) error {
	return c.coredb.ResourceDel(uclaim.TenentId, rid)
}

func (c *Controller) ResourceList(uclaim *claim.Session) ([]*entities.Resource, error) {
	return c.coredb.ResourceList(uclaim.TenentId)
}

func (c *Controller) ResourceListByPlug(uclaim *claim.Session, plugId string) ([]*entities.Resource, error) {
	return c.coredb.ResourceListByPlug(uclaim.TenentId, plugId)
}

func (c *Controller) ResourceAgentList(uclaim *claim.Session, req *vmodels.ResourceQuery) ([]*entities.Resource, error) {
	// agent, err := c.coredb.AgentGet(uclaim.TenentId, req.PlugId, req.AgentId)
	// if err != nil {
	// 	return nil, err
	// }

	// // vals := make([]string, 0, len(agent.Resources))
	// // for _, v := range agent.Resources {
	// // 	vals = append(vals, v)
	// // }

	// resources, err := c.coredb.ResourcesMulti(uclaim.TenentId, vals...)
	// return resources, err

	return nil, nil
}
