package python

import (
	_ "embed"
)

// go:embed lib.py
var Lib []byte

// go:embed bootstrap.sh
var Bootstrap []byte
