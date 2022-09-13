package tasmsdk

import (
	"encoding/json"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func ListUsers(group string) ([]string, error) {
	gptr, glen := stringToPtr(group)
	var respPtr, respLen int32

	ok := _list_user(gptr, glen, intAddr(&respPtr), intAddr(&respLen))

	if ok {
		resp := make([]string, 0)
		err := json.Unmarshal(getBytes(respPtr), &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respPtr)
}

func MessageUser(group, user string, opts *bindx.UserMessage) error {
	gptr, glen := stringToPtr(group)
	uptr, ulen := stringToPtr(user)
	var respPtr, respLen int32

	optPtr, optLen := JsonPtr(opts)

	ok := _message_user(gptr, glen, uptr, ulen, int32(uintptr(optPtr)), optLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respPtr)
}

func GetUser(group, user string) (*entities.UserInfo, error) {
	gptr, glen := stringToPtr(group)
	uptr, ulen := stringToPtr(user)
	var respPtr, respLen int32

	ok := _get_user(gptr, glen, uptr, ulen, intAddr(&respPtr), intAddr(&respLen))
	if !ok {
		return nil, getErr(respPtr)
	}

	usr := &entities.UserInfo{}
	err := json.Unmarshal(getBytes(respPtr), usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func MessageCurrentUser(opts *bindx.UserMessage) error {
	var respPtr, respLen int32

	optPtr, optLen := JsonPtr(opts)

	ok := _message_current_user(int32(uintptr(optPtr)), optLen, intAddr(&respPtr), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respPtr)
}

func CurrentUser() (*entities.UserInfo, error) {
	var respPtr, respLen int32

	if ok := _current_user(intAddr(&respPtr), intAddr(&respLen)); ok {
		resp := &entities.UserInfo{}
		err := getJSON(respPtr, resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, getErr(respPtr)
}

// private

//go:wasm-module temphia1
//export list_user
func _list_user(gPtr, gLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export message_user
func _message_user(gPtr, gLen, uPtr, uLen, optPtr, optLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export message_user
func _get_user(gPtr, gLen, uPtr, uLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export message_user
func _message_current_user(optPtr, optLen, respPtr, respLen int32) bool

//go:wasm-module temphia1
//export current_user
func _current_user(respPtr, respLen int32) bool
