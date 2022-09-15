package tasmsdk

import "github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"

func HttpRaw(req *bindx.HttpRequest) (*bindx.HttpResponse, error) {

	mPtr, mLen := stringToPtr(req.Method)
	pPtr, pLen := stringToPtr(req.Path)
	bPtr, bLen := bytesToPtr(req.Body)
	hPtr, hLen := JsonPtr(req.Headers)

	var rStatus int32
	var rHeadPtr, rHeadLen int32
	var risJson int32
	var rbodyPtr int32
	var rbodyLen int32

	if _http_raw(
		mPtr, mLen,
		pPtr, pLen,
		int32(uintptr(hPtr)), hLen,
		bPtr, bLen,
		intAddr(&rStatus), intAddr(&rHeadPtr), intAddr(&rHeadLen),
		intAddr(&risJson), intAddr(&rbodyPtr), intAddr(&rbodyLen),
	) {

		resp := &bindx.HttpResponse{
			SatusCode: int(rStatus),
			Headers:   map[string][]string{},
			Json:      risJson == 0,
			Body:      getBytes(rbodyLen),
		}

		err := getJSON(rHeadPtr, &resp.Headers)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(rHeadPtr)
}

func HttpRawBatch(reqs []*bindx.HttpRequest) ([]bindx.HttpResponse, error) {
	var respOffset, respLen int32
	rPtr, rLen := JsonPtr(reqs)

	if _http_raw_batch(int32(uintptr(rPtr)), rLen, intAddr(&respOffset), intAddr(&respLen)) {
		resp := make([]bindx.HttpResponse, 0)
		err := getJSON(respOffset, &resp)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, getErr(respOffset)
}

func HttpQuickGet(url string, headers map[string]string) ([]byte, error) {

	var respOffset, respLen int32

	uPtr, uLen := stringToPtr(url)
	hPtr, hLen := JsonPtr(headers)

	if _http_quick_get(uPtr, uLen, int32(uintptr(hPtr)), hLen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)
}

func HttpQuickPost(url string, headers map[string]string, data []byte) ([]byte, error) {

	var respOffset, respLen int32

	uPtr, uLen := stringToPtr(url)
	hPtr, hLen := JsonPtr(headers)
	dPtr, dLen := bytesToPtr(data)

	if _http_quick_post(uPtr, uLen, int32(uintptr(hPtr)), hLen, dPtr, dLen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)

}
func HttpFormPost(url string, headers map[string]string, data []byte) ([]byte, error) {
	var respOffset, respLen int32

	uPtr, uLen := stringToPtr(url)
	hPtr, hLen := JsonPtr(headers)
	dPtr, dLen := bytesToPtr(data)

	if _http_form_post(uPtr, uLen, int32(uintptr(hPtr)), hLen, dPtr, dLen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)

}

func HttpJsonGet(url string, headers map[string]string) ([]byte, error) {

	var respOffset, respLen int32

	uPtr, uLen := stringToPtr(url)
	hPtr, hLen := JsonPtr(headers)

	if _http_json_get(uPtr, uLen, int32(uintptr(hPtr)), hLen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)

}

func HttpJsonPost(url string, headers map[string]string, data []byte) ([]byte, error) {
	var respOffset, respLen int32

	uPtr, uLen := stringToPtr(url)
	hPtr, hLen := JsonPtr(headers)
	dPtr, dLen := bytesToPtr(data)

	if _http_json_post(uPtr, uLen, int32(uintptr(hPtr)), hLen, dPtr, dLen, intAddr(&respOffset), intAddr(&respLen)) {
		return getBytes(respOffset), nil
	}

	return nil, getErr(respOffset)
}

//go:wasm-module temphia1
//export http_raw
func _http_raw(mPtr, mLen, pPtr, pLen, hPtr, hLen, bPtr, bLen,
	rStatus, rHeadPtr, rHeadLen, risJson, rbodyPtr, rbodyLen int32) bool

//go:wasm-module temphia1
//export http_raw_batch
func _http_raw_batch(reqsPtr, reqsLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export http_quick_get
func _http_quick_get(uPtr, uLen, hPtr, hLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export http_quick_post
func _http_quick_post(uPtr, uLen, hPtr, hLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export http_form_post
func _http_form_post(uPtr, uLen, hPtr, hLen, dPtr, dLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export http_json_get
func _http_json_get(uPtr, uLen, hPtr, hLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export http_json_post
func _http_json_post(uPtr, uLen, hPtr, hLen, dPtr, dLen, respOffset, respLen int32) bool
