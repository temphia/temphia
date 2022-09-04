package tasmsdk

import "unsafe"

var ROOT = map[uintptr][]byte{}
var emptyJsonObj = []byte(`{}`)
var emptyJsonArr = []byte(`[]`)

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
