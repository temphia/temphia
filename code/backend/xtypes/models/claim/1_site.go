package claim

import (
	"github.com/rs/xid"
)

type Site struct {
	TenentId        string            `json:"tenent_id,omitempty"`
	Type            string            `json:"type,omitempty"`
	XID             string            `json:"xid,omitempty"`
	Host            string            `json:"host,omitempty"`
	Scopes          []string          `json:"scopes,omitempty"`
	PinnedUserGroup string            `json:"ugroup,omitempty"`
	Attributes      map[string]string `json:"attributes,omitempty"`
}

func NewSiteClaim(tenantId, host string, scopes ...string) *Site {
	return &Site{
		TenentId:   tenantId,
		Type:       "site_claim",
		XID:        xid.New().String(),
		Attributes: make(map[string]string),
		Host:       host,
		Scopes:     scopes,
	}
}
