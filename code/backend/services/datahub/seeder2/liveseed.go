package seeder

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type LiveSeederOptions struct {
	TenantId  string
	Group     string
	Table     string
	UserId    string
	Source    store.DynSource
	MaxRecord int
}

type LiveSeeder struct {
	mainTable        *entities.Table
	mainTableColumns map[string]*entities.Column
	siblingTables    map[string]*entities.Table
	siblingColumns   map[string][]*entities.Column

	fileCache map[string][]string
	userCache map[string][]string

	recordCache map[string]map[int64]map[string]any
	group       *entities.TableGroup
	options     LiveSeederOptions
}

func NewLiveSeeder(opts LiveSeederOptions) (*LiveSeeder, error) {

	tgrp, err := opts.Source.GetGroup(opts.Group)
	if err != nil {
		return nil, err
	}

	tables, err := opts.Source.ListTables(opts.Group)
	if err != nil {
		return nil, err
	}

	tablesIndexed := make(map[string]*entities.Table, len(tables))
	for _, t := range tables {
		tablesIndexed[t.Slug] = t
	}

	columns, err := opts.Source.ListColumns(opts.Group, opts.Table)
	if err != nil {
		return nil, err
	}

	columnIndexed := make(map[string]*entities.Column, len(columns))
	for _, c := range columns {
		columnIndexed[c.Slug] = c
	}

	ls := &LiveSeeder{
		mainTable:        tablesIndexed[opts.Table],
		mainTableColumns: columnIndexed,
		siblingTables:    tablesIndexed,

		// lazyload these
		siblingColumns: make(map[string][]*entities.Column),
		fileCache:      make(map[string][]string),
		userCache:      make(map[string][]string),
		recordCache:    make(map[string]map[int64]map[string]any),
		group:          tgrp,
		options:        opts,
	}

	return ls, nil
}

func (l *LiveSeeder) Seed() error {
	seedRecords := make([]map[string]any, 0, (l.options.MaxRecord))

	for i := 0; i < l.options.MaxRecord; i = i + 1 {
		currmap := make(map[string]any, len(l.mainTableColumns))
		for _, col := range l.mainTableColumns {

			// fixme => column skip ?
			record := NewRecord(col, l)
			currmap[col.Slug] = record
		}

		seedRecords = append(seedRecords, currmap)
	}

	for _, record := range seedRecords {

		_, err := l.options.Source.NewRow(0, store.NewRowReq{
			TenantId: l.options.TenantId,
			Group:    l.options.Group,
			Table:    l.options.Table,
			Data:     record,
			ModCtx: store.ModCtx{
				UserId: l.options.UserId,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// private

func (s *LiveSeeder) Files(column string) []string {
	return s.fileCache[column]
}

func (s *LiveSeeder) Users(column string) []string {
	return s.userCache[column]
}
