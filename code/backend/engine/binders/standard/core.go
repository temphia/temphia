package standard

import (
	"encoding/json"
	"time"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx/logid"
)

func (b *Binder) Log(msg string) {

	b.logInfo().
		Str("message_data", msg).
		Msg(logid.BinderExecutionLog)

	b.logDebugRoom(&etypes.DebugMessage{
		Messages: []string{msg},
		EventId:  b.Handle.EventId,
		PlugId:   b.Handle.PlugId,
		AgentId:  b.Handle.AgentId,
	})
}

func (b *Binder) LazyLog(msgs []string) {
	b.logInfo().
		Strs("message_datas", msgs).
		Msg(logid.BinderExecutionLog)

	b.logDebugRoom(&etypes.DebugMessage{
		Messages: msgs,
		EventId:  b.Handle.EventId,
		PlugId:   b.Handle.PlugId,
		AgentId:  b.Handle.AgentId,
	})
}

func (b *Binder) Sleep(msec int32) {
	time.Sleep(time.Millisecond * time.Duration(msec))
}

func (b *Binder) GetFileWithMeta(file string) ([]byte, int64, error) {
	out, err := b.self.SelfGetFile(file)

	return out, 0, err
}

func (b *Binder) GetApp() any {
	return b.Handle.Deps.App
}

// private

func (b *Binder) logDebugRoom(msg *etypes.DebugMessage) {
	out, err := json.Marshal(msg)
	if err != nil {
		pp.Println(err)
		return
	}

	b.Handle.Deps.Sockd.SendBroadcast(b.Handle.Namespace, "plugs_dev", []int64{}, out)
}
