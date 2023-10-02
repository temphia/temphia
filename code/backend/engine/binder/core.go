package binder

import (
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
)

func (b *Binder) Log(eid, msg string) {

	b.logInfo().
		Str("message_data", msg).
		Msg(logid.BinderExecutionLog)

	b.logDebugRoom(&etypes.DebugMessage{
		Messages: []string{msg},
		EventId:  eid,
		PlugId:   b.PlugId,
		AgentId:  b.AgentId,
	})
}

func (b *Binder) LazyLog(eid string, msgs []string) {
	b.logInfo().
		Strs("message_datas", msgs).
		Msg(logid.BinderExecutionLog)

	b.logDebugRoom(&etypes.DebugMessage{
		Messages: msgs,
		EventId:  eid,
		PlugId:   b.PlugId,
		AgentId:  b.AgentId,
	})
}

func (b *Binder) Sleep(msec int32) {
	time.Sleep(time.Millisecond * time.Duration(msec))
}

func (b *Binder) GetFileWithMeta(file string) ([]byte, int64, error) {
	bstore := b.Deps.Pacman.GetBprintFileStore()

	// fixme => folder support ?

	out, err := bstore.GetBlob(b.Namespace, b.BprintId, "", file)
	return out, 0, err
}

func (b *Binder) GetApp() any {
	return b.Deps.App
}

// private
