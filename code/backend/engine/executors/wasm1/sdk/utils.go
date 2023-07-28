package tasmsdk

import (
	"encoding/json"
	"errors"
	"unsafe"
)

var ROOT = map[uintptr][]byte{}

func allocBytes(size int32) int32 {
	data := make([]byte, size)
	pointer := (unsafe.Pointer(&data[0]))
	return int32(uintptr(pointer))
}

func freeBytes(addr int32) {
	delete(ROOT, uintptr(addr))
}

// only called internally also deletes from ROOT
func getBytes(ptr int32) []byte {
	key := uintptr(ptr)
	resp := ROOT[key]
	delete(ROOT, key)
	return resp
}

func getErr(ptr int32) error {
	return errors.New(string(getBytes(ptr)))
}

func stringToPtr(s string) (int32, int32) {
	return bytesToPtr([]byte(s))
}

func bytesToPtr(buf []byte) (int32, int32) {
	ptr := &buf[0]
	return int32(uintptr(unsafe.Pointer(ptr))), int32(len(buf))
}

func intAddr(i *int32) int32 {
	return int32(uintptr(unsafe.Pointer(i)))
}

func JsonPtr(obj any) (unsafe.Pointer, int32) {
	if obj == nil {
		return unsafe.Pointer(nil), 0
	}

	out, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return unsafe.Pointer(&out[0]), int32(len(out))
}

func getJSON(ptr int32, target any) error {
	return json.Unmarshal(getBytes(ptr), target)
}
