package evgoja

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type Builder struct {
	instances map[string][]*GojaInstance
	mLock     sync.Mutex
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	b := Builder{
		instances: make(map[string][]*GojaInstance),
	}

	return etypes.ExecBuilderFunc(b.Instance), nil

}

func (b *Builder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	b.mLock.Lock()
	instances, ok := b.instances[opts.TenantId]
	if !ok {
		instances = []*GojaInstance{}
		b.instances[opts.TenantId] = instances
	}
	b.mLock.Unlock()

	var i *GojaInstance

	for _, instance := range instances {
		if !instance.needsClose {
			continue
		}
	}

	if i == nil {
		i = NewInstance(opts.TenantId)
		err := i.Init(opts.Binder)
		if err != nil {
			return nil, err
		}

		b.mLock.Lock()
		b.instances[opts.TenantId] = append(b.instances[opts.TenantId], i)
		b.mLock.Unlock()
	}

	return &Goja{
		instance: i,
	}, nil
}
