package tasmsdk

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func ListUsers(group string) ([]string, error) {
	gptr, glen := stringToPtr(group)
	var respOffset, respLen int32

	ok := _list_user(gptr, glen, intAddr(&respOffset), intAddr(&respLen))

	if ok {
		resp := make([]string, 0)
		err := json.Unmarshal(getBytes(respOffset), &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respOffset)
}

func MessageUser(group, user string, opts *bindx.UserMessage) error {
	gptr, glen := stringToPtr(group)
	uptr, ulen := stringToPtr(user)
	var respOffset, respLen int32

	optPtr, optLen := JsonPtr(opts)

	ok := _message_user(gptr, glen, uptr, ulen, int32(uintptr(optPtr)), optLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func GetUser(group, user string) (*entities.UserInfo, error) {
	gptr, glen := stringToPtr(group)
	uptr, ulen := stringToPtr(user)
	var respOffset, respLen int32

	ok := _get_user(gptr, glen, uptr, ulen, intAddr(&respOffset), intAddr(&respLen))
	if !ok {
		return nil, getErr(respOffset)
	}

	usr := &entities.UserInfo{}
	err := json.Unmarshal(getBytes(respOffset), usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func MessageCurrentUser(opts *bindx.UserMessage) error {
	var respOffset, respLen int32

	optPtr, optLen := JsonPtr(opts)

	ok := _message_current_user(int32(uintptr(optPtr)), optLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func CurrentUser() (*entities.UserInfo, error) {
	var respOffset, respLen int32

	if ok := _current_user(intAddr(&respOffset), intAddr(&respLen)); ok {
		resp := &entities.UserInfo{}
		err := getJSON(respOffset, resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, getErr(respOffset)
}

// private

//go:wasm-module temphia1
//export list_user
func _list_user(gPtr, gLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export message_user
func _message_user(gPtr, gLen, uPtr, uLen, optPtr, optLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export message_user
func _get_user(gPtr, gLen, uPtr, uLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export message_user
func _message_current_user(optPtr, optLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export current_user
func _current_user(respOffset, respLen int32) bool
