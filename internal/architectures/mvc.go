package architectures

func MVC(args ...string) {
	dirs := getMVCDirs()

	createDirs(dirs, args...)
}

func getMVCDirs() *cleanDirs {
	model := &cleanSubDirs{
		parent:   "model",
		children: []string{},
	}

	view := &cleanSubDirs{
		parent:   "view",
		children: []string{},
	}

	controller := &cleanSubDirs{
		parent:   "controller",
		children: []string{},
	}

	mvc := &cleanDirs{
		mainDir:  "internal",
		children: []*cleanSubDirs{model, view, controller},
	}

	return mvc
}