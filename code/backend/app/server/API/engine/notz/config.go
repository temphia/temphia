package notz

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/app/server/API/engine/router"
)

func (a *Notz) getRouteConfig(tenantId, plugId, agentId string) *routeCache {

	key := tenantId + plugId + agentId

	a.rMutext.Lock()
	conf := a.routesCaches[key]
	a.rMutext.Unlock()

	if conf != nil {
		return conf
	}

	plug := a.ecache.GetPlug(tenantId, plugId)

	agent := a.ecache.GetAgent(tenantId, plugId, agentId)
	r := agent.WebOptions["router"]
	if r == "" {
		return nil
	}

	store := a.pacman.GetBprintFileStore()

	out, err := store.GetBlob(tenantId, plug.BprintId, "", r)
	if err != nil {
		return nil
	}

	conf = &routeCache{
		bprintId: plug.BprintId,
		config:   &router.RouteConfig{},
	}

	err = json.Unmarshal(out, conf.config)
	if err != nil {
		return nil
	}

	a.rMutext.Lock()
	c2, ok := a.routesCaches[key]
	if ok {
		a.rMutext.Unlock()
		return c2
	}
	a.routesCaches[key] = conf
	a.rMutext.Unlock()

	return conf

}
