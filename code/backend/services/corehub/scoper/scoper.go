package scoper

import (
	"strings"
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Manager struct {
	scopes map[string]*GroupScoper
	mLock  sync.RWMutex
}

func New() Manager {
	return Manager{
		scopes: make(map[string]*GroupScoper),
		mLock:  sync.RWMutex{},
	}
}

func (s *Manager) Get(key string) *GroupScoper {
	s.mLock.RLock()
	gs := s.scopes[key]
	s.mLock.RUnlock()
	return gs
}

func (s *Manager) Set(key string, model *entities.UserGroup) {

	gs := &GroupScoper{
		scopes: make(map[string]struct{}),
	}

	for _, scope := range strings.Split(model.Scopes, ",") {
		gs.scopes[scope] = struct{}{}
	}

	s.scopes[key] = gs
}
