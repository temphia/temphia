package main

import (
	"log"
	"os"

	reclient "github.com/temphia/temphia/code/executors/re/langs/golang"
)

func main() {

	opts := reclient.Options{
		Port:     os.Getenv("TEMPHIA_RE_PORT"),
		Token:    os.Getenv("TEMPHIA_RE_TOKEN"),
		AgentId:  os.Getenv("TEMPHIA_RE_AGENT_ID"),
		PlugId:   os.Getenv("TEMPHIA_RE_PLUG_ID"),
		TenantId: os.Getenv("TEMPHIA_RE_TENANT"),
	}

	log.Println("staring_reclient ctx:", opts)

	rc := reclient.New(opts)

	err := rc.Init()
	if err != nil {
		panic(err)
	}

}
