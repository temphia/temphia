package seeder

import (
	"math/rand"

	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type Options struct {
	TenantId        string
	UserId          string
	GroupId         string
	Dydb            dyndb.DynDB
	SelectableFiles []string
	SelectableUsers []string
}

type AutoSeeder struct {
	opts        Options
	recordCache map[string]map[int64]map[string]any
}

func NewAutoSeeder(opts Options) *AutoSeeder {
	return &AutoSeeder{
		opts:        opts,
		recordCache: make(map[string]map[int64]map[string]any),
	}
}

func (a *AutoSeeder) TableGenerate(table string, no int) ([]map[string]any, error) {
	return a.generateTableSeed(table, no)
}

func (a *AutoSeeder) generateTableSeed(table string, no int) ([]map[string]any, error) {

	cache := a.opts.Dydb.GetCache()
	cols, err := cache.CachedColumns(a.opts.TenantId, a.opts.GroupId, table)
	if err != nil {
		return nil, err
	}

	datas := make([]map[string]any, 0, no)

	for i := 0; i <= no; i = i + 1 {
		data := make(map[string]any)
		//	data[dyndb.KeyPrimary] = i + 1

	columnloop:
		for _, c := range cols {

			if !c.NotNullable {
				if rand.Int()%3 == 1 {
					continue
				}

			}

			if c.RefType != "" {
				// fixme => proper ref type
				switch c.RefType {
				case dyndb.RefHardPriId, dyndb.RefSoftPriId:
					data[c.Slug] = gofakeit.Number(1, no)
					continue columnloop
				case dyndb.RefHardText:
				case dyndb.RefSoftText:
				case dyndb.RefHardMulti:
				default:
				}

			}

			switch c.Ctype {
			case dyndb.CtypeShortText:

				switch c.Slug {
				case "name":
					data[c.Slug] = gofakeit.Name()
				case "addr":
					data[c.Slug] = gofakeit.Address().Address
				default:
					data[c.Slug] = gofakeit.HipsterWord()
				}

			case dyndb.CtypeLongText:
				data[c.Slug] = gofakeit.HipsterSentence(20)
			case dyndb.CtypePhone:
				data[c.Slug] = gofakeit.Phone()
			case dyndb.CtypeSelect, dyndb.CtypeMultSelect:
				if c.Options != nil {
					data[c.Slug] = gofakeit.RandomString(c.Options)
				}
			case dyndb.CtypeRFormula:
				if c.NotNullable {
					data[c.Slug] = "1 + 1"
				}
			case dyndb.CtypeFile, dyndb.CtypeMultiFile:
				data[c.Slug] = gofakeit.RandomString(a.opts.SelectableFiles)
			case dyndb.CtypeCheckBox:
				data[c.Slug] = gofakeit.Bool()
			case dyndb.CtypeCurrency:
				data[c.Slug] = gofakeit.Price(10, 200)
			case dyndb.CtypeNumber:
				data[c.Slug] = gofakeit.Number(0, 400)
			case dyndb.CtypeLocation:
				data[c.Slug] = [2]float64{gofakeit.Latitude(), gofakeit.Longitude()}
			case dyndb.CtypeDateTime:
				data[c.Slug] = gofakeit.Date().UTC()
			case dyndb.CtypeSingleUser, dyndb.CtypeMultiUser:
				data[c.Slug] = gofakeit.RandomString(a.opts.SelectableUsers)
			case dyndb.CtypeEmail:
				data[c.Slug] = gofakeit.Email()
			case dyndb.CtypeJSON:
				data[c.Slug] = "{}"
			case dyndb.CtypeRangeNumber:
				data[c.Slug] = gofakeit.Price(40, 130)
			case dyndb.CtypeColor:
				data[c.Slug] = gofakeit.HexColor()
			default:
				fmt.Println("skipping ", c)
			}

		}

		datas = append(datas, data)
	}

	return datas, nil
}
