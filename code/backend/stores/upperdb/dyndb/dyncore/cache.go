package dyncore

import (
	"sync"

	"github.com/hashicorp/golang-lru/simplelru"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type GetColumnsFunc func(tenantId, gslug, tslug string) ([]*entities.Column, error)

type dcache struct {
	schemaCache simplelru.LRU
	scLock      sync.Mutex
	getCol      GetColumnsFunc
}

func NewCache(getCol GetColumnsFunc) *dcache {
	// lru size from config ?

	lru, err := simplelru.NewLRU(30, nil)
	if err != nil {
		return nil
	}

	return &dcache{
		scLock:      sync.Mutex{},
		schemaCache: *lru,
		getCol:      getCol,
	}
}

func (d *dcache) CachedColumns(tenantId, group, table string) (map[string]*entities.Column, error) {
	d.scLock.Lock()
	defer d.scLock.Unlock()

	val, ok := d.schemaCache.Get(tenantId + group + table)
	if ok {
		return val.(map[string]*entities.Column), nil
	}

	pCols := make(map[string]*entities.Column)

	colums, err := d.getCol(tenantId, group, table)
	if err != nil {
		return nil, err
	}

	for _, col := range colums {
		pCols[col.Slug] = col
	}

	d.schemaCache.Add(tenantId+group+table, pCols)
	return pCols, nil
}

func (d *dcache) EvictColumns(tenantId, group, table string) {
	d.scLock.Lock()
	defer d.scLock.Unlock()

	d.schemaCache.Remove(tenantId + group + table)
}
