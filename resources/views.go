package resources

import (
	"embed"
	"html/template"
	"io/fs"
)

const (
	layoutsDir  = "views/layouts"
	partialsDir = "views/partials"
	viewsDir    = "views"
)

var (
	//go:embed views/* views/layouts/* views/partials/*
	files embed.FS
	Views map[string]*template.Template
)

func LoadViews() error {
	if Views == nil {
		Views = make(map[string]*template.Template)
	}

	viewFiles, err := fs.ReadDir(files, viewsDir)
	if err != nil {
		return err
	}

	for _, view := range viewFiles {
		if view.IsDir() {
			continue
		}

		parsedTemplate, err := template.ParseFS(
			files,
			viewsDir+"/"+view.Name(),
			layoutsDir+"/*.html",
			partialsDir+"/*.html",
		)
		if err != nil {
			return err
		}

		Views[view.Name()] = parsedTemplate
	}

	return nil
}
