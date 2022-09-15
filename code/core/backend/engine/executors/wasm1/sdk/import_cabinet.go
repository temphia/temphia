package tasmsdk

import (
	"encoding/json"
	"errors"
)

func CabinetAddFile(folder, file string, data []byte) error {

	folderPtr, folderLen := stringToPtr(folder)
	filePtr, fileLen := stringToPtr(file)

	conPtr, conLen := bytesToPtr(data)

	var respOffset, respLen int32

	ok := _cabinet_add_file(folderPtr, folderLen, filePtr, fileLen, conPtr, conLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func CabinetListFolder(folder string) ([]string, error) {
	folderPtr, folderLen := stringToPtr(folder)
	var respOffset, respLen int32

	ok := _cabinet_list_folder(folderPtr, folderLen, intAddr(&respOffset), intAddr(&respLen))

	resp := getBytes(respOffset)
	if !ok {
		return nil, errors.New(string(resp))
	}
	folders := make([]string, 0)

	err := json.Unmarshal(resp, &folder)
	if err != nil {
		return nil, err
	}

	return folders, nil
}

func CabinetGetFile(folder, file string) ([]byte, error) {
	folderPtr, folderLen := stringToPtr(folder)
	filePtr, fileLen := stringToPtr(file)
	var respOffset, respLen int32

	ok := _cabinet_get_file(folderPtr, folderLen, filePtr, fileLen, intAddr(&respOffset), intAddr(&respLen))
	resp := getBytes(respOffset)
	if ok {
		return resp, nil
	}

	return nil, errors.New(string(resp))
}

func CabinetDelFile(folder, file string) error {
	folderPtr, folderLen := stringToPtr(folder)
	filePtr, fileLen := stringToPtr(file)
	var respOffset, respLen int32

	ok := _cabinet_del_file(folderPtr, folderLen, filePtr, fileLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func CabinetGenerateTkt(folder string, opts map[string]any) (string, error) {
	folderPtr, folderLen := stringToPtr(folder)
	optPtr, optLen := JsonPtr(opts)

	var respOffset, respLen int32

	ok := _cabinet_generate_tkt(folderPtr, folderLen, int32(uintptr(optPtr)), optLen, intAddr(&respOffset), intAddr(&respLen))
	resp := getBytes(respOffset)
	if ok {
		return string(resp), nil
	}

	return "", errors.New(string(resp))
}

// private

//go:wasm-module temphia1
//export cabinet_add_file
func _cabinet_add_file(fPtr, fLen, filePtr, fileLen, conPtr, conLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export cabinet_list_folder
func _cabinet_list_folder(fPtr, fLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export cabinet_get_file
func _cabinet_get_file(fPtr, fLen, filePtr, fileLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export cabinet_del_file
func _cabinet_del_file(fPtr, fLen, filePtr, fileLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export cabinet_generate_tkt
func _cabinet_generate_tkt(fPtr, fLen, optPtr, optLen, respOffset, respLen int32) bool
