package tasmsdk

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
)

func SelfAddFile(file string, data []byte) error {
	var respOffset, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_add_file(fptr, flen, dPtr, dLen, intAddr(&respOffset), intAddr(&respLen)) {
		return nil
	}

	return getErr(respOffset)
}

func SelfUpdateFile(file string, data []byte) error {

	var respOffset, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_update_file(fptr, flen, dPtr, dLen, intAddr(&respOffset), intAddr(&respLen)) {
		return nil
	}

	return getErr(respOffset)
}

func SelfAddDataFile(file string, data []byte) error {
	var respOffset, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_add_data_file(fptr, flen, dPtr, dLen, intAddr(&respOffset), intAddr(&respLen)) {
		return nil
	}

	return getErr(respOffset)
}

func SelfUpdateDataFile(file string, data []byte) error {
	var respOffset, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_update_data_file(fptr, flen, dPtr, dLen, intAddr(&respOffset), intAddr(&respLen)) {
		return nil
	}

	return getErr(respOffset)
}

func SelfGetDataFile(file string) ([]byte, error) {
	var respOffset, respLen int32

	fptr, flen := stringToPtr(file)

	if _self_get_data_file(fptr, flen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)
}

func SelfListDataFile() (map[string]string, error) {
	var respOffset, respLen int32

	if _self_list_data_file(intAddr(&respOffset), intAddr(&respLen)) {
		resp := make(map[string]string)
		err := getJSON(respOffset, &resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, getErr(respOffset)
}

func SelfDeleteDataFile(file string) error {
	var respOffset, respLen int32
	fptr, flen := stringToPtr(file)

	if _self_del_data_file(fptr, flen, intAddr(&respOffset), intAddr(&respLen)) {
		return nil
	}

	return getErr(respOffset)
}

// private

//go:wasm-module temphia1
//export self_add_file
func _self_add_file(fPtr, fLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_update_file
func _self_update_file(fPtr, fLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_add_data_file
func _self_add_data_file(fPtr, fLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_update_data_file
func _self_update_data_file(fPtr, fLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_get_data_file
func _self_get_data_file(fptr, flen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_list_data_file
func _self_list_data_file(respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_del_data_file
func _self_del_data_file(fptr, flen, respOffset, respLen int32) bool

// other

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

func SelfLinkExec(name, method string, data []byte, async, detached bool) ([]byte, error) {
	var respOffset, respLen int32

	nptr, nlen := stringToPtr(name)
	mptr, mlen := stringToPtr(method)
	dptr, dlen := bytesToPtr(data)

	var asyncI, detachedI int32
	if async {
		asyncI = 1
	}
	if detached {
		detachedI = 1
	}

	if _self_link_exec(nptr, nlen, mptr, mlen, dptr, dlen, asyncI, detachedI, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)
}

func SelfModuleExec(name, method, path string, data []byte) ([]byte, error) {

	var respOffset, respLen int32

	nptr, nlen := stringToPtr(name)
	mptr, mlen := stringToPtr(method)
	dptr, dlen := bytesToPtr(data)

	if _self_module_exec(nptr, nlen, mptr, mlen, dptr, dlen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)
}

func SelfForkExec(method string, data []byte) error {
	var respOffset, respLen int32
	mptr, mlen := stringToPtr(method)

	if _self_fork_exec(mptr, mlen, intAddr(&respOffset), intAddr(&respLen)) {
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
func _self_link_exec(nPtr, nLen, mPtr, mLen, dPtr, dLen, async, detached, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_module_exec
func _self_module_exec(nPtr, nLen, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export self_fork_exec
func _self_fork_exec(mPtr, mLen, respOffset, respLen int32) bool
