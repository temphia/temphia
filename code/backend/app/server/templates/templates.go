package templates

import (
	_ "embed"
)

//go:embed agent_boot.html
var AgentBoot []byte

//go:embed portal.html
var Portal []byte

//go:embed root.html
var Root []byte
