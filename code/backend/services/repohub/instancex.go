package repohub

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/services/repohub/instancers/sheet"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

func (p *PacMan) GetInstanceHub() repox.InstancHub {
	return &p.instancer
}

type InstancHub struct {
	pacman *PacMan
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

func (i *InstancHub) SheetTemplate(tenantId, source, gslug string, template *xbprint.NewSheetGroup) (*xinstance.Response, error) {
	sintancer := i.pacman.instancers[xbprint.TypeDataSheet].(*sheet.SheetInstancer)
	return sintancer.DirectInstance(tenantId, source, gslug, template)
}

func (i *InstancHub) ManualSingle(opts repox.InstanceOptions) (any, error) {

	pp.Println("INSTANCE OPTS |>", opts)

	instancer, ok := i.pacman.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound()
	}

	return instancer.Instance(xinstance.Options{
		TenantId:     opts.UserSession.TenantId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserSession.UserID,
		UserData:     opts.UserConfigData,
		Automatic:    false,
		Handle: &Handle{
			instanced: make(map[string]*xinstance.Response),
			opts:      opts,
			pacman:    i.pacman,
		},
	})
}

func (i *InstancHub) ManualBundleItem(opts repox.InstanceOptions) (any, error) {

	instancer, ok := i.pacman.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound()
	}

	return instancer.Instance(xinstance.Options{
		TenantId:     opts.UserSession.TenantId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserSession.UserID,
		UserData:     opts.UserConfigData,
		Automatic:    false,
		Handle: &Handle{
			instanced: make(map[string]*xinstance.Response),
			opts:      opts,
			pacman:    i.pacman,
		},
	})

}

func (i *InstancHub) AutomaticBundle(opts repox.InstanceOptions) (any, error) {
	bundle := xbprint.Install{}
	err := i.pacman.loadFile(opts.UserSession.TenantId, opts.BprintId, opts.File, &bundle)
	if err != nil {
		return nil, err
	}

	iObjs := make(map[string]*xinstance.Response)

	pp.Println("ALL BUNDLE ||>>", bundle)

	allok := true

	for _, bitem := range bundle.Items {
		pp.Println("INSTANCING BITEM", bitem)

		instancer, ok := i.pacman.instancers[bitem.Type]
		if !ok {
			pp.Println("NOT FOUND", bitem.Type)
			return nil, easyerr.NotFound()
		}

		resp, err := instancer.Instance(xinstance.Options{
			TenantId:     opts.UserSession.TenantId,
			BprintId:     opts.BprintId,
			InstanceType: bitem.Type,
			File:         bitem.File,
			UserId:       opts.UserSession.UserID,
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

func (i *InstancHub) AutomaticSingle(opts repox.InstanceOptions) (any, error) {

	instancer, ok := i.pacman.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound()
	}

	iresp, err := instancer.Instance(xinstance.Options{
		TenantId:     opts.UserSession.TenantId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserSession.UserID,
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

func (p *PacMan) loadFile(tenantId, bid string, file string, target any) error {

	out, err := p.BprintGetBlob(tenantId, bid, file)
	if err != nil {
		return err
	}
	return json.Unmarshal(out, target)
}

type AutoResp struct {
	AllOk   bool                           `json:"all_ok"`
	Objects map[string]*xinstance.Response `json:"objects"`
}