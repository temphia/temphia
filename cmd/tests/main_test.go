package tests

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/executors/re/retest"
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

func TestDummyTest(t *testing.T) {

}

func TestRemoteExecutor(t *testing.T) {
	WithPath("code/executors/re", t)(retest.TestRe)
}

func WithPath(path string, t *testing.T) func(func(t *testing.T)) {
	return func(f func(t *testing.T)) {

		dir, err := os.Getwd()
		if err != nil {
			t.Fatal("COULD NOT GET WORKING DIR PATH", err)
		}

		os.Chdir(path)
		f(t)
		os.Chdir(dir)

	}
}
