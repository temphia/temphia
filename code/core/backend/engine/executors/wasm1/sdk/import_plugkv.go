package tasmsdk

import (
	"encoding/json"
	"errors"
)

func PlugKvSet(txid int32, key, value string, opts map[string]any) error {

	keyPtr, keyLen := stringToPtr(key)
	valuePtr, valueLen := stringToPtr(value)
	optPtr, optLen := JsonPtr(opts)

	var respPtr, respLen int32

	ok := _plugkv_set(txid, keyPtr, keyLen, valuePtr, valueLen, int32(uintptr(optPtr)), optLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	resp := getBytes(respPtr)
	return errors.New(string(resp))
}

func PlugKvUpdate(txid int32, key, value string, opts map[string]any) error {

	keyPtr, keyLen := stringToPtr(key)
	valuePtr, valueLen := stringToPtr(value)
	optPtr, optLen := JsonPtr(opts)

	var respPtr, respLen int32

	ok := _plugkv_update(txid, keyPtr, keyLen, valuePtr, valueLen, int32(uintptr(optPtr)), optLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	resp := getBytes(respPtr)
	return errors.New(string(resp))
}

func PlugKvGet(txid int32, key string) (map[string]any, error) {
	keyPtr, keyLen := stringToPtr(key)
	var respPtr, respLen int32

	ok := _plugkv_get(txid, keyPtr, keyLen, intAddr(&respPtr), intAddr(&respLen))
	resp := getBytes(respPtr)
	if !ok {
		return nil, errors.New(string(resp))
	}

	data := make(map[string]any)
	err := json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PlugKvDelete(txid int32, key string) error {
	keyPtr, keyLen := stringToPtr(key)
	var respPtr, respLen int32

	ok := _plugkv_delete(txid, keyPtr, keyLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))
}

func PlugKvDeleteBatch(txid int32, keys []string) error {

	keysJson, err := json.Marshal(&keys)
	if err != nil {
		return err
	}

	keyPtr, keyLen := bytesToPtr(keysJson)
	var respPtr, respLen int32

	ok := _plugkv_delete_batch(txid, keyPtr, keyLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))

}

func PlugKvQuery(txid int32, opt map[string]any) ([]map[string]any, error) {

	optPtr, optLen := JsonPtr(opt)
	var respPtr, respLen int32

	ok := _plugkv_query(txid, int32(uintptr(optPtr)), optLen, intAddr(&respPtr), intAddr(&respLen))

	resp := getBytes(respPtr)

	if !ok {
		return nil, errors.New(string(resp))
	}

	data := make([]map[string]any, 0)

	err := json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PlugKvTxnNew() (int32, error) {
	var txid, respPtr, respLen int32

	ok := _plugkv_txn_new(intAddr(&txid), intAddr(&respPtr), intAddr(&respLen))
	if !ok {
		resp := getBytes(respPtr)
		return 0, errors.New(string(resp))
	}

	return txid, nil
}

func PlugKvTxnRollback(txid int32) error {
	var respPtr, respLen int32

	ok := _plugkv_txn_rollback(txid, intAddr(&respPtr), intAddr(&respLen))
	if !ok {
		resp := getBytes(respPtr)
		return errors.New(string(resp))
	}

	return nil
}

func PlugKvTxnCommit(txid int32) error {
	var respPtr, respLen int32

	ok := _plugkv_txn_commit(txid, intAddr(&respPtr), intAddr(&respLen))
	if !ok {
		resp := getBytes(respPtr)
		return errors.New(string(resp))
	}

	return nil
}

// private

//go:wasm-module temphia
//export plugkv_set
func _plugkv_set(txid, keyPtr, keyLen, valuePtr, valueLen, optsPtr, optsLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_update
func _plugkv_update(txid, keyPtr, keyLen, valuePtr, valueLen, optsPtr, optsLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_get
func _plugkv_get(txid, keyPtr, keyLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_del
func _plugkv_delete(txid, keyPtr, keyLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_del_batch
func _plugkv_delete_batch(txid, keyPtr, keyLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_query
func _plugkv_query(txid, optsPtr, optsLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_txn_new
func _plugkv_txn_new(txidPtr, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_txn_rollback
func _plugkv_txn_rollback(txid, respPtr, respLen int32) bool

//go:wasm-module temphia
//export plugkv_txn_commit
func _plugkv_txn_commit(txid, respPtr, respLen int32) bool
