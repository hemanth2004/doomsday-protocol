package guides

import (
	"errors"
	"os"
	"path/filepath"
)

type Guide string

func GuidesFolderFromWorkingDirectory() string {
	return filepath.Join(getWorkingDirectory(), "packaged", "Guides")
}

func (g Guide) GetFileName() string {
	return filepath.Base(string(g))
}

func (g Guide) GetFullPath() string {
	exPath := getWorkingDirectory()

	return filepath.Join(exPath, string(g))
}

func (g Guide) CheckIfExists() bool {
	fullPath := g.GetFullPath()

	_, err := os.Stat(fullPath)
	return err != nil
}

type ChangeViewingGuideMsg string

func getWorkingDirectory() string {
	mydir, err := os.Getwd()
	if err == nil {
		return mydir
	} else {
		panic(errors.New("failed to get working directory"))
	}

}
