package data

import (
	_ "embed"
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

//go:embed sheet_templates.json
var tplbytes []byte

var Templates = map[string]xbprint.NewSheetGroup{}

func init() {

	json.Unmarshal(tplbytes, &Templates)
}

func (c *Controller) ListSheetTemplates(uclaim *claim.Session) (map[string]xbprint.NewSheetGroup, error) {
	return Templates, nil
}

type QuickSheetInstance struct {
	Name     string `json:"name,omitempty"`
	Info     string `json:"info,omitempty"`
	Template string `json:"template,omitempty"`
	Source   string `json:"source,omitempty"`
}

func (c *Controller) InstanceSheet(uclaim *claim.Session, req QuickSheetInstance) (*xinstance.Response, error) {

	tpl, ok := Templates[req.Template]
	if !ok {
		return nil, easyerr.NotFound("template")
	}

	tpl.Name = req.Name
	tpl.Info = req.Info

	return c.repoman.GetInstancerHubV1().InstanceSheetDirect(repox.InstanceSheetOptions{
		Source:      req.Source,
		Group:       "",
		Template:    &tpl,
		UserContext: uclaim.AsUserCtx(),
	})

}
