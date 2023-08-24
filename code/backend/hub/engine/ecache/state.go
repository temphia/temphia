package ecache

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type agentstate struct {
	wchan    chan *entities.Agent
	plug     string
	agent    string
	tenantId string
}
