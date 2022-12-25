package repobuild

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
)

func TestM(t *testing.T) {

	fout, err := os.ReadFile("testdata/repo.json")
	if err != nil {
		t.Fatal(err)
	}

	tb, err := New(fout)
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(tb.BuildOne("baghchal", false))

}
