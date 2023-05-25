package repox

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

type InstanceOptionsV1 struct {
	BprintId       string             `json:"bprint_id,omitempty"`
	InstancerType  string             `json:"instancer_type,omitempty"`
	File           string             `json:"file,omitempty"`
	UserConfigData []byte             `json:"data,omitempty"`
	Auto           bool               `json:"auto,omitempty"`
	UserContext    *claim.UserContext `json:"-"`
}

type InstanceSheetOptions struct {
	Source      string                 `json:"source,omitempty"`
	Group       string                 `json:"group,omitempty"`
	Template    *xbprint.NewSheetGroup `json:"template,omitempty"`
	UserContext *claim.UserContext     `json:"-"`
}

type InstancerHubV1 interface {
	Instance(opts InstanceOptionsV1) (any, error)
	InstanceSheetDirect(opts InstanceSheetOptions) (*xinstance.Response, error)
}

type InstanceOptionsV2 struct {
	BprintId    string             `json:"bprint_id,omitempty"`
	UserSession *claim.UserContext `json:"-"`
	InstanceId  string             `json:"instance_id,omitempty"`
}

type UpdateOptionsV2 struct {
	BprintId    string             `json:"bprint_id,omitempty"`
	Items       map[string]string  `json:"items,omitempty"`
	UserSession *claim.UserContext `json:"-"`
}

type InstanceResponseV2 struct {
	Items map[string]string `json:"items,omitempty"`
}

type InstancerHubV2 interface {
	Instance(opts InstanceOptionsV2) (*InstanceResponseV2, error)
	Upgrade(opts UpdateOptionsV2) error

	InstanceSheetDirect(opts InstanceSheetOptions) (*xinstance.Response, error)
}
