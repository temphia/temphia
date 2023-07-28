package processer

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/twpayne/go-geom/encoding/wkt"
)

type PGCtypeProcesser struct {
	columns map[string]*entities.Column
}

func (pg *PGCtypeProcesser) FromRowsDBType(rows []map[string]interface{}) error {
	for _, row := range rows {
		err := pg.FromRowDBType(row)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pg *PGCtypeProcesser) FromRowDBType(row map[string]interface{}) error {

	delete(row, dyndb.KeyModSig)

	for k, v := range row {

		col := pg.columns[k]
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
			point, err := PgLocationFromDBType(v)
			if err != nil {
				return err
			}
			row[k] = point
		case dyndb.CtypeJSON:
			switch vv := v.(type) {
			case string:
				row[k] = vv
			case []uint8:
				row[k] = string(vv)
			default:
				continue
			}
		default:
			continue
		}

	}

	return nil
}

func (pg *PGCtypeProcesser) ToRowsDBType(rows []map[string]interface{}) error {

	for _, v := range rows {
		err := pg.ToRowDBType(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pg *PGCtypeProcesser) ToRowDBType(row map[string]interface{}) error {
	for k, v := range row {
		if dyndb.IsMeta(k) {
			continue
		}

		col := pg.columns[k]
		if col == nil {
			// yolo
			continue
		}

		if v == nil {
			continue
		}

		switch col.Ctype {
		case dyndb.CtypeLocation:
			row[k] = PgLocationToDBType(convertToFloat(v))
		case dyndb.CtypeJSON:

			switch vv := v.(type) {
			case map[string]any:
				out, err := json.Marshal(&vv)
				if err != nil {
					return err
				}
				row[k] = out
			default:
				continue
			}

		default:
			continue
		}

	}
	return nil
}

func PgLocationToDBType(p [2]float64) string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p[0], p[1])
}

func PgLocationFromDBType(val interface{}) ([2]float64, error) {
	var p [2]float64

	lstr := ""

	switch lval := val.(type) {
	case []uint8:
		lstr = string(lval)
	case string:
		lstr = lval
	}

	if strings.HasPrefix(lstr, "SRID=4326;POINT") {
		w, err := wkt.Unmarshal(lstr)
		if err != nil {
			return p, err
		}

		cod := w.FlatCoords()

		p[0] = cod[0]
		p[1] = cod[1]

		return p, nil
	}

	b, err := hex.DecodeString(lstr)
	if err != nil {
		pp.Println("@lstr", lstr)
		pp.Println("@hex_err", err)
		b = []byte(lstr)
	}

	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return p, err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return p, easyerr.Error(fmt.Sprintf("Invalid byte order %d", wkbByteOrder))
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return p, err
	}

	return p, binary.Read(r, byteOrder, &p)
}

func convertToFloat(val interface{}) [2]float64 {
	switch point := val.(type) {
	case []interface{}:
		return [2]float64{
			point[0].(float64),
			point[1].(float64),
		}
	case [2]float64:
		return point
	default:
		panic("not reachabale")
	}
}
