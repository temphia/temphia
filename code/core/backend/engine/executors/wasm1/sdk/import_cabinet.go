package tasmsdk

import (
	"encoding/json"
	"errors"
)

func CabinetAddFile(folder, file string, data []byte) error {

	folderPtr, folderLen := stringToPtr(folder)
	filePtr, fileLen := stringToPtr(file)

	conPtr, conLen := bytesToPtr(data)

	var respPtr, respLen int32

	ok := _cabinet_add_file(folderPtr, folderLen, filePtr, fileLen, conPtr, conLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return errors.New(string(getBytes(respPtr)))
}

func CabinetListFolder(folder string) ([]string, error) {
	folderPtr, folderLen := stringToPtr(folder)
	var respPtr, respLen int32

	ok := _cabinet_list_folder(folderPtr, folderLen, intAddr(&respPtr), intAddr(&respLen))

	resp := getBytes(respPtr)
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
	var respPtr, respLen int32

	ok := _cabinet_get_file(folderPtr, folderLen, filePtr, fileLen, intAddr(&respPtr), intAddr(&respLen))

	resp := getBytes(respPtr)
	if ok {
		return resp, nil
	}

	return nil, errors.New(string(resp))
}

func CabinetDelFile(folder, file string) error {
	folderPtr, folderLen := stringToPtr(folder)
	filePtr, fileLen := stringToPtr(file)
	var respPtr, respLen int32

	ok := _cabinet_del_file(folderPtr, folderLen, filePtr, fileLen, intAddr(&respPtr), intAddr(&respLen))

	resp := getBytes(respPtr)
	if ok {
		return nil
	}

	return errors.New(string(resp))
}

// private

//go:wasm-module temphia
//export cabinet_add_file
func _cabinet_add_file(fPtr, fLen, filePtr, fileLen, conPtr, conLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export cabinet_list_folder
func _cabinet_list_folder(fPtr, fLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export cabinet_get_file
func _cabinet_get_file(fPtr, fLen, filePtr, fileLen, respPtr, respLen int32) bool

//go:wasm-module temphia
//export cabinet_del_file
func _cabinet_del_file(fPtr, fLen, filePtr, fileLen, respPtr, respLen int32) bool
