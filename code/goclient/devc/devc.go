package devc

import (
	"context"
	"fmt"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/goclient/devapi"
	"github.com/tidwall/pretty"
)

type DevClient struct {
	api *devapi.DevAPI
}

func New(url, token string) *DevClient {

	return &DevClient{
		api: devapi.New(url, token),
	}

}

func (dc *DevClient) Watch(plug string, agent string) {

	conn, _, _, err := ws.Dial(context.TODO(), dc.api.ExecWatchURL(plug, agent))
	if err != nil {
		pp.Println("@err", err)
		return
	}

	for {

		out, _, err := wsutil.ReadServerData(conn)
		if err != nil {
			pp.Println(err)
			return
		}

		fmt.Print(string(pretty.Pretty(out)))
	}

}

func (dc *DevClient) PushFile(name, path string) error {
	out, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return dc.api.BprintFilePush(name, out)
}

func (dc *DevClient) Reset(plug, agent string) error {
	return dc.api.ExecReset(plug, agent)
}

func (dc *DevClient) ExecRun(pid, aid, action string, payload any) ([]byte, error) {
	return dc.api.ExecRun(pid, aid, action, payload)
}
