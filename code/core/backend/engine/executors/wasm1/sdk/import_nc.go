package tasmsdk

import "unsafe"

func NCPut(key string, value []byte, expire int64) error {
	var respOffset, respLen int32

	kPtr, kLen := stringToPtr(key)
	vPtr, vLen := bytesToPtr(value)

	if _nc_put(kPtr, kLen, vPtr, vLen, intAddr(&respOffset), intAddr(&respLen), expire) {
		return nil
	}

	return getErr(respOffset)

}

func NCPutCAS(key string, value []byte, version, expire int64) error {
	var respOffset, respLen int32

	kPtr, kLen := stringToPtr(key)
	vPtr, vLen := bytesToPtr(value)

	if _nc_put_cas(kPtr, kLen, vPtr, vLen, intAddr(&respOffset), intAddr(&respLen), version, expire) {
		return nil
	}

	return getErr(respOffset)
}

func NCGet(key string) (data []byte, version int64, expire int64, err error) {
	var respOffset, respLen int32
	kPtr, kLen := stringToPtr(key)

	if _nc_get(kPtr, kLen, intAddr(&respOffset), intAddr(&respLen), int32(uintptr(unsafe.Pointer(&expire)))) {
		data = getBytes(respOffset)
		return
	}

	err = getErr(respOffset)
	return
}

func NCExpire(key string) error {
	var respOffset, respLen int32
	kPtr, kLen := stringToPtr(key)

	if _nc_expire(kPtr, kLen, intAddr(&respOffset), intAddr(&respLen)) {
		return nil
	}

	return getErr(respOffset)
}

//go:wasm-module temphia1
//export nc_put
func _nc_put(kPtr, kLen, vPtr, vLen, respOffset, respLen int32, expire int64) bool

//go:wasm-module temphia1
//export nc_put_cas
func _nc_put_cas(kPtr, kLen, vPtr, vLen, respOffset, respLen int32, version, expire int64) bool

//go:wasm-module temphia1
//export nc_get
func _nc_get(kPtr, kLen, respOffset, respLen, rexpire int32) bool

//go:wasm-module temphia1
//export nc_expire
func _nc_expire(kPtr, kLen, respOffset, respLen int32) bool
