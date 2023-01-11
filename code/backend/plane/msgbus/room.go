package msgbus

import (
	"strings"
	"sync"
)

type BusRoom struct {
	topics     map[string][]Subscription
	mlock      sync.RWMutex
	subCounter int32
}

func (br *BusRoom) AddSub(topic string, cb func(msg Message)) int32 {
	topics := strings.Split(topic, ".")

	br.mlock.Lock()
	defer br.mlock.Unlock()

	subs := br.topics[topics[0]]
	if subs == nil {
		subs = make([]Subscription, 0, 2)
	}

	br.subCounter = br.subCounter + 1

	subs = append(subs, Subscription{
		Id:   br.subCounter,
		Path: topics[1:],
		Func: cb,
	})

	br.topics[topics[0]] = subs
	return br.subCounter
}

func (br *BusRoom) RemoveSub(topic string, subId int32) {
	topics := strings.Split(topic, ".")

	br.mlock.Lock()
	defer br.mlock.Unlock()

	subs := br.topics[topics[0]]
	if subs == nil {
		return
	}

	newSubs := make([]Subscription, 0, len(subs))

	for _, sub := range subs {
		if subId != (sub.Id) {
			newSubs = append(newSubs, sub)
		}
	}
	br.topics[topics[0]] = newSubs
}

func (br *BusRoom) SendSub(topic string, msg Message) bool {
	topics := strings.Split(topic, ".")

	br.mlock.RLock()
	defer br.mlock.RUnlock()

	subs := br.topics[topics[0]]
	if len(subs) == 0 {
		return false
	}

	for _, sub := range subs {
		// fixme => check sub topics

		sub.Func(msg)
	}

	return true

}
