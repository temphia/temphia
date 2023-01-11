package wasmer2

/*

func (w *wasmer2) buildSockd() {

	w.bind(BindOptions{
		name:    "_sockd_send_direct",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.sockdSendDirect,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_sockd_send_broadcast",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.sockdSendBroadcast,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_sockd_send_tagged",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.sockdSendTagged,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_sockd_add_to_room",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.sockdAddToRoom,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_sockd_kick_from_room",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.sockdKickFromRoom,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_sockd_list_room_conns",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.sockdListRoomConns,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_sockd_bann_conn",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.sockdBanConn,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

}

func (w *wasmer2) sockdSendDirect(args []wasmer.Value) ([]wasmer.Value, error) {
	sd := w.binder.GetSockdBindings()

	roomPtr := args[0].I32()
	roomLen := args[1].I32()
	connsPtr := args[2].I32()
	connsLen := args[3].I32()
	dataPtr := args[4].I32()
	dataLen := args[5].I32()

	mem := w.getMemory()

	conns := strings.Split(getStr(mem, connsPtr, connsLen), ",")

	err := sd.SendDirect(
		getStr(mem, roomPtr, roomLen),
		conns,
		getByte(mem, dataPtr, dataLen),
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) sockdSendBroadcast(args []wasmer.Value) ([]wasmer.Value, error) {
	sd := w.binder.GetSockdBindings()

	roomPtr := args[0].I32()
	roomLen := args[1].I32()
	// connsPtr := args[2].I32()
	// connsLen := args[3].I32()
	dataPtr := args[4].I32()
	dataLen := args[5].I32()

	mem := w.getMemory()

	// conns := strings.Split(getStr(mem, connsPtr, connsLen), ",")

	err := sd.SendBroadcast(
		getStr(mem, roomPtr, roomLen),
		getByte(mem, dataPtr, dataLen),
	)
	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) sockdSendTagged(args []wasmer.Value) ([]wasmer.Value, error) {
	sd := w.binder.GetSockdBindings()

	roomPtr := args[0].I32()
	roomLen := args[1].I32()
	tagPtr := args[2].I32()
	tagLen := args[3].I32()
	igConnsPtr := args[4].I32()
	igConnsLen := args[5].I32()
	dataPtr := args[6].I32()
	dataLen := args[7].I32()

	mem := w.getMemory()

	tags := strings.Split(getStr(mem, tagPtr, tagLen), ",")
	igConns := strings.Split(getStr(mem, igConnsPtr, igConnsLen), ",")

	err := sd.SendTagged(
		getStr(mem, roomPtr, roomLen),
		tags,
		igConns,
		getByte(mem, dataPtr, dataLen),
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) sockdAddToRoom(args []wasmer.Value) ([]wasmer.Value, error) {

	sd := w.binder.GetSockdBindings()

	roomPtr := args[0].I32()
	roomLen := args[1].I32()
	connPtr := args[2].I32()
	connLen := args[3].I32()
	tagsPtr := args[4].I32()
	tagsLen := args[5].I32()

	mem := w.getMemory()

	tags := strings.Split(getStr(mem, tagsPtr, tagsLen), ",")

	err := sd.AddToRoom(
		getStr(mem, roomPtr, roomLen),
		getStr(mem, connPtr, connLen),
		tags,
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) sockdKickFromRoom(args []wasmer.Value) ([]wasmer.Value, error) {

	sd := w.binder.GetSockdBindings()

	roomPtr := args[0].I32()
	roomLen := args[1].I32()
	connPtr := args[2].I32()
	connLen := args[3].I32()

	mem := w.getMemory()

	err := sd.KickFromRoom(
		getStr(mem, roomPtr, roomLen),
		getStr(mem, connPtr, connLen),
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) sockdListRoomConns(args []wasmer.Value) ([]wasmer.Value, error) {

	sd := w.binder.GetSockdBindings()

	roomPtr := args[0].I32()
	roomLen := args[1].I32()

	mem := w.getMemory()

	resp, err := sd.ListRoomConns(
		getStr(mem, roomPtr, roomLen),
	)

	out, err := json.Marshal(&resp)
	return w.wasmResp(out, err), nil
}

func (w *wasmer2) sockdBanConn(args []wasmer.Value) ([]wasmer.Value, error) {

	sd := w.binder.GetSockdBindings()

	connPtr := args[0].I32()
	connLen := args[1].I32()

	mem := w.getMemory()

	err := sd.BannConn(getStr(mem, connPtr, connLen))
	return w.wasmResp(nil, err), nil
}

*/
