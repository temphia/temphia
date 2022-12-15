package seeder

import (
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type SeederCore struct {
	Group            string
	MainTable        *entities.Table
	MainTableColumns map[string]*entities.Table
	SiblingTables    map[string]*entities.Table
	SiblingColumns   map[string][]*entities.Column
	UserId           string
}

func (s *SeederCore) files(column string) []string {
	return []string{}
}

func (s *SeederCore) users(column string) []string {
	return []string{}
}

func (s *SeederCore) generateTableSeed(no int, cols []*entities.Column, nullables map[string]bool) []map[string]any {

	datas := make([]map[string]any, 0, no)

	for i := 0; i <= no; i = i + 1 {
		data := make(map[string]any)
		data[store.KeyPrimary] = i + 1

	columnloop:
		for _, c := range cols {

			if nullables[c.Slug] {
				if rand.Int()%3 == 1 {
					continue
				}
			}

			if c.RefType != "" {
				switch c.RefType {
				case store.RefHardPriId, store.RefSoftPriId:
					data[c.Slug] = gofakeit.Number(1, no)
					continue columnloop
				case store.RefHardText:
				case store.RefSoftText:
				case store.RefHardMulti:
				default:
				}

			}

			switch c.Ctype {
			case store.CtypeShortText:

				switch c.Slug {
				case "name":
					data[c.Slug] = gofakeit.Name()
				case "addr":
					data[c.Slug] = gofakeit.Address().Address
				default:
					data[c.Slug] = gofakeit.HipsterWord()
				}

			case store.CtypeLongText:
				data[c.Slug] = gofakeit.HipsterSentence(20)
			case store.CtypePhone:
				data[c.Slug] = gofakeit.Phone()
			case store.CtypeSelect, store.CtypeMultSelect:
				if c.Options != nil {
					data[c.Slug] = gofakeit.RandomString(c.Options)
				}
			case store.CtypeRFormula:
				if !nullables[c.Slug] {
					data[c.Slug] = "1 + 1"
				}
			case store.CtypeFile, store.CtypeMultiFile:
				data[c.Slug] = gofakeit.RandomString(s.files(c.Slug))
			case store.CtypeCheckBox:
				data[c.Slug] = gofakeit.Bool()
			case store.CtypeCurrency:
				data[c.Slug] = gofakeit.Price(10, 200)
			case store.CtypeNumber:

				data[c.Slug] = gofakeit.Number(0, 400)
			case store.CtypeLocation:
				data[c.Slug] = [2]float64{gofakeit.Latitude(), gofakeit.Longitude()}
			case store.CtypeDateTime:
				data[c.Slug] = gofakeit.Date().UTC()
			case store.CtypeSingleUser, store.CtypeMultiUser:
				data[c.Slug] = gofakeit.RandomString(s.users(c.Slug))
			case store.CtypeEmail:
				data[c.Slug] = gofakeit.Email()
			case store.CtypeJSON:
				data[c.Slug] = "{}"
			case store.CtypeRangeNumber:
				data[c.Slug] = gofakeit.Price(40, 130)
			case store.CtypeColor:
				data[c.Slug] = gofakeit.HexColor()
			default:
				fmt.Println("skipping ", c)
			}

		}

		datas = append(datas, data)
	}

	return datas
}
