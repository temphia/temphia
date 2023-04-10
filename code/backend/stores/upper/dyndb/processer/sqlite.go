package processer

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type SqliteCtypeProcesser struct {
	columns map[string]*entities.Column
}

func (scp *SqliteCtypeProcesser) FromRowsDBType(rows []map[string]interface{}) error {

	for _, row := range rows {
		err := scp.FromRowDBType(row)
		if err != nil {
			return err
		}
	}

	return nil
}

func (scp *SqliteCtypeProcesser) FromRowDBType(row map[string]interface{}) error {

	delete(row, dyndb.KeyModSig)

	for k, v := range row {

		col := scp.columns[k]
		if col == nil {
			// yolo
			continue
		}

		if v == nil {
			continue
		}

		switch col.Ctype {

		case dyndb.CtypeCurrency:
			fstr := ""

			switch vv := v.(type) {
			case string:
				fstr = vv
			case []uint8:
				fstr = string(vv)
			default:
				continue
			}

			s, err := strconv.ParseFloat(fstr, 64)
			if err != nil {
				return err
			}
			row[k] = s
		case dyndb.CtypeNumber:
		case dyndb.CtypeLocation:
			var lstr []byte
			switch vv := v.(type) {
			case string:
				lstr = []byte(vv)
			case []uint8:
				lstr = vv
			default:
				continue
			}

			point := GeoJSON{}
			err := json.Unmarshal(lstr, &point)
			if err != nil {
				return err
			}
			row[k] = point.Coordinates
		default:
			continue
		}

	}

	return nil

}

func (scp *SqliteCtypeProcesser) ToRowsDBType(rows []map[string]interface{}) error {

	for _, v := range rows {
		err := scp.ToRowDBType(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (scp *SqliteCtypeProcesser) ToRowDBType(row map[string]interface{}) error {

	for k, v := range row {
		if dyndb.IsMeta(k) {
			continue
		}

		col := scp.columns[k]
		if col == nil {
			// yolo
			continue
		}

		if v == nil {
			continue
		}

		switch col.Ctype {

		case dyndb.CtypeJSON:
			switch v.(type) {
			case map[string]any:
				bytes, err := json.Marshal(k)
				if err != nil {
					pp.Println("@err", err)
					continue
				}

				row[k] = string(bytes)
			}

		case dyndb.CtypeLocation:

			switch vv := v.(type) {
			case []any:
				row[k] = fmt.Sprintf(`{"type":"Point", "coordinates":[%v, %v]}`, vv[0], vv[1])
			default:
				panic(fmt.Sprintf("Wrong location type key: %v value: %v", k, v))
			}

		default:
			continue
		}

	}

	return nil

}

type GeoJSON struct {
	Type        string     `json:"type,omitempty"`
	Coordinates [2]float32 `json:"coordinates,omitempty"`
}
