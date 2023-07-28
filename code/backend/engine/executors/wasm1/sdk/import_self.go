package tasmsdk

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

func SelfListResources() ([]*bindx.Resource, error) {
	var respOffset, respLen int32

	if _self_list_resources(intAddr(&respOffset), intAddr(&respLen)) {
		resp := make([]*bindx.Resource, 0)
		err := getJSON(respOffset, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respOffset)

}

func SelfGetResource(name string) (*bindx.Resource, error) {
	nPtr, nLen := stringToPtr(name)

	var respOffset, respLen int32

	if _self_get_resource(nPtr, nLen, intAddr(&respOffset), intAddr(&respLen)) {
		resp := &bindx.Resource{}
		err := getJSON(respOffset, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respOffset)

}
func SelfInLinks() ([]bindx.Link, error) {
	var respOffset, respLen int32

	if _self_in_links(intAddr(&respOffset), intAddr(&respLen)) {
		resp := make([]bindx.Link, 0)
		err := getJSON(respOffset, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respOffset)

}
func SelfOutLinks() ([]bindx.Link, error) {
	var respOffset, respLen int32

	if _self_out_links(intAddr(&respOffset), intAddr(&respLen)) {
		resp := make([]bindx.Link, 0)
		err := getJSON(respOffset, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respOffset)
}

func SelfLinkExec(name, method string, data []byte) ([]byte, error) {
	var respOffset, respLen int32

	nptr, nlen := stringToPtr(name)
	mptr, mlen := stringToPtr(method)
	dptr, dlen := bytesToPtr(data)

	if _self_link_exec(nptr, nlen, mptr, mlen, dptr, dlen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)
}

func SelfNewModule(name string, data []byte) (int32, error) {

	var respOffset, respLen int32

	nptr, nlen := stringToPtr(name)

	dptr, dlen := bytesToPtr(data)

	mid := _self_new_module(nptr, nlen, dptr, dlen, intAddr(&respOffset), intAddr(&respLen))

	if mid == 0 {
		return 0, getErr(respOffset)
	}

	return mid, nil
}

func SelfModuleExec(mid int32, method string, data []byte) ([]byte, error) {

	var respOffset, respLen int32

	mptr, mlen := stringToPtr(method)

	dptr, dlen := bytesToPtr(data)

	if _self_module_exec(mid, mptr, mlen, dptr, dlen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)
}

func SelfForkExec(method string, data []byte) error {
	var respOffset, respLen int32
	mptr, mlen := stringToPtr(method)
	dptr, dlen := bytesToPtr(data)

	if _self_fork_exec(mptr, mlen, dptr, dlen, intAddr(&respOffset), intAddr(&respLen)) {
		return nil
	}

	return getErr(respOffset)
}

//go:wasm-module temphia1
//export self_list_resources
func _self_list_resources(respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_get_resource
func _self_get_resource(nPtr, nLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_in_links
func _self_in_links(respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_out_links
func _self_out_links(respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_link_exec
func _self_link_exec(nPtr, nLen, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_new_module
func _self_new_module(nPtr, nLen, dPtr, dLen, respOffset, respLen int32) int32

//go:wasm-module temphia1
//export self_module_exec
func _self_module_exec(mid, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_module_ticket
func _self_module_ticket(mid, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_fork_exec
func _self_fork_exec(mPtr, mLen, dPtr, dLen, respOffset, respLen int32) bool
