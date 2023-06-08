package adapter

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type AdapterMod struct {
	adapterHub httpx.AdapterHub
	adapterId  int64
	inBinder   bindx.Invoker
	tenantId   string
}

func (am *AdapterMod) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {

	uctx := am.inBinder.ContextUser()
	if uctx == nil {
		return nil, easyerr.NotFound("user ctx @adpter ctx")
	}

	out, err := args.AsJsonBytes()
	if err != nil {
		return nil, err
	}

	resp, err := am.adapterHub.PreformEditorAction(httpx.AdapterEditorContext{
		Id: am.adapterId,
		User: &claim.UserContext{
			TenantId:  am.tenantId,
			UserID:    uctx.UserID,
			UserGroup: uctx.UserGroup,
			SessionID: uctx.SessionID,
			DeviceId:  uctx.DeviceId,
		},
		Name: method,
		Data: out,
	})
	if err != nil {
		return nil, err
	}

	return lazydata.NewAnyData(resp), nil
}

func (am *AdapterMod) Close() error {
	am.adapterHub = nil
	return nil
}
