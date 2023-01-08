package repohub

import (
	"encoding/json"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox/xinstance"
)

func (p *PacMan) GetInstanceHub() repox.InstancHub {

	return nil
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

func (i *InstancHub) ManualSingle(opt repox.InstanceOptions) (any, error) {
	instancer, ok := i.pacman.instancers[opt.InstancerType]
	if !ok {
		return nil, easyerr.NotFound()
	}

	return instancer.Instance(xinstance.Options{
		TenantId:     opt.UserSession.TenentId,
		BprintId:     opt.BprintId,
		InstanceType: opt.InstancerType,
		File:         opt.File,
		UserId:       opt.UserSession.UserID,
		UserData:     opt.UserConfigData,
		Automatic:    false,
	})
}

func (i *InstancHub) ManualBundleItem(opts repox.InstanceOptions) (any, error) {

	instancer, ok := i.pacman.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound()
	}

	return instancer.Instance(xinstance.Options{
		TenantId:     opts.UserSession.TenentId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserSession.UserID,
		UserData:     opts.UserConfigData,
		Automatic:    false,
	})

}

func (i *InstancHub) AutomaticBundle(opts repox.InstanceOptions) (any, error) {
	bundle := xbprint.Bundle{}
	err := i.pacman.loadFile(opts.UserSession.TenentId, opts.BprintId, opts.File, opts.RepoId, &bundle)
	if err != nil {
		return nil, err
	}

	iObjs := make(map[string]*xinstance.Response)

	for _, bitem := range bundle.Items {
		instancer, ok := i.pacman.instancers[bitem.Type]
		if !ok {
			return nil, easyerr.NotFound()
		}

		resp, err := instancer.Instance(xinstance.Options{
			TenantId:     opts.UserSession.TenentId,
			BprintId:     opts.BprintId,
			InstanceType: opts.InstancerType,
			File:         opts.File,
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
			return nil, err
		}

		iObjs[bitem.Name] = resp
	}

	return nil, nil
}

func (i *InstancHub) AutomaticSingle(opts repox.InstanceOptions) (any, error) {

	instancer, ok := i.pacman.instancers[opts.InstancerType]
	if !ok {
		return nil, easyerr.NotFound()
	}

	return instancer.Instance(xinstance.Options{
		TenantId:     opts.UserSession.TenentId,
		BprintId:     opts.BprintId,
		InstanceType: opts.InstancerType,
		File:         opts.File,
		UserId:       opts.UserSession.UserID,
		UserData:     opts.UserConfigData,
		Automatic:    true,
	})
}

func (p *PacMan) loadFile(tenantId, bid string, file string, repoId int64, target any) error {

	out, err := p.RepoSourceGetBlob(tenantId, "", bid, repoId, file)
	if err != nil {
		return err
	}
	return json.Unmarshal(out, target)
}
