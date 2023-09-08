package data

import (
	_ "embed"
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

//go:embed sheet_templates.json
var tplbytes []byte

var Templates = map[string]xpackage.NewSheetGroup{}

func init() {
	err := json.Unmarshal(tplbytes, &Templates)
	if err != nil {
		panic(err)
	}
}

func (c *Controller) ListSheetTemplates(uclaim *claim.Session) (map[string]xpackage.NewSheetGroup, error) {
	return Templates, nil
}

type QuickSheetInstance struct {
	Name     string `json:"name,omitempty"`
	Info     string `json:"info,omitempty"`
	Template string `json:"template,omitempty"`
	Source   string `json:"source,omitempty"`
}

func (c *Controller) InstanceSheet(uclaim *claim.Session, req QuickSheetInstance) (*xinstancer.Response, error) {

	tpl, ok := Templates[req.Template]
	if !ok {
		return nil, easyerr.NotFound("template")
	}

	tpl.Name = req.Name
	tpl.Info = req.Info

	return c.repoman.GetInstancer().InstanceSheetDirect(xinstancer.SheetOptions{
		Source:      req.Source,
		Template:    &tpl,
		UserContext: uclaim.AsUserCtx(),
	})
}
