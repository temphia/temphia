package instancers

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"gopkg.in/yaml.v2"
)

var (
	gFunc                      = xtypes.GetSlugGenerator(5)
	_     xinstancer.Instancer = (*instancer)(nil)
)

type instancer struct {
	corehub store.CoreHub
	bstore  xpacman.BStore
	datahub dyndb.DataHub
}

func New(corehub store.CoreHub, bstore xpacman.BStore, datahub dyndb.DataHub) *instancer {
	return &instancer{
		corehub: corehub,
		datahub: datahub,
		bstore:  bstore,
	}
}

func (i *instancer) Instance(opts xinstancer.Options) (*xinstancer.Response, error) {
	as, err := i.loadAppSchema(opts.TenantId, opts.BprintId)
	if err != nil {
		return nil, err
	}

	if opts.PlugId == "" {
		opts.PlugId = gFunc()
	}

	err = i.corehub.PlugNew(opts.TenantId, &entities.Plug{
		Id:               opts.PlugId,
		Name:             as.Name,
		BprintId:         opts.BprintId,
		InstancedObjects: entities.JsonStrMap{},
		BprintItemId:     "",
		StepHead:         "",
		TenantId:         opts.TenantId,
	})
	if err != nil {
		return nil, err
	}

	resp, err := i.runAppBundleStep(as, opts)
	if err != nil {
		return nil, err
	}

	// update plug here

	err = i.corehub.PlugUpdate(opts.TenantId, opts.PlugId, map[string]any{
		"instanced_objects": entities.JsonStrMap(resp.Items),
		"step_head":         resp.StepHead,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *instancer) Upgrade(opts xinstancer.Options) error {

	as, err := i.loadAppSchema(opts.TenantId, opts.NextBprintId)
	if err != nil {
		return err
	}

	if opts.PlugId == "" {
		opts.PlugId = gFunc()
	}

	plug, err := i.corehub.PlugGet(opts.TenantId, opts.PlugId)
	if err != nil {
		return err
	}

	opts.InstancedIds = plug.InstancedObjects

	// fixme => pass last step_head

	i.runAppBundleStep(as, opts)

	return nil

}

// private

func (i *instancer) readMigration(tenantId, bprintid, file string) (xpackage.MigrateOptions, error) {

	return xpackage.MigrateOptions{}, nil
}

func (i *instancer) loadAppSchema(tenantId, bprintid string) (*xpackage.AppSchema, error) {

	as := &xpackage.AppSchema{}
	out, err := i.loadBprintFile(tenantId, bprintid, "", "app.json")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(out, as)
	if err != nil {
		return nil, err
	}

	return as, nil
}

func (i *instancer) loadBprintFile(tenantId, bid, folder, file string) ([]byte, error) {

	out, err := i.bstore.GetBlob(tenantId, bid, folder, file)
	if err != nil {
		return nil, err
	}

	return out, err
}
