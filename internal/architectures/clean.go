package architectures

func Clean(args ...string) {
	dirs := getCleanDirs()

	createDirs(dirs, args...)
}

type cleanSubDirs struct {
	parent   string
	children []string
}

type cleanDirs struct {
	mainDir  string
	children []*cleanSubDirs
}

func getCleanDirs() *cleanDirs {
	domain := &cleanSubDirs{
		parent:   "domain",
		children: []string{"model", "exception"},
	}

	application := &cleanSubDirs{
		parent:   "application",
		children: []string{"usecase", "dto", "repository", "service"},
	}

	interfaces := &cleanSubDirs{
		parent:   "interface",
		children: []string{"repository", "controller", "service"},
	}

	config := &cleanSubDirs{
		parent:   "config",
		children: []string{},
	}

	cleanDirectories := &cleanDirs{
		mainDir:  "internal",
		children: []*cleanSubDirs{domain, application, interfaces, config},
	}

	return cleanDirectories
}