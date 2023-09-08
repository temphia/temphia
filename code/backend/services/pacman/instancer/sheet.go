package instancers

import (
	"encoding/json"

	_ "embed"

	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

//go:embed _sheet_schema.json
var sheetSchema []byte

func (i *instancer) InstanceSheetDirect(opts xinstancer.SheetOptions) (*xinstancer.Response, error) {
	parsedSchema := &xpackage.NewTableGroup{}

	err := json.Unmarshal(sheetSchema, parsedSchema)
	if err != nil {
		return nil, err
	}

	parsedSchema.Slug = gFunc()
	parsedSchema.Name = opts.Template.Name
	parsedSchema.Description = opts.Template.Info

	err = i.datahub.GetDynDB().NewGroup(opts.UserContext.TenantId, parsedSchema)
	if err != nil {
		return nil, err
	}

	return &xinstancer.Response{
		StepHead: "",
		Items:    map[string]string{},
	}, nil
}
