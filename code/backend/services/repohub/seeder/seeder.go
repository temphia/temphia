package seeder

import (
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/vmodels"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store"

	"github.com/goccy/go-yaml"
)

type Seeder struct {
	userId           string
	tg               *xbprint.NewTableGroup
	model            *entities.BPrint
	pacman           repox.Hub
	source           store.DynSource
	tenant           string
	group            string
	selectableImages []string
	selectableUsers  []string
}

func New(schema *xbprint.NewTableGroup, pman repox.Hub, dsource store.DynSource, tenantId, dataGroup, userId string) *Seeder {
	return &Seeder{
		tg:               schema,
		model:            nil,
		pacman:           pman,
		source:           dsource,
		tenant:           tenantId,
		group:            dataGroup,
		selectableImages: []string{},
		selectableUsers:  []string{},
		userId:           userId,
	}

}

func (s *Seeder) getTable(name string) *xbprint.NewTable {
	for _, tbl := range s.tg.Tables {
		if tbl.Slug == name {
			return tbl
		}
	}
	return nil
}

func (s *Seeder) DataSeed() error {

	bytes, err := s.pacman.BprintGetBlob(s.tenant, s.model.ID, "data.yaml")
	if err != nil {
		return err
	}
	data := &vmodels.DynData{}
	err = yaml.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	return s.applySeed(data.Data)
}

//

func (s *Seeder) applySeed(data map[string][]map[string]any) error {

	pp.Println("applying seed")

	doneRows := make(map[string]map[int64]int64)

	for _, table := range s.tg.ExecOrder {
		doneRows[table] = make(map[int64]int64)
	}

	for _, table := range s.tg.ExecOrder {
		tdata, ok := data[table]
		if !ok {
			pp.Println("skipping ", table)
			continue
		}

		tableDoneRows := doneRows[table]

		tmodel := s.getTable(table)

		pp.Println("@processing table", table)

		for _, rdata := range tdata {

			pp.Println("@processing row table")
			fmt.Println(rdata)

			for _, ref := range tmodel.ColumnRef {
				if ref.Type == store.RefHardPriId || ref.Type == store.RefSoftPriId {
					targetDoneTbl := doneRows[ref.Target]

					pp.Println("REF =>>>>", ref)

					_oldRowAlias, ok := rdata[ref.FromCols[0]]
					if !ok {
						continue
					}
					oldRowAlias, ok := _oldRowAlias.(int64)
					if !ok {
						continue
					}

					rdata[ref.FromCols[0]] = targetDoneTbl[oldRowAlias]
				}
			}

			_possibleId := rdata[store.KeyPrimary]

			delete(rdata, store.KeyPrimary)

			_rid, err := s.source.NewRow(0, store.NewRowReq{
				TenantId: s.tenant,
				Group:    s.group,
				Table:    table,
				Data:     rdata,
				ModCtx: store.ModCtx{
					UserId: s.userId,
				},
			})
			if err != nil {
				pp.Println("SEED ERROR =>", err)
				continue
			}
			if _possibleId == nil {
				continue
			}

			aliasId, ok := _possibleId.(uint64)
			if !ok {
				continue
			}

			tableDoneRows[int64(aliasId)] = _rid
		}

	}

	return nil

}

func (s *Seeder) GeneratedSeed(no int) error {

	pp.Println("Generating seed")

	data := make(map[string][]map[string]any)

	for _, etbl := range s.tg.ExecOrder {
		pp.Println("generating for table", etbl)

		tbl := s.getTable(etbl)
		nullables := make(map[string]bool)

		cols := store.ExtractColumns(tbl, s.tenant, s.group)

		for _, col := range tbl.Columns {
			nullables[col.Slug] = !col.NotNullable
		}

		data[etbl] = s.generateTableSeed(no, cols, nullables)
	}

	return s.applySeed(data)
}

func (s *Seeder) generateTableSeed(no int, cols []*entities.Column, nullables map[string]bool) []map[string]any {

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
				data[c.Slug] = gofakeit.RandomString(s.selectableImages)
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
				data[c.Slug] = gofakeit.RandomString(s.selectableUsers)
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

	pp.Println("@=>>", datas)

	return datas
}
