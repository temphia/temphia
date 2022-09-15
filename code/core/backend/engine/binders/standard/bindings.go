package standard

import "github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"

func (b *Binder) PlugKVBindingsGet() bindx.PlugKV   { return &b.plugKV }
func (b *Binder) SockdBindingsGet() bindx.Sockd     { return &b.sockd }
func (b *Binder) CabinetBindingsGet() bindx.Cabinet { return &b.cabinet }
func (b *Binder) NodeCacheGet() bindx.NodeCache     { return &b.ncache }
func (b *Binder) UserBindingsGet() bindx.User       { return &b.user }
func (b *Binder) SelfBindingsGet() bindx.Self       { return &b.self }
func (b *Binder) NetGet() bindx.Net                 { return &b.net }
