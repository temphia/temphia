package resource

import (
	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type ResInstancer struct {
	app    xtypes.App
	pacman repox.Hub
	syncer store.SyncDB
}

func New(app xtypes.App) xinstance.Instancer {

	deps := app.GetDeps()

	return &ResInstancer{
		app:    app,
		pacman: deps.RepoHub().(repox.Hub),
		syncer: deps.CoreHub().(store.CoreHub),
	}
}

func (pi *ResInstancer) Instance(opts xinstance.Options) (*xinstance.Response, error) {
	res := xbprint.NewResource{}
	err := opts.Handle.LoadFile(opts.File, &res)
	if err != nil {
		return nil, err
	}

	id := xid.New().String()

	pp.Println("@resource |>", res)

	err = pi.syncer.ResourceNew(opts.TenantId, &entities.Resource{
		Id:        id,
		Name:      res.Name,
		Type:      res.Type,
		SubType:   res.SubType,
		Target:    "",
		Payload:   res.Payload,
		Policy:    res.Policy,
		PlugId:    "",
		ExtraMeta: res.Meta,
		TenantId:  opts.TenantId,
	})

	if err != nil {
		return nil, err
	}

	return &xinstance.Response{
		Ok:      true,
		Message: "",
		Slug:    id,
		Data:    nil,
	}, nil
}
