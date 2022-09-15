package tasmsdk

import (
	"bytes"
	"unsafe"
)

func SendDirect(room string, connId int64, payload []byte) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)

	var respOffset, respLen int32

	ok := _sockd_send_direct(connId, roomPtr, roomLen, payloadPtr, payloadLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func SendDirectBatch(room string, connIds []int64, payload []byte) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)
	var respOffset, respLen int32

	connIdsPtr := int32(uintptr(unsafe.Pointer(&connIds[0])))
	connIdsLen := int32(len(connIds))

	ok := _sockd_send_direct_batch(roomPtr, roomLen, connIdsPtr, connIdsLen, payloadPtr, payloadLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)

}

func SendBroadcast(room string, payload []byte, ignores ...int64) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)

	var respOffset, respLen int32

	igPtr := int32(uintptr(unsafe.Pointer(&ignores[0])))
	igLen := int32(len(ignores))

	ok := _sockd_send_broadcast(roomPtr, roomLen, igPtr, igLen, payloadPtr, payloadLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)

}

func SendTagged(room string, tags []string, payload []byte) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)

	var respOffset, respLen int32

	var tagBuf bytes.Buffer
	for i, v := range tags {
		tagBuf.WriteString(v)
		if len(tags) != i+1 {
			tagBuf.WriteByte(',')
		}
	}

	tb := tagBuf.Bytes()
	tPtr, tLen := bytesToPtr(tb)

	ok := _sockd_send_tagged(
		roomPtr,
		roomLen,
		tPtr,
		tLen,
		payloadPtr,
		payloadLen,
		intAddr(&respOffset),
		intAddr(&respLen),
	)
	if ok {
		return nil
	}

	return getErr(respOffset)

}

func RoomUpdateTags(connId int64, room string, opts map[string]any) error {
	roomPtr, roomLen := stringToPtr(room)
	optsPtr, optsLen := JsonPtr(opts)
	var respOffset, respLen int32

	ok := _sockd_room_update_tags(connId, roomPtr, roomLen, int32(uintptr(optsPtr)), optsLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

// private

//go:wasm-module temphia1
//export sockd_send_direct
func _sockd_send_direct(connId int64, roomPtr, roomLen, payloadPtr, payloadLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export sockd_send_direct_batch
func _sockd_send_direct_batch(roomPtr, roomLen, connIdsPtr, connIdsLen, payloadPtr, payloadLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export sockd_send_broadcast
func _sockd_send_broadcast(roomPtr, roomLen, igPtr, igLen, payloadPtr, payloadLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export sockd_send_tagged
func _sockd_send_tagged(roomPtr, roomLen, tagsPtr, tagsLen, payloadPtr, payloadLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export sockd_room_update_tags
func _sockd_room_update_tags(connId int64, roomPtr, roomLen, optsPtr, optsLen, respOffset, respLen int32) bool
