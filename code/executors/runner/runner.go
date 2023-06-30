package runner

import (
	"os"
	"os/exec"
	"path"
	"sync"
)

type Options struct {
	BootstrapFile   string
	ExecutorLibData string
	ExecutorLibName string
	BootstrapFunc   func(ctx BootstrapContext) error
	RootFilesFunc   func() ([]byte, error)
}

type Runner struct {
	opts  *Options
	token string

	blines     map[string]*bindingsLine
	blinesLock sync.Mutex

	controlLine *controlLine
	clineLock   sync.Mutex
}

func New(opts *Options) *Runner {
	r := &Runner{
		opts:        opts,
		token:       "",
		controlLine: nil,
		clineLock:   sync.Mutex{},
		blines:      make(map[string]*bindingsLine),
		blinesLock:  sync.Mutex{},
	}

	return r
}

func (r *Runner) init() error {

	err := r.startServer()
	if err != nil {
		return err
	}

	// ExRunner

	tdir, err := os.MkdirTemp("exrunner", "*")
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(tdir, "bootstrap.sh"), []byte(r.opts.BootstrapFile), 0777)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("bootstrap.sh")
	cmd.Dir = tdir

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
