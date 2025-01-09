package guides

import (
	"os"
	"path/filepath"
)

type Guide string

func (g Guide) GetFileName() string {
	return filepath.Base(string(g))
}

func (g Guide) GetFullPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return filepath.Join(exPath, string(g))
}

func (g Guide) CheckIfExists() bool {
	fullPath := g.GetFullPath()

	_, err := os.Stat(fullPath)
	return err != nil
}

type ChangeViewingGuideMsg string
