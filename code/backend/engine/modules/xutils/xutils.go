package xutils

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
)

type Xutils struct {
	modipc *modipc.ModIPC
}

func (x *Xutils) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {
	return x.modipc.Handle(method, args)
}

func (x *Xutils) Close() error {
	return nil
}
