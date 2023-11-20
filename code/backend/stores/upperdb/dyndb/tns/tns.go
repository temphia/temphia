package tns

import "errors"

var (
	ErrLongSlug        = errors.New("longer slug length than allowed")
	ErrNotAllowedChars = errors.New("slug has not allowed chars")
)

type TNS interface {
	Table(tenantId, groupId, tableId string) string
	ActivityTable(tenantId, groupId, tableId string) string
	CheckGroupSlug(string) error
	CheckTableSlug(string) error
	CheckColumnSlug(string) error
	MetaTableGroup(tenantId string) string
	MetaTable(tenantId string) string
	MetaColumn(tenantId string) string
}

func New(mode string) TNS {
	switch mode {
	case "shared":
		return &tnsShared{}
	case "sharded":
		fallthrough
	default:
		panic("invalid option")
	}
}
