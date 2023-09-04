package notz

func (a *Notz) getBid(tenantId, plugId string) string {

	a.bLock.RLock()
	bid := a.bprintIdx[plugId]
	a.bLock.RUnlock()

	if bid != "" {
		return bid
	}

	plug, err := a.corehub.PlugGet(tenantId, plugId)
	if err != nil {
		return ""
	}

	a.bLock.Lock()
	a.bprintIdx[plugId] = plug.BprintId
	a.bLock.Unlock()

	return plug.BprintId
}
