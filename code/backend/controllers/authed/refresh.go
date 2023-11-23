package authed

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type RefreshReq struct {
	UserToken string         `json:"user_token,omitempty"`
	Options   map[string]any `json:"options,omitempty"`
	OldToken  string         `json:"old_token,omitempty"`
}

type RefreshResp struct {
	Token  string `json:"token,omitempty"`
	Expiry string `json:"expiry,omitempty"`
}

func (c *Controller) RefreshService(uclaim *claim.User, opts RefreshReq) (*RefreshResp, error) {
	return c.sessionClaim(uclaim, opts)
}

func (c *Controller) sessionClaim(uclaim *claim.User, opts RefreshReq) (*RefreshResp, error) {

	_, err := c.coredb.GetUserDevice(uclaim.TenantId, uclaim.UserID, uclaim.DeviceId)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &RefreshResp{
		Token:  token,
		Expiry: "",
	}, nil

}
