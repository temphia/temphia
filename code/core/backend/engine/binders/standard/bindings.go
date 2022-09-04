package standard

import "github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"

func (b *Binder) PlugKVBindingsGet() bindx.BindPlugKV   { return &b.plugKV }
func (b *Binder) SockdBindingsGet() bindx.BindSockd     { return &b.sockd }
func (b *Binder) CabinetBindingsGet() bindx.BindCabinet { return &b.cabinet }
func (b *Binder) NodeCacheGet() bindx.BindNodeCache     { return &b.ncache }
func (b *Binder) UserBindingsGet() bindx.BindUser       { return &b.user }
func (b *Binder) SelfBindingsGet() bindx.BindSelf       { return &b.self }
