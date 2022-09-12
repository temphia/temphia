package tasmsdk

import (
	"bytes"
	"errors"
	"unsafe"
)

func SendDirect(room string, connId int64, payload []byte) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)

	var respPtr, respLen int32

	ok := _sockd_send_direct(connId, roomPtr, roomLen, payloadPtr, payloadLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))
}

func SendDirectBatch(room string, connId []int64, payload []byte) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)
	var respPtr, respLen int32

	connIdsPtr := int32(uintptr(unsafe.Pointer(&connId[0])))
	connIdsLen := int32(len(connId))

	ok := _sockd_send_direct_batch(roomPtr, roomLen, connIdsPtr, connIdsLen, payloadPtr, payloadLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))

}

func SendBroadcast(room string, payload []byte) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)

	var respPtr, respLen int32

	ok := _sockd_send_broadcast(roomPtr, roomLen, payloadPtr, payloadLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))

}

func SendTagged(room string, tags []string, payload []byte) error {

	roomPtr, roomLen := stringToPtr(room)
	payloadPtr, payloadLen := bytesToPtr(payload)

	var respPtr, respLen int32

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
		intAddr(&respPtr),
		intAddr(&respLen),
	)
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))

}

func RoomUpdateTags(connId int64, room string, opts map[string]any) error {
	roomPtr, roomLen := stringToPtr(room)
	optsPtr, optsLen := JsonPtr(opts)
	var respPtr, respLen int32

	ok := _sockd_room_update_tags(connId, roomPtr, roomLen, int32(uintptr(optsPtr)), optsLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))
}

// private

//go:wasm-module temphia1
//export sockd_send_direct
func _sockd_send_direct(connId int64, roomPtr, roomLen, payloadPtr, payloadLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export sockd_send_direct_batch
func _sockd_send_direct_batch(roomPtr, roomLen, connIdsPtr, connIdsLen, payloadPtr, payloadLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export sockd_send_broadcast
func _sockd_send_broadcast(roomPtr, roomLen, payloadPtr, payloadLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export sockd_send_tagged
func _sockd_send_tagged(roomPtr, roomLen, tagsPtr, tagsLen, payloadPtr, payloadLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export sockd_room_update_tags
func _sockd_room_update_tags(connId int64, roomPtr, roomLen, optsPtr, optsLen, respPtr, respLen int32) bool
