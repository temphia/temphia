package seeder

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type SeedType interface {
	Files(column string) []string
	Users(column string) []string
}

func NewRecord(c *entities.Column, stype SeedType) any {

	switch c.Ctype {
	case dyndb.CtypeShortText:

		switch c.Slug {
		case "name", "title":
			return gofakeit.Name()
		case "addr":
			return gofakeit.Address().Address
		default:
			return gofakeit.HipsterWord()
		}

	case dyndb.CtypeLongText:
		return gofakeit.HipsterSentence(20)
	case dyndb.CtypePhone:
		return gofakeit.Phone()
	case dyndb.CtypeSelect, dyndb.CtypeMultSelect:
		if c.Options != nil {
			return gofakeit.RandomString(c.Options)
		}

	case dyndb.CtypeFile, dyndb.CtypeMultiFile:
		opts := stype.Files(c.Slug)
		if len(opts) > 0 {
			return gofakeit.RandomString(stype.Files(c.Slug))
		}

	case dyndb.CtypeCheckBox:
		return gofakeit.Bool()
	case dyndb.CtypeCurrency:
		return gofakeit.Price(10, 200)
	case dyndb.CtypeNumber:
		return gofakeit.Number(0, 400)
	case dyndb.CtypeLocation:
		return [2]float64{gofakeit.Latitude(), gofakeit.Longitude()}
	case dyndb.CtypeDateTime:
		return gofakeit.Date().UTC()
	case dyndb.CtypeSingleUser, dyndb.CtypeMultiUser:
		opts := stype.Users(c.Slug)
		if len(opts) > 0 {
			return gofakeit.RandomString(opts)
		}
	case dyndb.CtypeEmail:
		return gofakeit.Email()
	case dyndb.CtypeJSON:
		return "{}"
	case dyndb.CtypeRangeNumber:
		// fixme => get ranges from column
		return gofakeit.Price(40, 130)
	case dyndb.CtypeColor:
		return gofakeit.HexColor()
	default:
		//		case dyndb.CtypeRFormula:
		// if !nullables[c.Slug] {
		// 	data[c.Slug] = "1 + 1"
		// }
		return nil
	}

	return nil

}
