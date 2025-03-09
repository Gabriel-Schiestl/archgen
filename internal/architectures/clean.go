package architectures

func Clean(lang string, args ...string) {
	dirs := getCleanDirs(lang)

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

func getCleanDirs(lang string) *cleanDirs {
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

	mainDir := ""

	switch lang {
	case "go":
		mainDir = "internal"
	case "node", "python":
		mainDir = "src"
	case "java":
		mainDir = "src/main/java/com/empresa/app"
	}

	cleanDirectories := &cleanDirs{
		mainDir:  mainDir,
		children: []*cleanSubDirs{domain, application, interfaces, config},
	}

	return cleanDirectories
}