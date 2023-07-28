package authed

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
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

	_, err := c.coredb.GetUserDevice(uclaim.TenantId, uclaim.UserID, uclaim.DeviceId)
	if err != nil {

		pp.Println("@get_user_device_error", err)

		return &RefreshResp{
			Token:    "",
			Message:  "device not found",
			StatusOk: false,
		}
	}

	serviceId := c.sessionNode.Generate().Int64()

	if opts.OldToken != "" {
		sess, err := c.signer.ParseSession(uclaim.TenantId, opts.OldToken)
		if err != nil {
			serviceId = sess.SessionID
		}
	}

	sclaim := uclaim.DeriveSession(serviceId)

	token, err := c.signer.SignSession(uclaim.TenantId, sclaim)
	if err != nil {
		pp.Println("@sign_err", err)

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
