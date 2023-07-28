package agent

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type agentState struct {
	webFiles       map[string]string
	spaConfig      *entities.SPAOptions
	ssrConfig      *entities.SSROptions
	templateConfig any
}
