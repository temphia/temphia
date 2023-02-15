package sheet

import (
	"github.com/temphia/temphia/code/backend/services/datahub2/handle"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type Sheet struct {
	inner    dyndb.DynDB
	handle   *handle.Handle
	source   string
	tenantId string
	group    string
}

func New(inner dyndb.DynDB, handle *handle.Handle, source string, tenantId string, group string) *Sheet {
	return &Sheet{
		inner:    inner,
		handle:   handle,
		source:   source,
		tenantId: tenantId,
		group:    group,
	}

}

func (s *Sheet) ListSheetGroup(txid uint32) (*dyndb.ListSheetGroupResp, error)
func (s *Sheet) LoadSheet(txid uint32, data *dyndb.LoadSheetReq) (*dyndb.LoadSheetResp, error)
func (s *Sheet) ListSheet(txid uint32) ([]map[string]any, error)
func (s *Sheet) NewSheet(txid uint32, data map[string]any) error
func (s *Sheet) GetSheet(txid uint32, id int64) (map[string]any, error)
func (s *Sheet) UpdateSheet(txid uint32, id int64, data map[string]any) error
func (s *Sheet) DeleteSheet(txid uint32, id int64) error
func (s *Sheet) ListSheetColumn(txid uint32, sid int64) ([]map[string]any, error)
func (s *Sheet) NewSheetColumn(txid uint32, sid int64, data map[string]any) (int64, error)
func (s *Sheet) GetSheetColumn(txid uint32, sid, cid int64) (map[string]any, error)
func (s *Sheet) UpdateSheetColumn(txid uint32, sid, cid int64, data map[string]any) error
func (s *Sheet) DeleteSheetColumn(txid uint32, sid, cid int64) error
func (s *Sheet) NewRowWithCell(txid uint32, sid int64, data map[int64]map[string]any) (map[int64]map[string]any, error)
func (s *Sheet) UpdateRowWithCell(txid uint32, sid, rid int64, data map[int64]map[string]any) (map[int64]map[string]any, error)
