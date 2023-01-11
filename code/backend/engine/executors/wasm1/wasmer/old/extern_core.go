package wasmer2

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func (w *wasmer2) buildCore() {

	// @link => rustsdk core

	w.bind(BindOptions{
		name:    "_get_resp",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32},
		fn:      w.getResp,
		returns: []wasmer.ValueKind{},
	})

	w.bind(BindOptions{
		name:    "_log",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32},
		fn:      w.log,
		returns: []wasmer.ValueKind{},
	})

	w.bind(BindOptions{
		name:    "_sleep",
		kinds:   []wasmer.ValueKind{wasmer.I32},
		fn:      w.sleep,
		returns: []wasmer.ValueKind{},
	})

	/*
		1st step => tells what to load (ctx_var?, payload?) => return length of items
		2nd step => give list of ptr/offset to write data_to
	*/

	w.bind(BindOptions{
		name:    "_get_event_step1",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.getEventStep1,
		returns: []wasmer.ValueKind{},
	})

	w.bind(BindOptions{
		name:    "_get_event_step2",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.getEventStep2,
		returns: []wasmer.ValueKind{},
	})

	w.bind(BindOptions{
		name:    "_set_event_reply",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.setEventReply,
		returns: []wasmer.ValueKind{},
	})

}

func (w *wasmer2) getResp(args []wasmer.Value) ([]wasmer.Value, error) {
	copy(w.getMemory()[args[0].I32():], w.pendingResp)
	return []wasmer.Value{}, nil
}

func (w *wasmer2) log(args []wasmer.Value) ([]wasmer.Value, error) {
	mem := w.getMemory()
	message := string(mem[args[0].I32() : args[0].I32()+args[1].I32()])
	w.binder.Log(message)
	return []wasmer.Value{}, nil
}

func (w *wasmer2) sleep(args []wasmer.Value) ([]wasmer.Value, error) {
	w.binder.Sleep(args[0].I32())
	return []wasmer.Value{}, nil
}

func (w *wasmer2) getEventStep1(args []wasmer.Value) ([]wasmer.Value, error) {
	loadPayload := args[0].I32()
	loadCtxvar := args[1].I32()

	rid_len := args[2].I32()
	rtype_len := args[3].I32()
	rname_len := args[4].I32()
	rctx_var_len := args[5].I32()
	rdata_len := args[6].I32()

	var id_len int32
	var type_len int32
	var name_len int32
	var ctx_var_len int32
	var data_len int32

	mem := w.getMemory()

	if loadPayload == 1 {
		// data_len = int32(len(w.eventRequest.Data))
	}

	if loadCtxvar == 1 {
		// out, err := json.Marshal(w.eventRequest.Vars)
		// if err != nil {
		// 	return nil, err
		// }
		// w.pendingResp = out
		// ctx_var_len = int32(len(out))
	}

	id_len = int32(len(w.eventRequest.Id))
	type_len = int32(len("FIXME"))
	name_len = int32(len(w.eventRequest.Name))

	writeInt(mem, rid_len, int32(id_len))
	writeInt(mem, rtype_len, int32(type_len))
	writeInt(mem, rname_len, int32(name_len))
	writeInt(mem, rctx_var_len, int32(ctx_var_len))
	writeInt(mem, rdata_len, int32(data_len))

	return []wasmer.Value{}, nil
}

func (w *wasmer2) getEventStep2(args []wasmer.Value) ([]wasmer.Value, error) {
	idptr := args[0].I32()
	typeptr := args[1].I32()
	nameptr := args[2].I32()
	ctxvarptr := args[3].I32()
	//pptr := args[4].I32()

	mem := w.getMemory()

	pp.Println(args)

	copy(mem[idptr:], kosher.Byte(w.eventRequest.Id))
	copy(mem[typeptr:], kosher.Byte("FIXME"))
	copy(mem[nameptr:], kosher.Byte(w.eventRequest.Name))

	if ctxvarptr != 0 {
		copy(mem[ctxvarptr:], w.pendingResp)
	}

	// if pptr != 0 {
	// 	copy(mem[pptr:], w.eventRequest.Data)
	// }

	return []wasmer.Value{}, nil
}

func (w *wasmer2) setEventReply(args []wasmer.Value) ([]wasmer.Value, error) {
	metaPtr := args[0].I32()
	metaLen := args[1].I32()
	// dataPtr := args[2].I32()
	// dataLen := args[3].I32()

	mem := w.getMemory()

	reply := &event.Response{
		Payload: nil,
	}

	if metaLen != 0 {
		meta := map[string]any{}
		err := json.Unmarshal(mem[metaPtr:metaPtr+metaLen], &mem)
		if err != nil {
			// fixme => store/log outside
			return nil, err
		}

		pp.Println(meta)

	}

	// if dataLen != 0 {
	// 	data := make([]byte, 0, dataLen)
	// 	data = append(data, mem[dataPtr:(dataPtr+dataLen)]...)
	// 	reply.Payload = data
	// }

	w.eventReply = reply

	return []wasmer.Value{}, nil
}
