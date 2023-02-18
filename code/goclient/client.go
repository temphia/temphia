package client

import "os"

type BprintDevContext struct {
	Token   string `json:"token,omitempty"`
	APIURL  string `json:"api_url,omitempty"`
	PlugId  string `json:"plug_id,omitempty"`
	AgentId string `json:"agent_id,omitempty"`
}

func ReadContext() *BprintDevContext {

	return &BprintDevContext{
		Token:   os.Getenv("TEMPHIA_BDEV_TOKEN"),
		APIURL:  os.Getenv("TEMPHIA_BDEV_API_URL"),
		PlugId:  os.Getenv("TEMPHIA_BDEV_PLUG_ID"),
		AgentId: os.Getenv("TEMPHIA_BDEV_AGENT_ID"),
	}
}
