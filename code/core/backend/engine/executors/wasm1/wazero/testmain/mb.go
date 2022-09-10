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
func (m *mb) GetApp() any                                        { return nil }
func (m *mb) PlugKVBindingsGet() bindx.PlugKV                    { return nil }
func (m *mb) SockdBindingsGet() bindx.Sockd                      { return nil }
func (m *mb) UserBindingsGet() bindx.User                        { return nil }
func (m *mb) CabinetBindingsGet() bindx.Cabinet                  { return nil }
func (m *mb) SelfBindingsGet() bindx.Self                        { return nil }
func (m *mb) NodeCacheGet() bindx.NodeCache                      { return nil }

// bindx
