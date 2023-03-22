package msgbus

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type SubStore struct {
	topics     map[string][]xplane.Subscription
	mlock      sync.Mutex
	subCounter int32
}

func (br *SubStore) addSub(topic string, ch chan xplane.Message) int32 {

	br.mlock.Lock()
	defer br.mlock.Unlock()

	subs := br.topics[topic]
	if subs == nil {
		subs = make([]xplane.Subscription, 0, 2)
	}

	br.subCounter = br.subCounter + 1

	subs = append(subs, xplane.Subscription{
		Id:   br.subCounter,
		Chan: ch,
	})

	br.topics[topic] = subs
	return br.subCounter
}

func (br *SubStore) removeSub(topic string, subId int32) {

	br.mlock.Lock()
	defer br.mlock.Unlock()

	subs := br.topics[topic]
	if subs == nil {
		return
	}

	newSubs := make([]xplane.Subscription, 0, len(subs))

	for _, sub := range subs {
		if subId != (sub.Id) {
			newSubs = append(newSubs, sub)
		}
	}
	br.topics[topic] = newSubs
}

func (br *SubStore) getSubs(topic string) []xplane.Subscription {

	br.mlock.Lock()
	defer br.mlock.Unlock()

	return br.topics[topic]
}
