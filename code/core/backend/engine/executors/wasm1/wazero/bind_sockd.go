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
		e.writeWithOffsetPtr(kosher.Byte(err.Error()), uint32(respPtr), uint32(respLen))
		return 0
	}

	return 1
}

// SendDirect(room string, connId int64, payload []byte) error
// SendDirectBatch(room string, conns []int64, payload []byte) error
// SendBroadcast(room string, ignores []int64, payload []byte) error
// SendTagged(room string, tags []string, ignores []int64, payload []byte) error
// RoomUpdateTags(room string, opts sockdx.UpdateTagOptions) error

func SendDirectBatch(ctx context.Context, cid, roomPtr, roomLen, dataPtr, dataLen, respPtr, respLen int32) int32 {

	return 1

}
