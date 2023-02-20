package repohub

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"

	"github.com/thoas/go-funk"
)

func (c *PacMan) BprintList(tenantid, group string) ([]*entities.BPrint, error) {
	return c.corehub.BprintList(tenantid, group)
}

func (c *PacMan) BprintCreate(tenantid string, bp *entities.BPrint) (string, error) {
	if bp.ID == "" {
		bp.ID = xid.New().String()
	}

	return bp.ID, c.corehub.BprintNew(tenantid, bp)
}

func (c *PacMan) BprintUpdate(tenantid string, bp *entities.BPrint) error {
	if bp.ID == "" || bp.Slug == "" {
		return easyerr.NotFound()
	}

	//_, err := c.localStore.BprintUpdate(tenantid, bp)

	return easyerr.NotImpl()
}

func (c *PacMan) BprintGet(tenantid, bid string) (*entities.BPrint, error) {
	return c.corehub.BprintGet(tenantid, bid)

}

func (c *PacMan) BprintRemove(tenantid, bid string) error {
	bprint, err := c.corehub.BprintGet(tenantid, bid)
	if err != nil {
		return err
	}
	for _, file := range bprint.Files {
		_ = c.blobStore(tenantid).DeleteBlob(context.TODO(), xtypes.BprintBlobFolder, ffile(bid, file))
	}
	return c.corehub.BprintDel(tenantid, bid)
}

func (c *PacMan) BprintListBlobs(tenantid, bid string) (map[string]string, error) {
	bprint, err := c.corehub.BprintGet(tenantid, bid)
	if err != nil {
		return nil, err
	}

	resp := make(map[string]string)
	for _, v := range bprint.Files {
		// fixme => <file_name><mod_data>
		resp[v] = ""
	}

	return resp, nil
}

func (c *PacMan) BprintNewBlob(tenantid, bid, file string, payload []byte, updateList bool) error {
	err := c.blobStore(tenantid).AddBlob(context.TODO(), xtypes.BprintBlobFolder, ffile(bid, file), payload)
	if err != nil {
		return err
	}

	if !updateList {
		return nil
	}

	return c.BprintUpdateFilesList(tenantid, bid, file)
}

func (c *PacMan) BprintUpdateFilesList(tenantid, bid string, files ...string) error {

	bprint, err := c.BprintGet(tenantid, bid)
	if err != nil {
		return err
	}

	if bprint.Files == nil {
		bprint.Files = entities.JsonArray{}
	}

	for _, v := range files {
		if funk.ContainsString(bprint.Files, v) {
			continue
		}
		bprint.Files = append(bprint.Files, v)
	}

	bfiles, err := json.Marshal(bprint.Files)
	if err != nil {
		return err
	}

	return c.corehub.BprintUpdate(tenantid, bid, map[string]any{
		"files": bfiles,
	})
}

func (c *PacMan) BprintUpdateBlob(tenantid, bid, file string, payload []byte) error {
	return c.blobStore(tenantid).AddBlob(context.TODO(), xtypes.BprintBlobFolder, ffile(bid, file), payload)
}

func (c *PacMan) BprintGetBlob(tenantid, bid, file string) ([]byte, error) {
	return c.blobStore(tenantid).GetBlob(context.TODO(), xtypes.BprintBlobFolder, ffile(bid, file))
}
func (c *PacMan) BprintDeleteBlob(tenantid, bid, file string) error {
	// fixme => also delete from list in  bprint
	return c.blobStore(tenantid).DeleteBlob(context.TODO(), ffile(bid, file), file)
}

func ffile(id, file string) string {
	return id + "_" + file
}
