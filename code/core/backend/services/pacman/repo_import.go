package pacman

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"

	"github.com/thoas/go-funk"
)

func (p *PacMan) RepoSourceImport(tenantid string, opts *service.RepoImportOpts) (string, error) {

	repo := p.getRepoSource(tenantid, opts.Source)
	if repo == nil {
		return "", easyerr.NotFound()
	}

	bp, err := repo.GetItem(tenantid, opts.Group, opts.Slug)
	if err != nil {
		return "", err
	}

	bp.ID = opts.NewId
	if bp.ID == "" {
		bp.ID = xid.New().String()
	}

	bp.TenantID = tenantid

	if opts.SkipFiles != nil || len(opts.SkipFiles) != 0 {
		files := make([]string, 0, len(bp.Files)-len(opts.SkipFiles))
		for _, f := range bp.Files {
			if !funk.ContainsString(opts.SkipFiles, f) {
				files = append(files, f)
			}
		}
		bp.Files = entities.JsonArray(files)
	}

	backFiles := bp.Files
	bp.Files = make(entities.JsonArray, 0)

	_, err = p.BprintCreate(tenantid, bp)
	if err != nil {
		return "", err
	}

	bstore := p.blobStore(tenantid)

	for _, fkey := range backFiles {

		if funk.ContainsString(opts.SkipFiles, fkey) {
			continue
		}

		bytes, err := repo.GetFile(tenantid, opts.Group, opts.Slug, fkey)
		if err != nil {
			return "", err
		}
		err = bstore.AddBlob(context.TODO(), xtypes.BprintBlobFolder, ffile(bp.ID, fkey), bytes)
		if err != nil {
			return "", err
		}

	}

	out, err := json.Marshal(backFiles)
	if err != nil {
		return "", nil
	}

	err = p.syncer.BprintUpdate(tenantid, bp.ID, map[string]any{
		"files": string(out),
	})

	if err != nil {
		return "", err
	}

	return bp.ID, nil
}
