package app

import (
	"os"
	"os/signal"
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/postfinance/single"
	"github.com/temphia/temphia/code/backend/xtypes"
)

var _ xtypes.App = (*App)(nil)

type App struct {
	nodeId    string
	clusterId string
	tenantId  string
	devmode   bool
	deps      AppDeps
	global    Global

	closer sync.Once
	single *single.Single
}

func (a *App) Run() error { return a.run() }

func (a *App) NodeId() string                 { return a.nodeId }
func (a *App) ClusterId() string              { return a.clusterId }
func (a *App) DevMode() bool                  { return a.devmode }
func (a *App) TenantId() string               { return a.tenantId }
func (a *App) GetDeps() xtypes.Deps           { return &a.deps }
func (a *App) GetServer() xtypes.Server       { return nil }
func (a *App) GetGlobalVar() xtypes.GlobalVar { return &a.global }

func (a *App) run() error {

	single, err := single.New("instance", single.WithLockPath(a.deps.confd.RootDataFolder()))
	if err != nil {
		return err
	}

	err = single.Lock()
	if err != nil {
		return err
	}

	a.single = single

	err = a.deps.start()
	if err != nil {
		return err
	}

	defer a.close()

	return a.deps.server.Listen()
}

func (a *App) close() error {
	var err error

	a.closer.Do(func() {
		err = a.deps.server.Close()
	})

	for _, xt := range a.deps.extensions {
		multierror.Append(err, xt.Close())
	}

	if a.single != nil {
		err = multierror.Append(err, a.single.Unlock())
	}

	return err
}

func (a *App) SetupSignalHandler() {
	SetupSignalHandler(func(signal os.Signal) int {
		// fixme => only only exist on -9/15 ??

		err := a.close()
		if err != nil {
			return 1
		}

		return 0
	})
}

func SetupSignalHandler(fn func(signal os.Signal) int) {
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	go func() {
		for {
			s := <-sigchnl
			// fixme => only only exist on -9/15 ??
			exitchnl <- fn(s)
		}
	}()

	exitcode := <-exitchnl
	os.Exit(exitcode)
}
