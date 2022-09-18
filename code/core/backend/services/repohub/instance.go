package repohub

import (
	"encoding/json"
	"strings"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints/instancer"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/instance"
	"gopkg.in/yaml.v2"
)

func (p *PacMan) Instance(tenantId string, opts *instance.RepoOptions) (any, error) {

	bprint, err := p.corehub.BprintGet(tenantId, opts.BprintId)
	if err != nil {
		return nil, err
	}

	instanceType := bprint.Type

	if bprint.Type == bprints.TypeAppBundle || bprint.Type == "bundle" {
		instanceType = opts.InstancerType
	}

	is, ok := p.instancers[instanceType]
	if !ok {
		return nil, easyerr.NotFound()
	}

	return is.Instance(instancer.Options{
		TenantId:     tenantId,
		Bid:          opts.BprintId,
		InstanceType: instanceType,
		File:         opts.File,
		UserId:       opts.UserId,
		Data:         opts.Data,
	})

}

// private

func (p *PacMan) ParseInstanceFile(tenantId, bid, file string, target any) error {
	return p.readInstanceFile(tenantId, bid, file, target)
}

func (p *PacMan) readInstanceFile(tenantId, bprint, file string, target any) error {

	out, err := p.BprintGetBlob(tenantId, bprint, file)
	if err != nil {
		return err
	}

	if strings.HasSuffix(file, ".json") {
		return json.Unmarshal(out, target)
	} else if strings.HasSuffix(file, ".yaml") {
		return yaml.Unmarshal(out, target)
	} else {
		panic("")
	}
}
