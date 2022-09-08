package datahub

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (d *dynSource) NewRow(txid uint32, req store.NewRowReq) (int64, error) {
	ddb := d.dynDB()

	id, err := ddb.NewRow(txid, req)
	if err != nil {
		return 0, err
	}

	req.Data[store.KeyPrimary] = id

	// fixme => when txid != 0
	if txid == 0 {
		err = d.hub.sockdhub.PushNewRow(d.source, req.TenantId, req.Group, req.Table, req.Data)
		if err != nil {
			pp.Println(err)
		}
	}

	return id, nil
}

func (d *dynSource) GetRow(txid uint32, req store.GetRowReq) (map[string]any, error) {
	ddb := d.dynDB()

	if txid != 0 || req.SkipCache {
		return ddb.GetRow(txid, req)
	}
	// fixme => use cache
	return ddb.GetRow(txid, req)
}

func (d *dynSource) UpdateRow(txid uint32, req store.UpdateRowReq) (map[string]any, error) {
	ddb := d.dynDB()

	data, err := ddb.UpdateRow(txid, req)
	if err != nil {
		return nil, err
	}

	err = d.hub.sockdhub.PushUpdateRow(d.source, d.tenantId, req.Group, req.Table, req.Id, req.Data)
	if err != nil {
		pp.Println(err)
	}

	return data, nil
}

func (d *dynSource) DeleteRows(txid uint32, req store.DeleteRowReq) error {
	ddb := d.dynDB()
	// fixme => push to sockd chan

	return ddb.DeleteRows(txid, req)
}

func (d *dynSource) SimpleQuery(txid uint32, req store.SimpleQueryReq) (*store.QueryResult, error) {
	ddb := d.dynDB()

	return ddb.SimpleQuery(txid, req)
}

func (d *dynSource) FTSQuery(txid uint32, req store.FTSQueryReq) (*store.QueryResult, error) {
	ddb := d.dynDB()
	return ddb.FTSQuery(txid, req)
}
