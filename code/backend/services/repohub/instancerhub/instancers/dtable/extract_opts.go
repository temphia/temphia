package dtable

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (di *dtabeInstancer) extractUserOptions(tenantId string, auto bool, userData []byte, schemaData *xbprint.NewTableGroup) (*DataGroupRequest, error) {
	return ExtractUserOptions(di.cabhub, di.coreHub, di.dynhub)(tenantId, auto, userData, schemaData)
}

func ExtractUserOptions(cabhub store.CabinetHub, coreHub store.CoreHub, dynhub dyndb.DataHub) func(tenantId string, auto bool, userData []byte, schemaData *xbprint.NewTableGroup) (*DataGroupRequest, error) {

	return func(tenantId string, auto bool, userData []byte, schemaData *xbprint.NewTableGroup) (*DataGroupRequest, error) {

		dopts := &DataGroupRequest{}
		if !auto {
			err := json.Unmarshal(userData, dopts)
			if err != nil {
				return nil, err
			}
		}

		grandom, _ := xutils.GenerateRandomString(5)

		dopts.DyndbSource = "default"

		csource := cabhub.Default(tenantId)
		dopts.CabinetSource = csource.Name()
		dopts.CabinetFolder = store.DefaultDataAssetsFolder

		dopts.GroupName = schemaData.Name
		dopts.GroupSlug = schemaData.Slug + grandom
		dopts.SeedType = dyndb.DynSeedTypeAutogen

		tblOpts := make(map[string]*DataTableOption, len(schemaData.Tables))
		for _, nt := range schemaData.Tables {
			tblOpts[nt.Slug] = &DataTableOption{
				Name:         nt.Name,
				Slug:         nt.Slug,
				ActivityType: dyndb.DynActivityTypeStrict,
				SyncType:     dyndb.DynSyncTypeEventAndData,
				Seed:         true,
			}
		}

		dopts.TableOptions = tblOpts

		return dopts, nil

	}

}
