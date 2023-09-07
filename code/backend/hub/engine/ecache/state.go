package ecache

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type agentState struct {
	wchan    chan *entities.Agent
	plug     string
	agent    string
	tenantId string
}

type plugState struct {
	wchan    chan *entities.Plug
	plug     string
	tenantId string
}
