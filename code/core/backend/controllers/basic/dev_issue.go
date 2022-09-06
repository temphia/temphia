package basic

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/controllers/admin/devtoken"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type DevIssueReq struct {
	AllPlugs bool     `json:"all_plugs,omitempty"`
	PlugIds  []string `json:"plug_ids,omitempty"`
	BprintId string   `json:"bprint_id,omitempty"`
	Encoded  bool     `json:"encoded,omitempty"`
}

func (c *Controller) DevIssueTktEncoded(uclaim *claim.Session, host string, req DevIssueReq) (string, error) {

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

	// fixme => check perms

	tok, err := c.signer.SignPlugDevTkt(uclaim.TenentId, &claim.PlugDevTkt{
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
		TenantId:  uclaim.TenentId,
		DevTicket: tok,
	}, nil

}
