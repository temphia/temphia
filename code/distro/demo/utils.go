package demo

import (
	"net"
	"os"
	"os/signal"
	"strconv"
)

func getPort() (int, error) {
	port := os.Getenv("TEMPHIA_DEMO_PG_PORT")
	if port == "" {
		// fixme check postgres file
		return getFreePort()
	}

	pport, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(pport), nil
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func setUpHandler(fn func(signal os.Signal)) {
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	go func() {
		for {
			s := <-sigchnl
			fn(s)
		}
	}()

	exitcode := <-exitchnl
	os.Exit(exitcode)
}
