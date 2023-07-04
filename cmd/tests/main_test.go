package tests

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
)

func TestMain(m *testing.M) {
	if os.Args[len(os.Args)-1] == "list_tests" {
		pp.Println(m)
		return
	}

	pp.Println("Test setup")

	exitCode := m.Run()

	pp.Println("Test teardown")

	os.Exit(exitCode)
}

func TestEntry(t *testing.T) {

	pp.Println("@todo")

}
