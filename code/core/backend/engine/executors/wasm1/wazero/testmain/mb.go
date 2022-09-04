package main

import (
	"log"
	"time"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
)

type mb struct{}

func (m *mb) Log(msg string)                                     { log.Println(msg) }
func (m *mb) LazyLog(msgs []string)                              { log.Println(msgs) }
func (m *mb) Sleep(msec int32)                                   { time.Sleep(time.Duration(msec * int32(time.Millisecond))) }
func (m *mb) GetFileWithMeta(file string) ([]byte, int64, error) { return nil, 0, nil }
func (m *mb) GetApp() interface{}                                { return nil }
func (m *mb) PlugKVBindingsGet() bindx.BindPlugKV                { return nil }
func (m *mb) SockdBindingsGet() bindx.BindSockd                  { return nil }
func (m *mb) UserBindingsGet() bindx.BindUser                    { return nil }
func (m *mb) CabinetBindingsGet() bindx.BindCabinet              { return nil }
func (m *mb) SelfBindingsGet() bindx.BindSelf                    { return nil }
func (m *mb) NodeCacheGet() bindx.BindNodeCache                  { return nil }

// bindx
