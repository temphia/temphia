package hubv1

import "github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"

type AutoResp struct {
	AllOk   bool                           `json:"all_ok"`
	Objects map[string]*xinstance.Response `json:"objects"`
}
