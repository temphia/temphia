package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type LogQuery struct {
	From    string            `json:"from,omitempty"`
	To      string            `json:"to,omitempty"`
	Filters map[string]string `json:"filter,omitempty"`
}

func (c *Controller) LensQueryApp(uclaim *claim.Session, query LogQuery) ([]logx.Log, error) {
	return c.log.QueryAppTenant(query.From, query.To, uclaim.TenentId, query.Filters)
}

func (c *Controller) LensQueryEngine(uclaim *claim.Session, query LogQuery) ([]logx.Log, error) {
	return c.log.QueryEngine(query.From, query.To, uclaim.TenentId, query.Filters)
}

func (c *Controller) LensQuerySite(uclaim *claim.Session, query LogQuery) ([]logx.Log, error) {
	return c.log.QuerySite(query.From, query.To, uclaim.TenentId, query.Filters)
}
