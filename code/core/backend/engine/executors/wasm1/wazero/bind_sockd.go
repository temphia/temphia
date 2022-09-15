package wazero

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/libx/xutils/kosher"
)

func SendDirect(ctx context.Context, cid, roomPtr, roomLen, dataPtr, dataLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)
	e.getBytes((respOffset), (respLen))

	contents, ok := e.instance.Memory().Read(e.context, uint32(dataPtr), uint32(dataLen))
	if !ok {
		panic(ErrOutofIndex)
	}

	err := e.bindSockd.SendDirect(
		e.getString((roomPtr), (roomLen)),
		int64(cid),
		contents,
	)

	if err != nil {
		e.writeBytesNPtr(kosher.Byte(err.Error()), (respOffset), (respLen))
		return 0
	}

	return 1
}

func SendDirectBatch(ctx context.Context, roomPtr, roomLen, connIdsPtr, connIdsLen, payloadPtr, payloadLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	room := e.getString(respOffset, respLen)
	data := e.getBytes(payloadPtr, payloadLen)
	conns := make([]int64, connIdsLen)

	for idx := range conns {
		cid, ok := e.mem.ReadUint64Le(e.context, (uint32(connIdsPtr) + uint32(8*idx)))
		if !ok {
			panic(ErrOutofIndex)
		}

		conns[idx] = int64(cid)
	}

	err := e.bindSockd.SendDirectBatch(room, conns, data)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	return 1
}

func SockdSendBroadcast(ctx context.Context, roomPtr, roomLen, igPtr, igLen, payloadPtr, payloadLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	room := e.getString(respOffset, respLen)
	data := e.getBytes(payloadPtr, payloadLen)

	igconns := make([]int64, igLen)

	for idx := range igconns {
		cid, ok := e.mem.ReadUint64Le(e.context, (uint32(igPtr) + uint32(8*idx)))
		if !ok {
			panic(ErrOutofIndex)
		}

		igconns[idx] = int64(cid)
	}

	e.bindSockd.SendBroadcast(room, igconns, data)

	return 1
}

func SockdSendTagged(ctx context.Context, roomPtr, roomLen, tagsPtr, tagsLen, payloadPtr, payloadLen, respOffset, respLen int32) int32 {
	return 1
}

func SockdRoomUpdateTags(connId int64, roomPtr, roomLen, optsPtr, optsLen, respOffset, respLen int32) int32 {

	return 1

}
