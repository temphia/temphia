package authed

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type RefreshReq struct {
	UserToken string         `json:"user_token,omitempty"`
	Options   map[string]any `json:"options,omitempty"`
	OldToken  string         `json:"old_token,omitempty"`
}

type RefreshResp struct {
	Token    string `json:"token,omitempty"`
	Expiry   string `json:"expiry,omitempty"`
	Message  string `json:"message,omitempty"`
	StatusOk bool   `json:"status_ok,omitempty"`
}

func (c *Controller) RefreshService(uclaim *claim.User, opts RefreshReq) *RefreshResp {
	return c.refreshService(uclaim, opts)
}

func (c *Controller) refreshService(uclaim *claim.User, opts RefreshReq) *RefreshResp {
	return c.sessionClaim(uclaim, opts)
}

func (c *Controller) sessionClaim(uclaim *claim.User, opts RefreshReq) *RefreshResp {
	deviceId := xid.New().String()
	serviceId := c.sessman.SessionId()
	if opts.OldToken != "" {
		sess, err := c.signer.ParseSession(uclaim.TenentId, opts.OldToken)
		if err != nil {
			deviceId = sess.DeviceId
			serviceId = sess.SessionID
		}
	}

	sclaim := &claim.Session{
		TenentId:   uclaim.TenentId,
		UserID:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		Type:       "session",
		Attributes: nil,
		SessionID:  serviceId,
		DeviceId:   deviceId,
	}

	token, err := c.signer.SignSession(uclaim.TenentId, sclaim)
	if err != nil {
		return &RefreshResp{
			Token:    "",
			Expiry:   "",
			Message:  err.Error(),
			StatusOk: false,
		}
	}

	return &RefreshResp{
		Token:    token,
		Expiry:   "",
		Message:  "",
		StatusOk: true,
	}

}
