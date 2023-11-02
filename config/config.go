package config

import (
	"html/template"
	"log"
)

// AppConfig structures holds application wide configuration information.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
