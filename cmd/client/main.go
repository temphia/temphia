package main

import (
	"github.com/k0kubun/pp"
	devapi "github.com/temphia/temphia/code/client/dev_api"
)

const TOKEN = "BAYHFdqaLD4vLox5tkrqyi32hsrFYTXvMi4v5ZJwURhdXTWlDPswE8G1Bl4X5OqSuFL1DeQxU1wZWA7pB9aRDrwySk04b9UWjDhWUd7ew3XknSR0FvO96i9WDQCBtbB0BbXSEqsb2CbzdSyJgIcj91AsTQE9FsPVWSpSYKAZW9y7fjBRVU7LTEyFlUdP7Y4kx6uS6tTdSwrNSvAcZzDQ965WL5"
const API = "http://temphia.local:4000/z/api/default0/v2"

func main() {

	api := devapi.New(API, TOKEN)

	pp.Println(api.BprintFileList())
	pp.Println(api.BprintFilePush("test.txt", []byte(`hello world`)))

}
