package architectures

import (
	"log"
	"os"
	"path/filepath"
)

func createDirs(dirs *cleanDirs, args ...string) {
	var baseDir string

	if len(args) > 0 && args[0] != "" {
		baseDir = args[0]
	} else {
		baseDir = ""
	}

	mainDirPath := filepath.Join(baseDir, dirs.mainDir)

	if err := os.MkdirAll(mainDirPath, os.ModePerm); err != nil {
		log.Fatal(err.Error())
	}

	for _, d := range dirs.children {
		if err := os.MkdirAll(filepath.Join(mainDirPath, d.parent), os.ModePerm); err != nil {
			log.Fatal(err.Error())
		}

		for _, sb := range d.children {
			if err := os.MkdirAll(filepath.Join(mainDirPath, d.parent, sb), os.ModePerm); err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}