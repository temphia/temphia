package admin

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/controllers/admin/devtoken"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/scopes"
)

type DevIssueReq struct {
	AllPlugs bool     `json:"all_plugs,omitempty"`
	PlugIds  []string `json:"plug_ids,omitempty"`
	BprintId string   `json:"bprint_id,omitempty"`
}

func (c *Controller) DevIssueTktEncoded(uclaim *claim.Session, host string, req DevIssueReq) (string, error) {
	if !c.HasScope(uclaim, "engine") {
		return "", scopes.ErrNoAdminEngineScope
	}

	// ErrNoAdminLensScope

	pd, err := c.DevIssueTkt(uclaim, host, req)
	if err != nil {
		return "", err
	}

	ftok, err := pd.ToString()
	if err != nil {
		return "", err
	}

	return ftok, nil
}

func (c *Controller) DevIssueTkt(uclaim *claim.Session, host string, req DevIssueReq) (*devtoken.Plug, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	// fixme => check perms

	tok, err := c.signer.SignPlugDevTkt(uclaim.TenantId, &claim.PlugDevTkt{
		UserId:    uclaim.UserID,
		UserGroup: uclaim.UserGroup,
		BprintId:  req.BprintId,
		PlugIds:   req.PlugIds,
		AllPlugs:  req.AllPlugs,
	})

	if err != nil {
		return nil, err
	}

	return &devtoken.Plug{
		HostAddrs: []string{fmt.Sprintf("http://%s", host)},
		TenantId:  uclaim.TenantId,
		DevTicket: tok,
	}, nil

}
