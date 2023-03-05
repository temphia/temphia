package syncer

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type DataSyncer struct {
	sockd sockdx.SockdCore
}

func NewData(sockd sockdx.SockdCore) *DataSyncer {
	return &DataSyncer{
		sockd: sockd,
	}
}

type RowMod struct {
	Table   string  `json:"table,omitempty"`
	Rows    []int64 `json:"rows,omitempty"`
	ModType string  `json:"mod_type,omitempty"`
	Data    any     `json:"data,omitempty"`
}

func (s *DataSyncer) PushNewRow(source, tenantId, groupId, table string, data map[string]any) error {

	iid, ok := data[dyndb.KeyPrimary]
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

func (s *DataSyncer) PushUpdateRow(source, tenantId, groupId, table string, id int64, data map[string]any) error {
	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    []int64{id},
		ModType: "update",
		Data:    data,
	})
}

func (s *DataSyncer) PushDeleteRow(source, tenantId, groupId, table string, id int64) error {
	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    []int64{id},
		ModType: "delete",
		Data:    nil,
	})
}

func (s *DataSyncer) pushRowMod(source, tenantId, groupId string, data *RowMod) error {
	pp.Println("@syncer", data)

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return s.sockd.SendTagged(
		tenantId,
		sockdx.ROOM_SYS_DATA,
		[]string{fmt.Sprintf("dgroup.%s.%s", source, groupId)},
		[]int64{},
		out,
	)
}
