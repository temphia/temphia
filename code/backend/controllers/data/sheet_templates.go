package data

import (
	_ "embed"
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
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

func (c *Controller) InstanceSheet(uclaim *claim.Session, req QuickSheetInstance) error {

	tpl, ok := Templates[req.Template]
	if !ok {
		return easyerr.NotFound()
	}

	slug, err := xutils.GenerateRandomString(5)
	if err != nil {
		return err
	}

	if req.Source == "" {
		req.Source = c.dynHub.DefaultSource(uclaim.TenantId).Name()
	}

	return c.repoman.GetInstanceHub().SheetTemplate(uclaim.TenantId, req.Source, slug, &tpl)
}
