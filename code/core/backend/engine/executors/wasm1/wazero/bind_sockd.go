package wazero

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/libx/xutils/kosher"
)

func SendDirect(ctx context.Context, cid, roomPtr, roomLen, dataPtr, dataLen, respPtr, respLen int32) int32 {

	e := getCtx(ctx)

	contents, ok := e.instance.Memory().Read(e.context, uint32(dataPtr), uint32(dataLen))
	if !ok {
		panic(ErrOutofIndex)
	}

	err := e.bindSockd.SendDirect(
		e.getString(uint32(roomPtr), uint32(roomLen)),
		int64(cid),
		contents,
	)

	if err != nil {
		e.write2(kosher.Byte(err.Error()), uint32(respPtr), uint32(respLen))
		return 0
	}

	return 1
}

func SendDirectBatch(ctx context.Context, cid, roomPtr, roomLen, dataPtr, dataLen, respPtr, respLen int32) int32 {

	return 1

}
