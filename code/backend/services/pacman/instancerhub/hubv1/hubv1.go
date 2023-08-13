package hubv1

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/services/pacman/instancerhub/instancers/sheet"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

var (
	_ repox.InstancerHubV1 = (*InstancHub)(nil)
)

type InstancHub struct {
	instancers map[string]xinstance.Instancer
	pacman     repox.BStore
}

/*

	instancehub
		- manual
			- bundle -> instancer
			- single -> instancer
		- automatic
			- bundle -> instancer
			- single -> instancer

*/

func New(instancers map[string]xinstance.Instancer, pacman repox.BStore) *InstancHub {
	return &InstancHub{
		instancers: instancers,
		pacman:     pacman,
	}
}

func (i *InstancHub) Instance(opts repox.InstanceOptionsV1) (any, error) {

	switch opts.InstancerType {
	case xbprint.TypeBundle:
		if opts.Auto {
			return i.automaticBundle(opts)
		}
		return i.manualBundleItem(opts)

	default:
		if opts.Auto {
			return i.automaticSingle(opts)
		}
		return i.manualSingle(opts)
	}

}

func (i *InstancHub) InstanceSheetDirect(opts repox.InstanceSheetOptions) (*xinstance.Response, error) {
	return i.sheetTemplate(opts.UserContext.TenantId, opts.Source, opts.Group, opts.Template)
}

func (i *InstancHub) sheetTemplate(tenantId, source, gslug string, template *xbprint.NewSheetGroup) (*xinstance.Response, error) {
	sintancer := i.instancers[xbprint.TypeDataSheet].(*sheet.SheetInstancer)
	return sintancer.DirectInstance(tenantId, source, gslug, template)
}

func (i *InstancHub) manualSingle(opts repox.InstanceOptionsV1) (any, error) {

	pp.Println("INSTANCE OPTS |>", opts)

	instancer, ok := i.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound("instancer1")
	}

	return instancer.Instance(xinstance.Options{
		TenantId:     opts.UserContext.TenantId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserContext.UserID,
		UserData:     opts.UserConfigData,
		Automatic:    false,
		Handle: &Handle{
			instanced: make(map[string]*xinstance.Response),
			opts:      opts,
			pacman:    i.pacman,
		},
	})
}

func (i *InstancHub) manualBundleItem(opts repox.InstanceOptionsV1) (any, error) {

	instancer, ok := i.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound("instancer2")
	}

	return instancer.Instance(xinstance.Options{
		TenantId:     opts.UserContext.TenantId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserContext.UserID,
		UserData:     opts.UserConfigData,
		Automatic:    false,
		Handle: &Handle{
			instanced: make(map[string]*xinstance.Response),
			opts:      opts,
			pacman:    i.pacman,
		},
	})

}

func (i *InstancHub) automaticBundle(opts repox.InstanceOptionsV1) (any, error) {
	bundle := xbprint.BundleV1{}
	err := i.loadFile(opts.UserContext.TenantId, opts.BprintId, opts.File, &bundle)
	if err != nil {
		return nil, err
	}

	iObjs := make(map[string]*xinstance.Response)

	pp.Println("ALL BUNDLE ||>>", bundle)

	allok := true

	for _, bitem := range bundle.Items {
		pp.Println("INSTANCING BITEM", bitem)

		instancer, ok := i.instancers[bitem.Type]
		if !ok {
			pp.Println("NOT FOUND", bitem.Type)
			return nil, easyerr.NotFound("instancer3")
		}

		resp, err := instancer.Instance(xinstance.Options{
			TenantId:     opts.UserContext.TenantId,
			BprintId:     opts.BprintId,
			InstanceType: bitem.Type,
			File:         bitem.File,
			UserId:       opts.UserContext.UserID,
			UserData:     nil,
			Automatic:    true,
			Handle: &Handle{
				instanced: iObjs,
				opts:      opts,
				pacman:    i.pacman,
			},
		})
		if err != nil {
			allok = false

			iObjs[bitem.Name] = &xinstance.Response{
				Ok:      false,
				Message: err.Error(),
				Type:    bitem.Type,
			}

			continue
		}

		resp.Type = bitem.Type
		iObjs[bitem.Name] = resp
	}

	return AutoResp{
		AllOk:   allok,
		Objects: iObjs,
	}, nil
}

func (i *InstancHub) automaticSingle(opts repox.InstanceOptionsV1) (any, error) {

	instancer, ok := i.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound("instancer4")
	}

	iresp, err := instancer.Instance(xinstance.Options{
		TenantId:     opts.UserContext.TenantId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserContext.UserID,
		UserData:     opts.UserConfigData,
		Automatic:    true,
		Handle: &Handle{
			instanced: make(map[string]*xinstance.Response),
			opts:      opts,
			pacman:    i.pacman,
		},
	})

	if err != nil {
		return nil, err
	}

	iresp.Type = opts.InstancerType

	return AutoResp{
		AllOk: err == nil,
		Objects: map[string]*xinstance.Response{
			"default": iresp,
		},
	}, nil

}

// private

func (i *InstancHub) loadFile(tenantId, bid string, file string, target any) error {
	out, err := i.pacman.GetBlob(tenantId, bid, "", file)
	if err != nil {
		return err
	}

	return json.Unmarshal(out, target)
}
