package dev

import (
	"fmt"
	"io"
	"os"
	"path"
)

func copyFilesFromBuildProd() error {
	prod := "../code/frontend/ui/build_prod/"
	target := "../code/frontend/ui/build_dev/"

	files, err := os.ReadDir(prod)
	if err != nil {
		return fmt.Errorf("error reading source directory: %w", err)
	}

	for _, file := range files {
		fmt.Println("|>", file.Name())

		if file.IsDir() || file.Name() == ".gitkeep" {
			continue
		}

		sourceFilePath := path.Join(prod, file.Name())
		targetFilePath := path.Join(target, file.Name())

		pfile, err := os.Open(sourceFilePath)
		if err != nil {
			return fmt.Errorf("error opening source file %s: %w", sourceFilePath, err)
		}
		defer pfile.Close()

		_, err = os.Stat(targetFilePath)
		if err == nil {
			fmt.Println("Already contains:", file.Name())
			continue
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("error checking target file %s: %w", targetFilePath, err)
		}

		tfile, err := os.Create(targetFilePath)
		if err != nil {
			return fmt.Errorf("error creating target file %s: %w", targetFilePath, err)
		}
		defer tfile.Close()

		_, err = io.Copy(tfile, pfile)
		if err != nil {
			return fmt.Errorf("error copying data to target file %s: %w", targetFilePath, err)
		}
	}

	return nil
}
