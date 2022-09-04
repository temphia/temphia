package operator

import (
	"github.com/k0kubun/pp"
	"github.com/rs/xid"

	"github.com/temphia/temphia/code/core/backend/controllers/operator/opsutils"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Controller struct {
	app        xtypes.App
	coredb     store.CoreHub // fixme => use control plane instead of coredb sirectly
	signer     service.Signer
	OpUser     string
	OpPassword string
	OpToken    string
}

func New(cdb store.CoreHub, signer service.Signer, app xtypes.App, opName, opPass string) *Controller {

	return &Controller{
		app:        app,
		signer:     signer,
		coredb:     cdb,
		OpUser:     opName,
		OpPassword: opPass,
		OpToken:    "",
	}
}

func (c *Controller) Login(data *vmodels.OperatorLoginReq) (*vmodels.OperatorLoginResp, error) {
	pp.Println("@@=>", c.OpUser, c.OpPassword, c.OpToken, data)

	if data.MasterOpToken != "" {
		if !c.verifyMasterToken() {
			return nil, easyerr.Error("Invalid Mastertoken")
		}
	} else {
		// fixme => security and stuff (constant time compare ?)
		if data.User != c.OpUser || data.Password != c.OpPassword {
			return nil, easyerr.Error("Invalid User crediantials")
		}
	}

	opclaim := &claim.Operator{
		XID:          xid.New().String(),
		Type:         "operator",
		BindDeviceId: "",
	}

	optoken, err := c.signer.SignOperator(opclaim)
	if err != nil {
		return nil, err
	}

	return &vmodels.OperatorLoginResp{
		Token: optoken,
	}, nil

}

func (c *Controller) AddTenant(data *vmodels.NewTenant) error {
	return opsutils.AddTenant(c.app, data)
}

func (c *Controller) UpdateTenant(slug string, data map[string]any) error {
	return c.coredb.UpdateTenant(slug, data)
}

func (c *Controller) ListTenant() ([]*entities.Tenant, error) {
	return c.coredb.ListTenant()
}

func (c *Controller) DeleteTenant(slug string) error {
	return c.coredb.RemoveTenant(slug)
}

func (c *Controller) Stats() {

}
func (c *Controller) TenantToken() {

}

// private

func (c *Controller) verifyMasterToken() bool {
	return false
}
