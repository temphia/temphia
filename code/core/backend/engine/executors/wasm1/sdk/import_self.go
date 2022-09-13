package tasmsdk

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
)

func SelfAddFile(file string, data []byte) error {
	var respPtr, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_add_file(fptr, flen, dPtr, dLen, intAddr(&respPtr), intAddr(&respLen)) {
		return nil
	}

	return getErr(respPtr)
}

func SelfUpdateFile(file string, data []byte) error {

	var respPtr, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_update_file(fptr, flen, dPtr, dLen, intAddr(&respPtr), intAddr(&respLen)) {
		return nil
	}

	return getErr(respPtr)
}

func SelfAddDataFile(file string, data []byte) error {
	var respPtr, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_add_data_file(fptr, flen, dPtr, dLen, intAddr(&respPtr), intAddr(&respLen)) {
		return nil
	}

	return getErr(respPtr)
}

func SelfUpdateDataFile(file string, data []byte) error {
	var respPtr, respLen int32
	fptr, flen := stringToPtr(file)
	dPtr, dLen := bytesToPtr(data)

	if _self_update_data_file(fptr, flen, dPtr, dLen, intAddr(&respPtr), intAddr(&respLen)) {
		return nil
	}

	return getErr(respPtr)
}

func SelfGetDataFile(file string) ([]byte, error) {
	var respPtr, respLen int32

	fptr, flen := stringToPtr(file)

	if _self_get_data_file(fptr, flen, intAddr(&respPtr), intAddr(&respLen)) {
		return getBytes(respPtr), nil
	}

	return nil, getErr(respPtr)
}

func SelfListDataFile() (map[string]string, error) {
	var respPtr, respLen int32

	if _self_list_data_file(intAddr(&respPtr), intAddr(&respLen)) {
		resp := make(map[string]string)
		err := getJSON(respPtr, &resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, getErr(respPtr)
}

func SelfDeleteDataFile(file string) error {
	var respPtr, respLen int32
	fptr, flen := stringToPtr(file)

	if _self_del_data_file(fptr, flen, intAddr(&respPtr), intAddr(&respLen)) {
		return nil
	}

	return getErr(respPtr)
}

// private

//go:wasm-module temphia1
//export self_add_file
func _self_add_file(fPtr, fLen, dPtr, dLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_update_file
func _self_update_file(fPtr, fLen, dPtr, dLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_add_data_file
func _self_add_data_file(fPtr, fLen, dPtr, dLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_update_data_file
func _self_update_data_file(fPtr, fLen, dPtr, dLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_get_data_file
func _self_get_data_file(fptr, flen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_list_data_file
func _self_list_data_file(respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_del_data_file
func _self_del_data_file(fptr, flen, respPtr, respLen int32) bool

// other

func SelfListResources() ([]*bindx.Resource, error) {
	var respPtr, respLen int32

	if _self_list_resources(intAddr(&respPtr), intAddr(&respLen)) {
		resp := make([]*bindx.Resource, 0)
		err := getJSON(respPtr, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respPtr)

}

func SelfGetResource(name string) (*bindx.Resource, error) {
	nPtr, nLen := stringToPtr(name)

	var respPtr, respLen int32

	if _self_get_resource(nPtr, nLen, intAddr(&respPtr), intAddr(&respLen)) {
		resp := &bindx.Resource{}
		err := getJSON(respPtr, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respPtr)

}
func SelfInLinks() ([]bindx.Link, error) {
	var respPtr, respLen int32

	if _self_in_links(intAddr(&respPtr), intAddr(&respLen)) {
		resp := make([]bindx.Link, 0)
		err := getJSON(respPtr, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respPtr)

}
func SelfOutLinks() ([]bindx.Link, error) {
	var respPtr, respLen int32

	if _self_out_links(intAddr(&respPtr), intAddr(&respLen)) {
		resp := make([]bindx.Link, 0)
		err := getJSON(respPtr, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respPtr)
}

func SelfLinkExec(name, method string, data []byte, async, detached bool) ([]byte, error) {
	var respPtr, respLen int32

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

	if _self_link_exec(nptr, nlen, mptr, mlen, dptr, dlen, asyncI, detachedI, intAddr(&respPtr), intAddr(&respLen)) {
		return getBytes(respPtr), nil
	}

	return nil, getErr(respPtr)
}

func SelfModuleExec(name, method, path string, data []byte) ([]byte, error) {

	var respPtr, respLen int32

	nptr, nlen := stringToPtr(name)
	mptr, mlen := stringToPtr(method)
	dptr, dlen := bytesToPtr(data)

	if _self_module_exec(nptr, nlen, mptr, mlen, dptr, dlen, intAddr(&respPtr), intAddr(&respLen)) {
		return getBytes(respPtr), nil
	}

	return nil, getErr(respPtr)
}

func SelfForkExec(method string, data []byte) error {
	var respPtr, respLen int32
	mptr, mlen := stringToPtr(method)

	if _self_fork_exec(mptr, mlen, intAddr(&respPtr), intAddr(&respLen)) {
		return nil
	}

	return getErr(respPtr)
}

//go:wasm-module temphia1
//export self_list_resources
func _self_list_resources(respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_get_resource
func _self_get_resource(nPtr, nLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_in_links
func _self_in_links(respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_out_links
func _self_out_links(respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_link_exec
func _self_link_exec(nPtr, nLen, mPtr, mLen, dPtr, dLen, async, detached, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_module_exec
func _self_module_exec(nPtr, nLen, mPtr, mLen, dPtr, dLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export self_fork_exec
func _self_fork_exec(mPtr, mLen, respPtr, respLen int32) bool
