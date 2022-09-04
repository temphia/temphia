package tns

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/thoas/go-funk"
)

type tnsShared struct{}

func (t *tnsShared) Table(tenantId, groupId, tableId string) string {
	return fmt.Sprintf("%s_%s_%s", tenantId, groupId, tableId)
}

func (t *tnsShared) MetaTableGroup(tenantId string) string { return "data_table_groups" }
func (t *tnsShared) MetaTable(tenantId string) string      { return "data_tables" }
func (t *tnsShared) MetaColumn(tenantId string) string     { return "data_table_columns" }

func (t *tnsShared) CheckGroupSlug(slug string) error {
	if len(slug) > 10 {
		return ErrLongSlug
	}

	if !checkSlug(slug) {
		return ErrNotAllowedChars
	}
	return nil
}

func (t *tnsShared) CheckTableSlug(slug string) error {
	if len(slug) > 10 {
		return ErrLongSlug
	}

	if !checkSlug(slug) {
		return ErrNotAllowedChars
	}

	return nil
}

func (t *tnsShared) CheckColumnSlug(slug string) error {
	if len(slug) > 10 {
		return ErrLongSlug
	}
	if !checkSlug(slug) {
		return ErrNotAllowedChars
	}

	return nil
}

func checkSlug(s string) bool {
	return s != "" && !funk.ContainsString(store.ReservedSlugs, s) && store.SlugExp.MatchString(s)
}

func DataTable(tenent_id, gslug, tslug string) string {
	return fmt.Sprintf("%s_%s_%s", tenent_id, gslug, tslug)
}

func DataGroup(tenent_id, gslug string) string {
	return fmt.Sprintf("%s_%s", tenent_id, gslug)
}
