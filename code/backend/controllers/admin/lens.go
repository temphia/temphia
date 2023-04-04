package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type LogQuery struct {
	From    string            `json:"from,omitempty"`
	To      string            `json:"to,omitempty"`
	Filters map[string]string `json:"filters,omitempty"`
	Cursor  string            `json:"cursor,omitempty"`
}

func (c *Controller) LensQuery(uclaim *claim.Session, query LogQuery) ([]logx.Log, error) {
	return c.log.Query(uclaim.TenantId, logx.QueryRequest{
		From:    query.From,
		To:      query.To,
		Cursor:  query.Cursor,
		Filters: query.Filters,
	})
}
