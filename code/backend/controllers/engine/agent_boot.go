package engine

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type BootRenderOptions struct {
	LoaderScript string
	ScriptFile   string
	StyleFile    string
	ExtScripts   map[string]string
	BootData     string
	ApiBaseURL   string
	PlugId       string
	AgentId      string
	Name         string
	ExecLoader   string
}

type BootData struct {
	ApiBaseURL string `json:"api_base_url,omitempty"`
	PlugId     string `json:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty"`
	EntryName  string `json:"entry_name,omitempty"`
	ExecLoader string `json:"exec_loader,omitempty"`
}

func (c *Controller) BootAgent(tenantId, host, plugId, agentId string) (*BootRenderOptions, error) {

	agent, err := c.corehub.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		return nil, err
	}

	aburl := httpx.ApiBaseURL(host, tenantId)

	data := BootData{
		ApiBaseURL: aburl,
		PlugId:     plugId,
		AgentId:    agentId,
		EntryName:  agent.WebEntry,
		ExecLoader: agent.WebLoader,
	}

	bdata, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return &BootRenderOptions{
		ApiBaseURL:   aburl,
		ScriptFile:   agent.WebScript,
		StyleFile:    agent.WebStyle,
		LoaderScript: "",
		PlugId:       plugId,
		AgentId:      agentId,
		ExecLoader:   agent.WebLoader,
		BootData:     string(bdata),
		ExtScripts:   map[string]string{},
		Name:         agent.Name,
	}, nil

}
