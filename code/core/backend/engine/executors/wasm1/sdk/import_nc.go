package tasmsdk

import "unsafe"

func NCPut(key string, value []byte, expire int64) error {
	var respPtr, respLen int32

	kPtr, kLen := stringToPtr(key)
	vPtr, vLen := bytesToPtr(value)

	if _nc_put(kPtr, kLen, vPtr, vLen, intAddr(&respPtr), intAddr(&respLen), expire) {
		return nil
	}

	return getErr(respPtr)

}

func NCPutCAS(key string, value []byte, version, expire int64) error {
	var respPtr, respLen int32

	kPtr, kLen := stringToPtr(key)
	vPtr, vLen := bytesToPtr(value)

	if _nc_put_cas(kPtr, kLen, vPtr, vLen, intAddr(&respPtr), intAddr(&respLen), version, expire) {
		return nil
	}

	return getErr(respPtr)
}

func NCGet(key string) (data []byte, version int64, expire int64, err error) {
	var respPtr, respLen int32
	kPtr, kLen := stringToPtr(key)

	if _nc_get(kPtr, kLen, intAddr(&respPtr), intAddr(&respLen), int32(uintptr(unsafe.Pointer(&expire)))) {
		data = getBytes(respPtr)
		return
	}

	err = getErr(respPtr)
	return
}

func NCExpire(key string) error {
	var respPtr, respLen int32
	kPtr, kLen := stringToPtr(key)

	if _nc_expire(kPtr, kLen, intAddr(&respPtr), intAddr(&respLen)) {
		return nil
	}

	return getErr(respPtr)
}

//go:wasm-module temphia1
//export nc_put
func _nc_put(kPtr, kLen, vPtr, vLen, respPtr, respLen int32, expire int64) bool

//go:wasm-module temphia1
//export nc_put_cas
func _nc_put_cas(kPtr, kLen, vPtr, vLen, respPtr, respLen int32, version, expire int64) bool

//go:wasm-module temphia1
//export nc_get
func _nc_get(kPtr, kLen, respPtr, respLen, rexpire int32) bool

//go:wasm-module temphia1
//export nc_expire
func _nc_expire(kPtr, kLen, respPtr, respLen int32) bool
