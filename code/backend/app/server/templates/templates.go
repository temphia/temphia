package templates

import (
	_ "embed"
)

//go:embed agent_boot.html
var AgentBootTemplate []byte
