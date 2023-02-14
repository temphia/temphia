package target

import (
	"fmt"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type TargetAppInstancer struct {
	app     xtypes.App
	pacman  repox.Hub
	corehub store.CoreHub
}

func NewTApp(app xtypes.App) xinstance.Instancer {

	deps := app.GetDeps()

	return &TargetAppInstancer{
		app:     app,
		pacman:  deps.RepoHub().(repox.Hub),
		corehub: deps.CoreHub().(store.CoreHub),
	}
}

func (ta *TargetAppInstancer) Instance(opts xinstance.Options) (*xinstance.Response, error) {

	tapp := entities.TargetApp{}
	err := opts.Handle.LoadFile(opts.File, &tapp)
	if err != nil {
		return nil, err
	}

	target := ""
	plugId := ""

	switch tapp.TargetType {
	case entities.TargetAppTypeDataTableWidget:
		tparts := strings.Split(tapp.Target, "/")
		if len(tparts) != 2 {
			return nil, easyerr.Error("target has invalid target value")
		}
		resp := opts.Handle.GetPrevObject(tparts[0])
		if resp == nil || !resp.Ok {
			return nil, easyerr.Error("could not load prev data ref")
		}
		target = fmt.Sprintf("%s/%s", resp.Slug, tparts[1])

		resp = opts.Handle.GetPrevObject(tapp.PlugId)
		if resp == nil || !resp.Ok {
			return nil, easyerr.Error("could not load prev plug ref")
		}
		plugId = resp.Slug
	}

	err = ta.corehub.AddTargetApp(&entities.TargetApp{
		Id:          0,
		Name:        tapp.Name,
		Icon:        tapp.Icon,
		Policy:      tapp.Policy,
		TargetType:  tapp.TargetType,
		Target:      target,
		ContextType: tapp.ContextType,
		PlugId:      plugId,
		AgentId:     tapp.AgentId,
		ExecDomain:  0,
		ExecMeta:    tapp.ExecMeta,
		ExtraMeta:   tapp.ExtraMeta,
		TenantId:    opts.TenantId,
	})
	if err != nil {
		return nil, err
	}

	return &xinstance.Response{
		Ok:      false,
		Message: "",
		Slug:    fmt.Sprintf("%d", 0), // fixme => return newly created id
		Data:    nil,
	}, nil
}
