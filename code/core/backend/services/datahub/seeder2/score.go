package seeder

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type SeedType interface {
	Files(column string) []string
	Users(column string) []string
}

func NewRecord(c *entities.Column, stype SeedType) any {

	switch c.Ctype {
	case store.CtypeShortText:

		switch c.Slug {
		case "name", "title":
			return gofakeit.Name()
		case "addr":
			return gofakeit.Address().Address
		default:
			return gofakeit.HipsterWord()
		}

	case store.CtypeLongText:
		return gofakeit.HipsterSentence(20)
	case store.CtypePhone:
		return gofakeit.Phone()
	case store.CtypeSelect, store.CtypeMultSelect:
		if c.Options != nil {
			return gofakeit.RandomString(c.Options)
		}

	case store.CtypeFile, store.CtypeMultiFile:
		opts := stype.Files(c.Slug)
		if len(opts) > 0 {
			return gofakeit.RandomString(stype.Files(c.Slug))
		}

	case store.CtypeCheckBox:
		return gofakeit.Bool()
	case store.CtypeCurrency:
		return gofakeit.Price(10, 200)
	case store.CtypeNumber:
		return gofakeit.Number(0, 400)
	case store.CtypeLocation:
		return [2]float64{gofakeit.Latitude(), gofakeit.Longitude()}
	case store.CtypeDateTime:
		return gofakeit.Date().UTC()
	case store.CtypeSingleUser, store.CtypeMultiUser:
		opts := stype.Users(c.Slug)
		if len(opts) > 0 {
			return gofakeit.RandomString(opts)
		}
	case store.CtypeEmail:
		return gofakeit.Email()
	case store.CtypeJSON:
		return "{}"
	case store.CtypeRangeNumber:
		// fixme => get ranges from column
		return gofakeit.Price(40, 130)
	case store.CtypeColor:
		return gofakeit.HexColor()
	default:
		//		case store.CtypeRFormula:
		// if !nullables[c.Slug] {
		// 	data[c.Slug] = "1 + 1"
		// }
		return nil
	}

	return nil

}
