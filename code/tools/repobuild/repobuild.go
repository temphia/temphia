package repobuild

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/temphia/temphia/code/core/backend/libx/xutils"
)

// RepoBuild is simple helper for building repo by calling underlying build system.
// underlying build system should generate `index.json` (which is like manifest file)
// and other build artifacts
type RepoBuild struct {
	config *BuildConfig
}

func New(conf []byte) (*RepoBuild, error) {
	bconf := &BuildConfig{}

	err := json.Unmarshal(conf, bconf)
	if err != nil {
		return nil, err
	}

	return &RepoBuild{
		config: bconf,
	}, nil

}

func (rb *RepoBuild) BuildAll() (*BuildResult, error) {

	result := &BuildResult{
		ErroredItems: make(map[string]error),
		Outputs:      make(map[string]string),
	}

	for k := range rb.config.Items {

		ofolder, err := rb.BuildOne(k, false)
		if err != nil {
			result.ErroredItems[k] = err
			continue
		}
		result.Outputs[k] = ofolder
	}

	return result, nil
}

func (rb *RepoBuild) BuildOne(name string, zip bool) (string, error) {
	of, err := rb.buildItem(name)
	if err != nil {
		return "", err
	}

	if !zip {
		return of, nil
	}

	panic("Zip not implemented")
}

func (rb *RepoBuild) buildItem(name string) (string, error) {
	item := rb.config.Items[name]

	buildPath := path.Join(rb.config.BuildFolder, name)
	outputPath := path.Join(rb.config.OutputFolder, name, item.OutputFolder)

	err := xutils.CreateIfNotExits(buildPath)
	if err != nil {
		return "", err
	}

	_, err = git.PlainClone(buildPath, false, &git.CloneOptions{
		URL:           item.GitURL,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(item.Branch),
		SingleBranch:  true,
		Depth:         1,
	})

	if err != nil {
		if !errors.Is(git.ErrRepositoryAlreadyExists, err) {
			return "", err
		}
	}

	cmd := exec.Command(item.BuildCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Path = buildPath

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, CopyDirectory(buildPath, outputPath)
}
