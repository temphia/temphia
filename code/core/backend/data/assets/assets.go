package assets

import (
	"embed"
)

//go:embed build
//go:embed lib
//go:embed static
var DataDir embed.FS
