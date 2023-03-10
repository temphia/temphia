package wasmer2

import (
	"bytes"
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx/ticket"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func (w *wasmer2) buildCabinet() {

	w.bind(BindOptions{
		name:    "_cabinet_add_file",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.cabinetAddFile,
		returns: []wasmer.ValueKind{wasmer.I32, wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_cabinet_list_folder",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32},
		fn:      w.cabinetListFolder,
		returns: []wasmer.ValueKind{wasmer.I32, wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_cabinet_get_file",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.cabinetGetFile,
		returns: []wasmer.ValueKind{wasmer.I32, wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_cabinet_del_file",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.cabinetDeleteFile,
		returns: []wasmer.ValueKind{wasmer.I32, wasmer.I32},
	})

	//w.bind("_cabinet_cache_file", wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32)(w.cabinetGetFile, wasmer.I32, wasmer.I32)

}
func (w *wasmer2) cabinetAddFile(args []wasmer.Value) ([]wasmer.Value, error) {

	cab := w.binder.CabinetBindingsGet()

	folderPtr := args[0].I32()
	folderLen := args[1].I32()
	filePtr := args[2].I32()
	fileLen := args[3].I32()
	contentPtr := args[4].I32()
	contentLen := args[5].I32()

	mem := w.getMemory()

	err := cab.AddFile(
		getStr(mem, folderPtr, folderLen),
		getStr(mem, filePtr, fileLen),
		getByte(mem, contentPtr, contentLen),
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) cabinetListFolder(args []wasmer.Value) ([]wasmer.Value, error) {
	cab := w.binder.CabinetBindingsGet()

	folderPtr := args[0].I32()
	folderLen := args[1].I32()

	mem := w.getMemory()

	resp, err := cab.ListFolder(
		getStr(mem, folderPtr, folderLen),
	)

	if err != nil {
		return w.wasmErr(err), nil
	}

	var buf bytes.Buffer

	rlen := len(resp)
	for i, value := range resp {
		buf.WriteString((value))
		if (i + 1) != rlen {
			buf.WriteByte(',')
		}
	}

	return w.wasmResp(buf.Bytes(), err), nil
}

func (w *wasmer2) cabinetGetFile(args []wasmer.Value) ([]wasmer.Value, error) {

	cab := w.binder.CabinetBindingsGet()

	folderPtr := args[0].I32()
	folderLen := args[1].I32()

	filePtr := args[2].I32()
	fileLen := args[3].I32()

	mem := w.getMemory()

	out, err := cab.GetFile(
		getStr(mem, folderPtr, folderLen),
		getStr(mem, filePtr, fileLen),
	)

	return w.wasmResp(out, err), nil
}

func (w *wasmer2) cabinetDeleteFile(args []wasmer.Value) ([]wasmer.Value, error) {

	cab := w.binder.CabinetBindingsGet()

	folderPtr := args[0].I32()
	folderLen := args[1].I32()

	filePtr := args[2].I32()
	fileLen := args[3].I32()

	mem := w.getMemory()

	err := cab.DeleteFile(
		getStr(mem, folderPtr, folderLen),
		getStr(mem, filePtr, fileLen),
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) cabinetGenerateTicket(args []wasmer.Value) ([]wasmer.Value, error) {

	cab := w.binder.CabinetBindingsGet()

	folderPtr := args[0].I32()
	folderLen := args[1].I32()

	ticketPtr := args[2].I32()
	ticketLen := args[3].I32()

	mem := w.getMemory()

	raw := getByte(mem, ticketPtr, ticketLen)

	opts := &ticket.CabinetFolder{}

	err := json.Unmarshal(raw, opts)
	if err != nil {
		return w.wasmErr(err), nil
	}

	out, err := cab.Ticket(
		getStr(mem, folderPtr, folderLen),
		opts,
	)

	return w.wasmResp(kosher.Byte(out), err), nil
}
