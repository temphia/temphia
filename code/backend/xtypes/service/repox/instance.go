package repox

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

type InstanceOptions struct {
	BprintId       string         `json:"bprint_id,omitempty"`
	InstancerType  string         `json:"instancer_type,omitempty"`
	File           string         `json:"file,omitempty"`
	UserConfigData []byte         `json:"data,omitempty"`
	Auto           bool           `json:"auto,omitempty"`
	UserSession    *claim.Session `json:"-"`
}

type InstancHub interface {
	SheetTemplate(tenantId, source, gslug string, template *xbprint.NewSheetGroup) error
	ManualSingle(opt InstanceOptions) (any, error)
	ManualBundleItem(opts InstanceOptions) (any, error)
	AutomaticBundle(opt InstanceOptions) (any, error)
	AutomaticSingle(opt InstanceOptions) (any, error)
}
