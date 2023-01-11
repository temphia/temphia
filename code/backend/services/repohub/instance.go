package repohub

// func (p *PacMan) Instance(tenantId string, opts *repox.InstanceOptions) (any, error) {

// 	bprint, err := p.corehub.BprintGet(tenantId, opts.BprintId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	instanceType := bprint.Type

// 	if bprint.Type == xbprint.TypeBundle || bprint.Type == "bundle" {
// 		instanceType = opts.InstancerType
// 	}

// 	is, ok := p.instancers[instanceType]
// 	if !ok {
// 		return nil, easyerr.NotFound()
// 	}

// 	return is.Instance(xinstance.Options{
// 		TenantId:     opts.UserSession.TenentId,
// 		BprintId:     opts.BprintId,
// 		InstanceType: opts.InstancerType,
// 		File:         opts.File,
// 		UserId:       opts.UserSession.UserID,
// 		UserData:     opts.UserConfigData,
// 	})
// }

// // private

// func (p *PacMan) ParseInstanceFile(tenantId, bid, file string, target any) error {
// 	return p.readInstanceFile(tenantId, bid, file, target)
// }

// func (p *PacMan) readInstanceFile(tenantId, bprint, file string, target any) error {

// 	out, err := p.BprintGetBlob(tenantId, bprint, file)
// 	if err != nil {
// 		return err
// 	}

// 	if strings.HasSuffix(file, ".json") {
// 		return json.Unmarshal(out, target)
// 	} else if strings.HasSuffix(file, ".yaml") {
// 		return yaml.Unmarshal(out, target)
// 	} else {
// 		panic("")
// 	}
// }
