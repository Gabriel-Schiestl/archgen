package architectures

import (
	"fmt"
	"os"
	"path/filepath"
)

func Clean(args ...string) {
	dirs := getCleanDirs()

	if err := os.MkdirAll(dirs.mainDir, os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range dirs.children {
		if err := os.MkdirAll(filepath.Join(dirs.mainDir, d.parent), os.ModePerm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, sb := range d.children {
			if err := os.MkdirAll(filepath.Join(dirs.mainDir, d.parent, sb), os.ModePerm); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

type cleanSubDirs struct {
	parent   string
	children []string
}

type cleanDirs struct {
	mainDir  string
	children []cleanSubDirs
}

func getCleanDirs() *cleanDirs {
	domain := cleanSubDirs{
		parent:   "domain",
		children: []string{"model", "exception"},
	}

	application := cleanSubDirs{
		parent:   "application",
		children: []string{"usecase", "dto", "repository", "service"},
	}

	interfaces := cleanSubDirs{
		parent:   "interface",
		children: []string{"repository", "controller", "service"},
	}

	config := cleanSubDirs{
		parent: "config",
		children: []string{},
	}

	cleanDirectories := &cleanDirs{
		mainDir:  "internal",
		children: []cleanSubDirs{domain, application, interfaces, config},
	}

	return cleanDirectories
}