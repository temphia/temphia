package seeder

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type LiveSeeder struct {
	Core        SeederCore
	RecordCache map[string]map[int64]map[string]any
	Group       *entities.TableGroup
	Pacman      repox.Hub
	Source      store.DynSource
	UserId      string
}

func NewLiveSeeder() *LiveSeeder {

	return &LiveSeeder{
		Core:        SeederCore{},
		RecordCache: make(map[string]map[int64]map[string]any),
		Group:       nil,
		Pacman:      nil,
		Source:      nil,
		UserId:      "",
	}
}

func (l *LiveSeeder) Seed() error {

	return nil
}
