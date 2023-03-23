package data

import (
	_ "embed"
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

//go:embed sheet_templates.json
var tplbytes []byte

var Templates = map[string]xbprint.NewSheetGroup{}

func init() {

	json.Unmarshal(tplbytes, &Templates)
}

func (c *Controller) ListSheetTemplates(uclaim *claim.Data) (map[string]xbprint.NewSheetGroup, error) {
	return Templates, nil
}

type QuickSheetInstance struct {
	Name     string `json:"name,omitempty"`
	Info     string `json:"info,omitempty"`
	Template string `json:"template,omitempty"`
	Source   string `json:"source,omitempty"`
}

func (c *Controller) InstanceSheet(uclaim *claim.Data, req QuickSheetInstance) error {

	// tpl, ok := Templates[req.Template]
	// if !ok {
	// 	return easyerr.NotFound()
	// }

	// sheet.New()

	return nil
}
