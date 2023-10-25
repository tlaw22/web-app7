package pkg

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders a template
// This method taxes the server more because it has to read from disk, each time a user requests a page.
func RenderTemplate2(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

// Create a map to store the cached pages
var tc = make(map[string]*template.Template)

// re-create the render template function to cache the files instead of loading the files from disk onload. This methods caches all of the HTML files inside of RAM. Having to read from disk 1000s of times per day could start a fire.
func RenderTemplate(w http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error
	// check to see if the template has been updated
	_, inMap := tc[t]
	if !inMap {
		// need to create template
		log.Println("Adding page to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// The template is in the cache
		log.Println("Using cached template...")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
