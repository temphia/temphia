package authz

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Manager struct {
	scopes map[string]*AuthZ
	mLock  sync.RWMutex
}

func New() Manager {
	return Manager{
		scopes: make(map[string]*AuthZ),
		mLock:  sync.RWMutex{},
	}
}

func (s *Manager) Get(key string) *AuthZ {
	s.mLock.RLock()
	gs := s.scopes[key]
	s.mLock.RUnlock()
	return gs
}

func (s *Manager) Set(key string, model *entities.UserGroup) {

	gs := &AuthZ{
		scoper: scoper{
			scopes: make(map[string]struct{}),
		},
	}

	gs.scoper.build(model.Scopes)

	s.scopes[key] = gs
}
