package architectures

func MVC(lang, dir string) {
	dirs := getMVCDirs()

	createDirs(dirs, dir, lang)
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