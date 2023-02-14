package datahub

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
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

func (d *dynSource) DeleteRowBatch(txid uint32, req store.DeleteRowBatchReq) error {
	ddb := d.dynDB()
	return ddb.DeleteRowBatch(txid, req)
}

func (d *dynSource) DeleteRowMulti(txid uint32, req store.DeleteRowMultiReq) error {
	ddb := d.dynDB()
	return ddb.DeleteRowMulti(txid, req)
}

func (d *dynSource) DeleteRow(txid uint32, req store.DeleteRowReq) error {
	ddb := d.dynDB()
	return ddb.DeleteRow(txid, req)
}

func (d *dynSource) SimpleQuery(txid uint32, req store.SimpleQueryReq) (*store.QueryResult, error) {
	ddb := d.dynDB()

	if req.Count == 0 {
		req.Count = store.DefaultQueryFetchCount
	}

	return ddb.SimpleQuery(txid, req)
}

func (d *dynSource) LoadTable(txid uint32, req store.LoadTableReq) (*store.LoadTableResp, error) {
	ddb := d.dynDB()

	views, err := ddb.ListView(req.TenantId, req.Group, req.Table)
	if err != nil {
		return nil, err
	}

	sqr := store.SimpleQueryReq{
		TenantId:    req.TenantId,
		Table:       req.Table,
		Group:       req.Group,
		Count:       50,
		FilterConds: make([]*store.FilterCond, 0),
		Page:        0,
		Selects:     nil,
		SearchTerm:  "",
	}

	finalResp := &store.LoadTableResp{
		ReverseRefs:   nil,
		Views:         views,
		DataWidgets:   nil,
		ActiveView:    "",
		FolderTickets: make(map[string]string),
		UserTickets:   make(map[string]string),
	}

	if req.ViewFilters != nil {
		sqr.FilterConds = req.ViewFilters
	} else if req.View != "" {
		for _, view := range views {
			if view.Name == req.View {
				sqr.Count = view.Count
				fconds, err := json.Marshal(view.FilterConds)
				if err == nil {
					json.Unmarshal(fconds, &sqr.FilterConds)
				}
				sqr.Selects = view.Selects
				finalResp.ActiveView = view.Name
				break
			}
		}
	}

	sqresp, err := d.SimpleQuery(txid, sqr)
	if err != nil {
		return nil, err
	}
	finalResp.QueryResponse = sqresp

	apps, err := d.hub.corehub.ListTargetAppByType(req.TenantId, entities.TargetAppTypeDataTableWidget, fmt.Sprintf("%s/%s", req.Group, req.Table))
	if err == nil {
		finalResp.DataWidgets = apps
	}

	refCols, err := d.ListReverseColumnRef(req.Group, req.Table)
	if err == nil {
		finalResp.ReverseRefs = refCols
	}

	pp.Println("@final_resp", finalResp)

	return finalResp, nil

}

func (d *dynSource) FTSQuery(txid uint32, req store.FTSQueryReq) (*store.QueryResult, error) {
	ddb := d.dynDB()
	return ddb.FTSQuery(txid, req)
}
