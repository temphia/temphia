package processer

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
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

	delete(row, store.KeyModSig)

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
		case store.CtypeShortText:
		case store.CtypePhone:
		case store.CtypeSelect:
		case store.CtypeRFormula:
		case store.CtypeFile:
		case store.CtypeMultiFile:
		case store.CtypeCheckBox:
		case store.CtypeCurrency:
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
		case store.CtypeNumber:
		case store.CtypeLocation:
			point, err := PgLocationFromDBType(v)
			if err != nil {
				return err
			}
			row[k] = point
		case store.CtypeDateTime:
		case store.CtypeMultSelect:
		case store.CtypeLongText:
		case store.CtypeSingleUser:
		case store.CtypeMultiUser:
		case store.CtypeEmail:
		case store.CtypeJSON:
		case store.CtypeRangeNumber:
		case store.CtypeColor:

		default:
			panic("not implemented")
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
		if store.IsMeta(k) {
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
		case store.CtypeShortText:
		case store.CtypePhone:
		case store.CtypeSelect:
		case store.CtypeRFormula:
		case store.CtypeFile:
		case store.CtypeMultiFile:
		case store.CtypeCheckBox:
		case store.CtypeCurrency:
			//
		case store.CtypeNumber:
		case store.CtypeLocation:
			row[k] = PgLocationToDBType(convertToFloat(v))
		case store.CtypeDateTime:

		case store.CtypeMultSelect:
		case store.CtypeLongText:
		case store.CtypeSingleUser:
		case store.CtypeMultiUser:
		case store.CtypeEmail:
		case store.CtypeJSON:
		case store.CtypeRangeNumber:
		case store.CtypeColor:

		default:
			panic("not implemented")
		}

	}
	return nil
}

func PgLocationToDBType(p [2]float64) string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p[0], p[1])
}

func PgLocationFromDBType(val interface{}) ([2]float64, error) {
	var p [2]float64

	b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return p, err
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
		return p, fmt.Errorf("Invalid byte order %d", wkbByteOrder)
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
