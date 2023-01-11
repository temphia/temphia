package wasmer2

import (
	"encoding/json"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/store"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func (w *wasmer2) buildPlugKV() {
	w.bind(BindOptions{
		name:    "_plugkv_set",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.pkvSet,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_update",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.pkvUpdate,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_get",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.pkvGet,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_del",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.pkvDel,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_del_batch",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.pkvDelBatch,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_query",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      w.pkvQuery,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_new_txn",
		kinds:   []wasmer.ValueKind{wasmer.I32},
		fn:      w.pkvNewTxn,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_rollback",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32},
		fn:      w.pkvRollback,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

	w.bind(BindOptions{
		name:    "_plugkv_commit",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32},
		fn:      w.pkvCommit,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

}

func (w *wasmer2) pkvSet(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()
	keyPtr := args[1].I32()
	keyLen := args[2].I32()
	valPtr := args[3].I32()
	valLen := args[4].I32()
	optPtr := args[5].I32()
	optLen := args[6].I32()

	mem := w.getMemory()

	var opts *store.SetOptions

	if optLen != 0 {
		opts = &store.SetOptions{}
		err := json.Unmarshal(getByte(mem, optPtr, optLen), opts)
		if err != nil {
			return w.wasmErr(err), nil
		}
	}

	err := pkv.Set(
		uint32(txid),
		getStr(mem, keyPtr, keyLen),
		getStr(mem, valPtr, valLen),
		opts,
	)

	return w.wasmResp(nil, err), nil

}

func (w *wasmer2) pkvUpdate(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()
	keyPtr := args[1].I32()
	keyLen := args[2].I32()
	valPtr := args[3].I32()
	valLen := args[4].I32()
	optPtr := args[5].I32()
	optLen := args[6].I32()

	mem := w.getMemory()

	var opts *store.UpdateOptions

	if optLen != 0 {
		opts = &store.UpdateOptions{}
		err := json.Unmarshal(getByte(mem, optPtr, optLen), opts)
		if err != nil {
			return w.wasmErr(err), nil
		}
	}

	err := pkv.Update(
		uint32(txid),
		getStr(mem, keyPtr, keyLen),
		getStr(mem, valPtr, valLen),
		opts,
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) pkvGet(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()
	keyPtr := args[1].I32()
	keyLen := args[2].I32()

	pp.Println(keyPtr, keyLen)

	mem := w.getMemory()

	key := getStr(mem, keyPtr, keyLen)

	pp.Println(key)

	out, err := pkv.Get(
		uint32(txid),
		key,
	)
	if err != nil {
		return w.wasmErr(err), nil
	}

	respBytes, err := json.Marshal(out)
	return w.wasmResp(respBytes, err), nil
}

func (w *wasmer2) pkvDel(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()
	keysPtr := args[1].I32()
	keysLen := args[2].I32()

	mem := w.getMemory()

	err := pkv.Del(
		uint32(txid),
		getStr(mem, keysPtr, keysLen),
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) pkvDelBatch(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()
	keyPtr := args[1].I32()
	keyLen := args[2].I32()

	mem := w.getMemory()

	err := pkv.DelBatch(
		uint32(txid),
		strings.Split(getStr(mem, keyPtr, keyLen), ","),
	)

	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) pkvQuery(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()
	qPtr := args[1].I32()
	qLen := args[2].I32()

	mem := w.getMemory()

	query := &store.PkvQuery{}
	err := json.Unmarshal(getByte(mem, qPtr, qLen), query)
	if err != nil {
		return w.wasmErr(err), nil
	}

	resp, err := pkv.Query(
		uint32(txid),
		query,
	)

	out, err := json.Marshal(resp)
	return w.wasmResp(out, err), nil
}

func (w *wasmer2) pkvNewTxn(args []wasmer.Value) ([]wasmer.Value, error) {
	pkv := w.binder.PlugKVBindingsGet()

	txid, err := pkv.NewTxn()
	if err != nil {
		return w.wasmErr(err), nil
	}

	return []wasmer.Value{wasmer.NewI32(wasmer.NewI32(txid))}, nil
}

func (w *wasmer2) pkvRollback(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()

	err := pkv.RollBack(
		uint32(txid),
	)
	return w.wasmResp(nil, err), nil
}

func (w *wasmer2) pkvCommit(args []wasmer.Value) ([]wasmer.Value, error) {

	pkv := w.binder.PlugKVBindingsGet()

	txid := args[0].I32()

	err := pkv.Commit(
		uint32(txid),
	)

	return w.wasmResp(nil, err), nil
}
