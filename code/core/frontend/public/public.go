package public

import (
	"embed"
)

//go:embed build/*
var BuildFiles embed.FS
