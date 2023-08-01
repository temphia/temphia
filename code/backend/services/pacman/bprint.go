package pacman

import (
	"context"
	"encoding/json"

	"github.com/jaevor/go-nanoid"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"

	"github.com/thoas/go-funk"
)

var bprintIdFunc, _ = nanoid.CustomASCII("abcdefghijklmnopqrstuvwxyz1234567890", 5)

func (c *PacMan) BprintList(tenantid, group string) ([]*entities.BPrint, error) {
	return c.corehub.BprintList(tenantid, group)
}

func (c *PacMan) BprintCreate(tenantid string, bp *entities.BPrint) (string, error) {
	if bp.ID == "" {
		bp.ID = bprintIdFunc()
	}

	return bp.ID, c.corehub.BprintNew(tenantid, bp)
}

func (c *PacMan) BprintUpdate(tenantid string, bp *entities.BPrint) error {
	if bp.ID == "" || bp.Slug == "" {
		return easyerr.NotFound("bprint id/slug")
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
		_ = c.cabinet.DeleteBlob(context.TODO(), tenantid, xtypes.BprintBlobFolder, ffile(bid, file))
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

	err := c.cabinet.AddBlob(context.TODO(), tenantid, xtypes.BprintBlobFolder, ffile(bid, file), payload)
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
	return c.cabinet.AddBlob(context.TODO(), tenantid, xtypes.BprintBlobFolder, ffile(bid, file), payload)
}

func (c *PacMan) BprintGetBlob(tenantid, bid, file string) ([]byte, error) {
	return c.cabinet.GetBlob(context.TODO(), tenantid, xtypes.BprintBlobFolder, ffile(bid, file))
}
func (c *PacMan) BprintDeleteBlob(tenantid, bid, file string) error {
	// fixme => also delete from list in  bprint
	return c.cabinet.DeleteBlob(context.TODO(), tenantid, ffile(bid, file), file)
}

func ffile(id, file string) string {
	return id + "_" + file
}