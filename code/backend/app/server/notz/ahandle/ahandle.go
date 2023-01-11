package ahandle

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

var (
	_ httpx.AdapterHandle = (*AHandle)(nil)
)

type Options struct {
	Corehub  store.CoreHub
	Logger   *zerolog.Logger
	DomainId int64
	TenantId string
}

type AHandle struct {
	corehub    store.CoreHub
	logger     *zerolog.Logger
	tenantId   string
	keyTypeKey string
}

func New(opts Options) *AHandle {
	return &AHandle{
		corehub:    opts.Corehub,
		logger:     opts.Logger,
		tenantId:   opts.TenantId,
		keyTypeKey: fmt.Sprintf("adapter-%d", opts.DomainId),
	}
}

func (ah *AHandle) KvAdd(key, value string) error {
	return ah.corehub.AddSystemKV(ah.tenantId, &entities.SystemKV{
		Key:      key,
		Type:     ah.keyTypeKey,
		Value:    value,
		TenantId: ah.tenantId,
	})
}

func (ah *AHandle) KvUpdate(key, value string) error {
	return ah.corehub.UpdateSystemKV(ah.tenantId, key, ah.keyTypeKey, map[string]any{
		"value": value,
	})
}

func (ah *AHandle) KvGet(key string) (string, error) {
	data, err := ah.corehub.GetSystemKV(ah.tenantId, key, ah.keyTypeKey)
	if err != nil {
		return "", err
	}

	return data.Value, err
}

func (ah *AHandle) KvRemove(key string) error {
	return ah.corehub.RemoveSystemKV(ah.tenantId, key, ah.keyTypeKey)
}

func (ah *AHandle) KvList(prefix string) ([]string, error) {
	resps, err := ah.corehub.ListSystemKV(ah.tenantId, ah.keyTypeKey, prefix, 0)
	if err != nil {
		return nil, err
	}

	final := make([]string, 0, len(resps))

	for _, sk := range resps {
		final = append(final, sk.Value)
	}

	return final, nil
}

// log
func (ah *AHandle) GetLogger() *zerolog.Logger {
	return ah.logger
}

func (ah *AHandle) LogInfo(rid int64) *zerolog.Event {
	return ah.logger.Info().Int64("rid", rid)
}

func (ah *AHandle) LogError(rid int64) *zerolog.Event {
	return ah.logger.Error().Int64("rid", rid)
}
