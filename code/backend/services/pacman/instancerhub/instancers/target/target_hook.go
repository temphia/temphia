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

type TargetHookInstancer struct {
	app     xtypes.App
	pacman  repox.Pacman
	corehub store.CoreHub
}

func NewTHook(app xtypes.App) xinstance.Instancer {

	deps := app.GetDeps()

	return &TargetHookInstancer{
		app:     app,
		pacman:  deps.RepoHub().(repox.Pacman),
		corehub: deps.CoreHub().(store.CoreHub)}
}

func (th *TargetHookInstancer) Instance(opts xinstance.Options) (*xinstance.Response, error) {

	thook := entities.TargetHook{}
	err := opts.Handle.LoadFile(opts.File, &thook)
	if err != nil {
		return nil, err
	}

	target := ""
	plugId := ""

	switch thook.TargetType {
	case entities.TargetHookTypeDataTableHook:
		tparts := strings.Split(thook.Target, "/")
		if len(tparts) != 2 {
			return nil, easyerr.Error("target has invalid target value")
		}
		resp := opts.Handle.GetPrevObject(tparts[0])
		if resp == nil || !resp.Ok {
			return nil, easyerr.Error("could not load prev data ref")
		}
		target = fmt.Sprintf("%s/%s", resp.Slug, tparts[1])

		resp = opts.Handle.GetPrevObject(thook.PlugId)
		if resp == nil || !resp.Ok {
			return nil, easyerr.Error("could not load prev plug ref")
		}
		plugId = resp.Slug
	}

	id, err := th.corehub.AddTargetHook(&entities.TargetHook{
		Id:         0,
		Name:       thook.Name,
		Policy:     thook.Policy,
		TargetType: thook.TargetType,
		Target:     target,
		EventType:  thook.EventType,
		PlugId:     plugId,
		AgentId:    thook.AgentId,
		ExecMeta:   thook.ExecMeta,
		ExtraMeta:  thook.ExtraMeta,
		TenantId:   opts.TenantId,
		Handler:    thook.Handler,
	})
	if err != nil {
		return nil, err
	}

	return &xinstance.Response{
		Ok:      false,
		Message: "",
		Slug:    fmt.Sprintf("%d", id), // fixme => return newly created id
		Data:    nil,
	}, nil

}
