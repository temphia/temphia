package templates

import (
	_ "embed"
)

//go:embed portal.html
var Portal []byte

//go:embed root.html
var Root []byte
