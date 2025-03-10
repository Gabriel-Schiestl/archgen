package architectures

import (
	"log"
	"os"
	"path/filepath"
)

func createDirs(dirs *cleanDirs, dir, lang string) {
	var baseDir string

	if dir != "" {
		baseDir = dir
	} else {
		baseDir = ""
	}

	mainDirPath := filepath.Join(baseDir, dirs.mainDir)

	if err := os.MkdirAll(mainDirPath, os.ModePerm); err != nil {
		log.Fatal(err.Error())
	}

	for _, d := range dirs.children {
		actualDir := filepath.Join(mainDirPath, d.parent)

		if err := os.MkdirAll(actualDir, os.ModePerm); err != nil {
			log.Fatal(err.Error())
		}

		if lang == "python" {
			createFileIfNotExists(filepath.Join(actualDir, "__init__.py"))
		}

		for _, sb := range d.children {
			if err := os.MkdirAll(filepath.Join(actualDir, sb), os.ModePerm); err != nil {
				log.Fatal(err.Error())
			}

			if lang == "python" {
				createFileIfNotExists(filepath.Join(actualDir, sb, "__init__.py"))
			}
		}
	}
}

func createFileIfNotExists(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0644)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err.Error())
	}

	file.Close()
}