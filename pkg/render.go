package pkg

import (
	"bytes"
	"fmt"
	"go-web-app7/pkg"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template
// This method taxes the server more because it has to read from disk, each time a user requests a page.
var app *pkg.AppConfig

func NewTemplate(a *pkg.AppConfig) {
	app = a
}
var tc map[string]*template.Template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	if app.TemplateCache == nil {
	tc := app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get template from cashe
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//	myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	// get all of the files named*,tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
