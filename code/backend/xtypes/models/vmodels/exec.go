package vmodels

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type ExecutorData struct {
	Plug   *entities.Plug
	Agent  *entities.Agent
	Bprint *entities.BPrint
}
