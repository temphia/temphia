package hubv2

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (h *HubV2) instanceData(handle Handle, item xbprint.InstallItem) error {
	tenantId := handle.opts.UserSession.TenantId
	dsrc := h.dtable.DefaultSource(tenantId)

	out, err := h.pacman.BprintGetBlob(tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	schema := DataSchemaV2{}
	err = json.Unmarshal(out, &schema)
	if err != nil {
		return err
	}

	dsrc.MigrateSchema(tenantId, dyndb.MigrateOptions{
		Steps: schema.Steps,
		New:   true,
		Slug:  "",
	})

	// handle.opts

	return nil
}
