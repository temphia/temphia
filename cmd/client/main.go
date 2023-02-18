package main

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/goclient/devcli"
)

const TOKEN = "BAYHFdqaLD4vLox5tkrqyi32hsrFYTXvMi4v5ZJwURhdXTWlDPswE8G1Bl4X5OqSuFL1DeQxU1wZWA7pB9aRDrwySk04b9UWjDhWUd7ew3XknSR0FvO96i9WDQCBtbB0BbXSEqsb2CbzdSyJgIcj91AsTQE9FsPVWSpSYKAZW9y7fjBRVU7LTEyFlUdP7Y4kx6uS6tTdSwrNSvAcZzDQ965WL5"
const API = "http://temphia.local:4000/z/api/default0/v2"

func main() {

	os.Setenv("TEMPHIA_BDEV_TOKEN", TOKEN)
	os.Setenv("TEMPHIA_BDEV_API_URL", API)
	os.Setenv("TEMPHIA_BDEV_PLUG_ID", "cfn44d0m4q7f892bpuu0")
	os.Setenv("TEMPHIA_BDEV_AGENT_ID", "default")

	cli := devcli.New()
	pp.Println(cli.Process())

}
