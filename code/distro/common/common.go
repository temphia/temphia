package common

import (
	"errors"

	"github.com/postfinance/single"
)

func IsRunning(folder string) bool {

	s, err := single.New("instance", single.WithLockPath(folder))
	if err != nil {
		panic(err)
	}

	err = s.Lock()
	if err != nil {
		if errors.Is(err, single.ErrAlreadyRunning) {
			return true
		}

		panic(err)
	}

	defer s.Unlock()

	return false

}
