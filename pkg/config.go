package pkg

import "html/template"

// AppConfig structures holds application wide configuration information.
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
