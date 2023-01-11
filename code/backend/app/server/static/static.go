package static

import _ "embed"

//go:embed root.html
var Root []byte

//go:embed operator.html
var Operator []byte

//go:embed portal.html
var Portal []byte
