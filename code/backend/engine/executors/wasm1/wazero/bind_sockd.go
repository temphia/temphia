package wazero

import (
	"context"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
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

	return e.writeFinal(respOffset, respLen, err)
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
	return e.writeFinal(respOffset, respLen, err)
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

	err := e.bindSockd.SendBroadcast(room, igconns, data)
	return e.writeFinal(respOffset, respLen, err)
}

func SockdSendTagged(ctx context.Context, roomPtr, roomLen, tagsPtr, tagsLen, igPtr, igLen, payloadPtr, payloadLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	room := e.getString(respOffset, respLen)
	data := e.getBytes(payloadPtr, payloadLen)
	tags := e.getString(respOffset, respLen)

	igconns := make([]int64, igLen)

	for idx := range igconns {
		cid, ok := e.mem.ReadUint64Le(e.context, (uint32(igPtr) + uint32(8*idx)))
		if !ok {
			panic(ErrOutofIndex)
		}

		igconns[idx] = int64(cid)
	}

	err := e.bindSockd.SendTagged(room, strings.Split(tags, ","), igconns, data)

	return e.writeFinal(respOffset, respLen, err)
}

func SockdRoomUpdateTags(ctx context.Context, connId int64, roomPtr, roomLen, optsPtr, optsLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	room := e.getString(respOffset, respLen)

	opts := sockdx.UpdateTagOptions{}

	err := e.getJSON(optsPtr, optsLen, &opts)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}
	return e.writeFinal(respOffset, respLen, e.bindSockd.RoomUpdateTags(room, opts))
}
