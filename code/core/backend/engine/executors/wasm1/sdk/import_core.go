package tasmsdk

import (
	"errors"
	"unsafe"
)

func Log(msg string)                                  { log(msg) }
func Sleep(msec uint32)                               { sleep(msec) }
func GetSelfFile(file string) ([]byte, error)         { return getSelfFile(file) }
func GetSelfFile2(file string) ([]byte, int64, error) { return getSelfFile2(file) }

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
func _get_self_file(file_ptr, file_size, resp_ptr, resp_len int32) bool

func getSelfFile(file string) ([]byte, error) {
	var respPtr, respLen int32 // we are not respLen but other impl might
	fptr, flen := stringToPtr(file)

	ok := _get_self_file(fptr, flen, intAddr(&respPtr), intAddr(&respLen))

	resp := getBytes(respPtr)

	if !ok {
		return nil, errors.New(string(resp))
	}

	return resp, nil
}

//go:wasm-module temphia1
//export get_self_file2
func _get_self_file2(file_ptr, file_size, resp_ptr, resp_len, mod int32) bool

func getSelfFile2(file string) ([]byte, int64, error) {
	var respPtr, respLen int32 // we are not respLen but other impl might
	var mod int64
	fptr, flen := stringToPtr(file)

	ok := _get_self_file2(fptr, flen, intAddr(&respPtr), intAddr(&respLen), int32(uintptr(unsafe.Pointer(&mod))))

	resp := getBytes(respPtr)

	if !ok {
		return nil, 0, errors.New(string(resp))
	}

	return resp, mod, nil
}
