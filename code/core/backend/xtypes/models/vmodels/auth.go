package vmodels

type PlugData struct {
	Name       string `json:"name,omitempty"`
	PlugId     string `json:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty"`
	ExecTicket string `json:"exec_ticket,omitempty"`
}

type UserInfo struct {
	Name  string `json:"name,omitempty"`
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Group string `json:"group,omitempty"`
}
