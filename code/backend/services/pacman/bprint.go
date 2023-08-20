package pacman

// import (
// 	"github.com/jaevor/go-nanoid"
// 	"github.com/temphia/temphia/code/backend/libx/easyerr"
// 	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
// )

// var bprintIdFunc, _ = nanoid.CustomASCII("abcdefghijklmnopqrstuvwxyz1234567890", 5)

// func (c *PacMan) BprintList(tenantid, group string) ([]*entities.BPrint, error) {
// 	return c.corehub.BprintList(tenantid, group)
// }

// func (c *PacMan) BprintCreate(tenantid string, bp *entities.BPrint) (string, error) {
// 	if bp.ID == "" {
// 		bp.ID = bprintIdFunc()
// 	}

// 	return bp.ID, c.corehub.BprintNew(tenantid, bp)
// }

// func (c *PacMan) BprintUpdate(tenantid, id string, data map[string]any) error {
// 	if id == "" {
// 		return easyerr.NotFound("bprint id/slug")
// 	}

// 	return c.corehub.BprintUpdate(tenantid, id, data)
// }

// func (c *PacMan) BprintGet(tenantid, bid string) (*entities.BPrint, error) {
// 	return c.corehub.BprintGet(tenantid, bid)

// }

// func (c *PacMan) BprintRemove(tenantid, bid string) error {
// 	err := c.corehub.BprintDel(tenantid, bid)
// 	if err != nil {
// 		return err
// 	}

// 	// fixme => del folder

// 	return nil

// }

// func (c *PacMan) BprintListBlobs(tenantid, bid string) (map[string]string, error) {

// 	files, err := c.bstore.ListBlob(tenantid, bid, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := make(map[string]string)
// 	for _, v := range files {
// 		resp[v.Name] = v.LastModified
// 	}

// 	return nil, nil
// }
