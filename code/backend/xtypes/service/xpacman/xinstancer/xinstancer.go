package xinstancer

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

type Options struct {
	BprintId    string             `json:"bprint_id,omitempty"`
	UserSession *claim.UserContext `json:"-"`
	// for upgrade
	PlugId       string            `json:"plug_id,omitempty"`
	NextBprintId string            `json:"next_bprint_id,omitempty"`
	TenantId     string            `json:"-"`
	InstancedIds map[string]string `json:"instanced,omitempty"`
}

type SheetOptions struct {
	Source      string                  `json:"source,omitempty"`
	Template    *xpackage.NewSheetGroup `json:"template,omitempty"`
	UserContext *claim.UserContext      `json:"-"`
}

type Response struct {
	StepHead string            `json:"step_head,omitempty"`
	Items    map[string]string `json:"items,omitempty"`
}

type Instancer interface {
	Instance(opts Options) (*Response, error)
	Upgrade(opts Options) error
	InstanceSheetDirect(opts SheetOptions) (*Response, error)
}
