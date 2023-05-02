package index

import (
	"encoding/json"
	"os"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/thoas/go-funk"
)

type Indexer struct {
	db   *DB
	file string
}

func New(file string) *Indexer {

	db := &DB{
		GroupIndex: make(map[string][]string),
		TagIndex:   make(map[string][]string),
		Items:      make(map[string]repox.BPrint),
	}

	fout, err := os.ReadFile(file)
	if err == nil {
		err := json.Unmarshal(fout, db)
		if err != nil {
			panic(err)
		}
	}

	return &Indexer{
		db:   db,
		file: file,
	}
}

func (dbi *Indexer) UpdateItemIndex(bprint *xbprint.LocalBprint, alias, version string) error {

	typeEntries, ok := dbi.db.GroupIndex[bprint.Type]
	if !ok {
		typeEntries = make([]string, 0)
	}

	if !funk.ContainsString(typeEntries, alias) {
		dbi.db.GroupIndex[bprint.Type] = append(typeEntries, alias)
	}

	versions := []string{}
	old, ok := dbi.db.Items[alias]
	if ok {
		versions = old.Versions
	}

	if !funk.ContainsString(versions, version) {
		versions = append(versions, version)
	}

	item := repox.BPrint{
		Name:        bprint.Name,
		Slug:        bprint.Slug,
		Type:        bprint.Type,
		SubType:     "", // fixme
		Description: bprint.Description,
		Icon:        bprint.Icon,
		Versions:    versions,
		Tags:        bprint.Tags,
	}

	dbi.db.Items[alias] = item

	// fixme => also index tags

	return nil
}

func (dbi *Indexer) Save() error {
	out, err := json.MarshalIndent(dbi.db, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbi.file, out, 0755)
}

type DB struct {
	GroupIndex map[string][]string     `json:"group_index" yaml:"group_index"`
	TagIndex   map[string][]string     `json:"tag_index" yaml:"tag_index"`
	Items      map[string]repox.BPrint `json:"items" yaml:"items"`
}
