package linked

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type Linked struct {
	linkedFromExecId string
	fromPlugId       string
	fromAgentId      string
	userctx          *claim.UserContext
}

func New(from, pid, aid string, userctx *claim.UserContext) invoker.Invoker {
	return &Linked{
		linkedFromExecId: from,
		fromPlugId:       pid,
		fromAgentId:      aid,
		userctx:          userctx,
	}
}

func (l *Linked) Type() string {
	return "linked"
}

func (l *Linked) ExecuteMethod(action string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return nil, easyerr.NotImpl()
}

func (l *Linked) UserContext() *claim.UserContext {
	return l.userctx
}

func (l *Linked) GetAttr(string) interface{} {
	return nil
}

func (l *Linked) GetAttrs() map[string]interface{} {
	return map[string]interface{}{}
}
