package helper

import (
	"fmt"
	"os"
)

type ExecutorHelper struct {
	dev          bool
	name         string
	staticAssets map[string][]byte
}

func New(name string, dev bool) *ExecutorHelper {
	return &ExecutorHelper{
		dev:          dev,
		name:         name,
		staticAssets: make(map[string][]byte),
	}
}

func (eh *ExecutorHelper) SetAsset(name string, data []byte) {
	eh.staticAssets[name] = data
}

func (eh *ExecutorHelper) Serve(file string) ([]byte, error) {
	if eh.dev {
		switch file {
		case "loader.css":
			return getData(eh.name + ".css")
		case "loader.js":
			return getData(eh.name + ".js")
		default:
			modFile := fmt.Sprintf("%s.js.map", eh.name)
			if modFile == file {
				return getData(modFile)
			}
		}
	}
	return eh.staticAssets[file], nil
}

func getData(file string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("code/frontend/public/build/%s", file))
}
