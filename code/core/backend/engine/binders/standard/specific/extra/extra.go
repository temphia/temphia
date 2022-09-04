package extra

/*

func (b *Binder) loadResource() {
	ress, err := b.factory.App.CoreHub().ResourceListByAgent(b.namespace, b.PlugId, b.AgentId)
	if err != nil {
		panic(err)
	}

	b.resources = make(map[string]*entities.Resource, len(ress))

	for _, res := range ress {
		b.resources[res.Id] = res
	}
}

func (b *Binder) mustResolveResource(name, rtype string) string {
	if b.resources == nil {
		b.loadResource()
	}

	res := b.resources[name]
	if res == nil {
		panic("Resource not found")
	}

	if res.Type != rtype {
		panic("Wrong resource type")
	}

	return name
}

func (b *Binder) getResource(name string) (*entities.Resource, error) {
	if b.resources == nil {
		b.loadResource()
	}

	res := b.resources[name]
	if res == nil {
		return nil, easyerr.NotFound()
	}

	return res, nil
}

*/

/*


var selfLocker = locker{
	locks: make(map[string]struct{}),
	llock: sync.Mutex{},
}

var resLocker = locker{
	locks: map[string]struct{}{},
	llock: sync.Mutex{},
}

func (b *Binder) selfLKey(key string) string {
	return b.namespace + b.PlugId + key
}

func (b *Binder) SelfLockWait(key string) error {
	return selfLocker.LockWatch(b.selfLKey(key))

}
func (b *Binder) SelfLock(key string, expiry int) error {
	return selfLocker.Lock(b.selfLKey(key), expiry)
}

func (b *Binder) SelfLockRenew(key string, expiry int) error {
	return selfLocker.LockRenew(b.selfLKey(key), expiry)
}

func (b *Binder) SelfUnLock(key string) error {
	return selfLocker.UnLock(b.selfLKey(key))
}

func (b *Binder) resLKey(res, key string) string {
	return b.namespace + res + key
}

func (b *Binder) ResourceLockWait(resource string, key string) error {
	return resLocker.LockWatch(b.resLKey(resource, key))
}

func (b *Binder) ResourceLock(resource string, key string, expiry int) error {
	return resLocker.Lock(b.resLKey(resource, key), expiry)
}

func (b *Binder) ResourceLockRenew(resource string, key string, expiry int) error {
	return resLocker.LockRenew(b.resLKey(resource, key), expiry)
}

func (b *Binder) ResourceUnLock(resource string, key string) error {
	return resLocker.UnLock(b.resLKey(resource, key))
}



type locker struct {
	locks map[string]struct{}
	llock sync.Mutex
}

func (l *locker) LockWatch(key string) error {
	return nil
}
func (l *locker) Lock(key string, expiry int) error {

	l.llock.Lock()
	defer l.llock.Unlock()
	_, ok := l.locks[key]
	if ok {
		return easyerr.Error("already locked")
	}

	return nil
}

func (l *locker) LockRenew(key string, expiry int) error {
	return nil
}

func (l *locker) UnLock(key string) error {
	l.llock.Lock()
	defer l.llock.Unlock()
	_, ok := l.locks[key]
	if !ok {
		return easyerr.Error("already unlocked")
	}

	delete(l.locks, key)

	return nil
}



func (b *Binder) loadResource() {
	ress, err := b.factory.App.CoreHub().ResourceListByAgent(b.namespace, b.PlugId, b.AgentId)
	if err != nil {
		panic(err)
	}

	b.resources = make(map[string]*entities.Resource, len(ress))

	for _, res := range ress {
		b.resources[res.Id] = res
	}
}

func (b *Binder) mustResolveResource(name, rtype string) string {
	if b.resources == nil {
		b.loadResource()
	}

	res := b.resources[name]
	if res == nil {
		panic("Resource not found")
	}

	if res.Type != rtype {
		panic("Wrong resource type")
	}

	return name
}

func (b *Binder) getResource(name string) (*entities.Resource, error) {
	if b.resources == nil {
		b.loadResource()
	}

	res := b.resources[name]
	if res == nil {
		return nil, easyerr.NotFound()
	}

	return res, nil
}



*/
