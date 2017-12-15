package app

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	// TemplateDir is the directory containing the html/template template files.
	TemplateDir = filepath.Join(defaultBase("github.com/donachys/kubevalonline/app"), "templates")
)

func renderTemplate(rw http.ResponseWriter, tmpl string) error {
	log.Println(tmpl)
	t, err := template.New(tmpl).ParseFiles(filepath.Join(TemplateDir, tmpl))
	if err != nil {
		return err
	}

	return t.Execute(rw, nil)
}
