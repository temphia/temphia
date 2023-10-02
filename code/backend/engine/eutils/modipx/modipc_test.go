package modipx

import (
	"testing"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
)

type TestMod struct {
	Count int
}

type pingOpts struct {
	Add int
}

func (t *TestMod) Ping(opts *pingOpts) (any, error) {

	return t.Count + opts.Add, nil
}

func TestModIPC(t *testing.T) {

	tmod := &TestMod{
		Count: 42,
	}

	modipc := NewModIPX(tmod)

	resp, err := modipc.Handle("Ping", lazydata.NewJsonData([]byte(`{"Add": 10}`)))
	if err != nil {
		t.Fatal(err)
	}

	bout, err := resp.AsJsonBytes()
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(string(bout))

	if string(bout) != "52" {

		t.Fatal("resp not right")
	}

}
