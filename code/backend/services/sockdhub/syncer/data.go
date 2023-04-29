package syncer

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

const (
	DataModTypeInsert  = "insert"
	DataModTypeUpdate  = "update"
	DataModTypeDelete  = "delete"
	DataModTypeComment = "comment"

	DataModTypeSheetInsert = "sheet_insert"
	DataModTypeSheetUpdate = "sheet_update"
	DataModTypeSheetDelete = "sheet_delete"
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

// datatable

func (s *DataSyncer) PushNewRow(source, tenantId, groupId, table string, ids []int64, data any) error {
	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    ids,
		ModType: DataModTypeInsert,
		Data:    data,
	})

}

func (s *DataSyncer) PushUpdateRow(source, tenantId, groupId, table string, ids []int64, data any) error {
	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    ids,
		ModType: DataModTypeUpdate,
		Data:    data,
	})
}

func (s *DataSyncer) PushDeleteRow(source, tenantId, groupId, table string, ids []int64) error {
	return s.pushRowMod(source, tenantId, groupId, &RowMod{
		Table:   table,
		Rows:    ids,
		ModType: DataModTypeDelete,
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

// datasheet

type RowSheetMod struct {
	SheetId int64   `json:"sheet_id,omitempty"`
	Rows    []int64 `json:"rows,omitempty"`
	ModType string  `json:"mod_type,omitempty"`
	Data    any     `json:"data,omitempty"`
}

func (s *DataSyncer) PushSheetNewRow(source, tenantId, groupId string, sheetId int64, ids []int64, data any) error {

	return s.pushSheetRowMod(source, tenantId, groupId, &RowSheetMod{
		SheetId: sheetId,
		Rows:    ids,
		ModType: DataModTypeSheetInsert,
		Data:    data,
	})
}

func (s *DataSyncer) PushSheetUpdateRow(source, tenantId, groupId string, sheetId int64, ids []int64, data any) error {
	return s.pushSheetRowMod(source, tenantId, groupId, &RowSheetMod{
		SheetId: sheetId,
		Rows:    ids,
		ModType: DataModTypeSheetUpdate,
		Data:    data,
	})
}

func (s *DataSyncer) PushSheetDeleteRow(source, tenantId, groupId string, sheetId int64, ids []int64) error {
	return s.pushSheetRowMod(source, tenantId, groupId, &RowSheetMod{
		SheetId: sheetId,
		Rows:    ids,
		ModType: DataModTypeSheetDelete,
		Data:    nil,
	})
}

func (s *DataSyncer) pushSheetRowMod(source, tenantId, groupId string, data *RowSheetMod) error {

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
