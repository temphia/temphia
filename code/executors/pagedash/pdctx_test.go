package pagedash

import (
	"testing"

	"github.com/dop251/goja"
)

func TestCtx(t *testing.T) {

	ctx := PdCtx{
		data: map[string]any{
			"x": 1,
		},
		model: nil,
		ok:    true,
		rt:    goja.New(),
	}

	ctx.rt.Set("test_fail", func(reason string) {
		t.Fatal(reason)
	})

	_, err := ctx.rt.RunScript("test", `
	
	function run_test(version) {
		if (get_data_value("x") !== 1) {
			test_fail("wrong get data value")
		}
	}
	`)

	if err != nil {
		t.Fatal(err)
	}

	ctx.bind()

	err = ctx.execute("run_test", "1")
	if err != nil {
		t.Fatal(err)
	}

}
