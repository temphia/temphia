package sockdhub

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type RowMod struct {
	Table   string  `json:"table,omitempty"`
	Rows    []int64 `json:"rows,omitempty"`
	ModType string  `json:"mod_type,omitempty"`
	Data    any     `json:"data,omitempty"`
}

func (s *SockdHub) PushNewRow(source, tenantId, groupId, table string, data map[string]any) error {

	iid, ok := data[store.KeyPrimary]
	if !ok {
		pp.Println("row id not found ", data)
		return nil
	}

	id := iid.(int64)

	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    []int64{id},
		ModType: "insert",
		Data:    data,
	})
}

func (s *SockdHub) PushUpdateRow(source, tenantId, groupId, table string, id int64, data map[string]any) error {
	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    []int64{id},
		ModType: "update",
		Data:    data,
	})
}

func (s *SockdHub) PushDeleteRow(source, tenantId, groupId, table string, id int64) error {
	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    []int64{id},
		ModType: "delete",
		Data:    nil,
	})
}

func (s *SockdHub) pushRowMod(source, tenantId, groupId string, data *RowMod) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return s.sockd.SendTagged(
		tenantId,
		ROOM_SYSTABLE,
		[]string{fmt.Sprintf("dgroup.%s.%s", source, groupId)},
		[]int64{},
		out,
	)
}
