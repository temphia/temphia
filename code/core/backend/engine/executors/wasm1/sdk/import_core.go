package tasmsdk

import (
	"errors"
	"unsafe"
)

func Log(msg string)                                 { log(msg) }
func Sleep(msec uint32)                              { sleep(msec) }
func GetSelfFile(file string) ([]byte, int64, error) { return getSelfFile(file) }

// private

//go:wasm-module temphia1
//export log
func _log(ptr, size int32)

func log(msg string) {
	_log(stringToPtr(msg))
}

//go:wasm-module temphia1
//export sleep
func _sleep(msec uint32)

func sleep(msec uint32) {
	_sleep(msec)
}

//go:wasm-module temphia1
//export get_self_file
func _get_self_file(file_ptr, file_size, resp_ptr, resp_len, modPtr int32) bool

func getSelfFile(file string) ([]byte, int64, error) {
	var respOffset, respLen int32 // we are not respLen but other impl might
	var mod int64
	fptr, flen := stringToPtr(file)

	ok := _get_self_file(fptr, flen, intAddr(&respOffset), intAddr(&respLen), int32(uintptr(unsafe.Pointer(&mod))))

	resp := getBytes(respOffset)

	if !ok {
		return nil, 0, errors.New(string(resp))
	}

	return resp, mod, nil
}
