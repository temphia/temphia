package store

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type PlugStateKV interface {
	Set(txid uint32, tenantId, plugId, key, value string, opts *SetOptions) error
	Update(txid uint32, tenantId, plugId, key, value string, opts *UpdateOptions) error
	Get(txid uint32, tenantId, plugId, key string) (*entities.PlugKV, error)
	Query(txid uint32, tenantId, plugId string, query *PkvQuery) ([]*entities.PlugKV, error)
	Del(txid uint32, tenantId, plugId, key string) error
	DelBatch(txid uint32, tenantId, plugId string, keys []string) error

	NewTxn() (uint32, error)
	RollBack(txid uint32) error
	Commit(txid uint32) error
}

type SetOptions struct {
	Tag1 string `json:"tag1,omitempty"`
	Tag2 string `json:"tag2,omitempty"`
	Tag3 string `json:"tag3,omitempty"`
	TTL  int    `json:"ttl,omitempty"`
}

type UpdateOptions struct {
	ForceVer    bool   `json:"force_ver,omitempty"`
	WithVerison bool   `json:"with_version,omitempty"`
	Version     int    `json:"version,omitempty"`
	SetTag1     bool   `json:"set_tag1,omitempty"`
	SetTag2     bool   `json:"set_tag2,omitempty"`
	SetTag3     bool   `json:"set_tag3,omitempty"`
	Tag1        string `json:"tag1,omitempty"`
	Tag2        string `json:"tag2,omitempty"`
	Tag3        string `json:"tag3,omitempty"`
	SetTTL      bool   `json:"set_ttl,omitempty"`
	TTL         int    `json:"ttl,omitempty"`
}

type PkvQuery struct {
	KeyPrefix string   `json:"key_prefix,omitempty"`
	LoadMeta  bool     `json:"load_meta,omitempty"`
	Tag1s     []string `json:"tag1s,omitempty"`
	Tag2s     []string `json:"tag2s,omitempty"`
	Tag3s     []string `json:"tag3s,omitempty"`
	PageCount uint     `json:"page_count,omitempty"`
	Page      uint     `json:"page,omitempty"`
	KeyCursor string   `json:"key_cursor,omitempty"`
}
