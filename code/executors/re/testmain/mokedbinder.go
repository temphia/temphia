package main

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type MB struct {
	bindx.Bindings
}

func (m *MB) GetFileWithMeta(file string) (data []byte, version int64, err error) {
	out, err := getTestZip()
	return out, 0, err
}

func (m *MB) Clone() bindx.Core {

	return m
}
