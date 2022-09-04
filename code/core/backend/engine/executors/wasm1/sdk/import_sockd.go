package tasmsdk

func SendDirect(room string, connId []string, payload []byte) error {

	return nil
}

func SendBroadcast(room string, payload []byte) error {
	return nil
}

func SendTagged(room string, tags []string, payload []byte) error {
	return nil
}

func AddToRoom(room string, connId string, tags []string) error {
	return nil
}

func KickFromRoom(room string, connId string) error {
	return nil
}

func ListRoomConns(room string) (map[string][]string, error) {
	return nil, nil
}

func BannConn(connId string) error {
	return nil
}

// private

//go:wasm-module temphia
//export sockd_send_direct
func _sockd_send_direct(roomPtr, roomLen, connIdsPtr, connIdsLen, payloadPtr, payloadLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export sockd_send_broadcast
func _sockd_send_broadcast(roomPtr, roomLen, payloadPtr, payloadLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export sockd_send_tagged
func _sockd_send_tagged(roomPtr, roomLen, tagsPtr, tagsLen, payloadPtr, payloadLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export sockd_add_to_room
func _sockd_add_to_room(roomPtr, roomLen, cidPtr, cidLen, tagsPtr, tagsLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export sockd_kick_from_room
func _sockd_kick_from_room(roomPtr, roomLen, cidPtr, cidLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export sockd_kick_from_room
func _sockd_list_room_conns(roomPtr, roomLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export sockd_kick_from_room
func _sockd_ban_conn(cidPtr, cidLen, respPtr, respLen int32) bool
