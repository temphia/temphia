package xinstancer

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

type Options struct {
	BprintId    string             `json:"bprint_id,omitempty"`
	InstanceId  string             `json:"instance_id,omitempty"`
	Items       map[string]string  `json:"items,omitempty"`
	UserSession *claim.UserContext `json:"-"`
}

type SheetOptions struct {
	Source      string                 `json:"source,omitempty"`
	Group       string                 `json:"group,omitempty"`
	Template    *xbprint.NewSheetGroup `json:"template,omitempty"`
	UserContext *claim.UserContext     `json:"-"`
}

type Response struct {
	Items map[string]string `json:"items,omitempty"`
}

type Instancer interface {
	Instance(opts Options) (*Response, error)
	Upgrade(opts Options) error
	InstanceSheetDirect(opts SheetOptions) (*xinstance.Response, error)
}
