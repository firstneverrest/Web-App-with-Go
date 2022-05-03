package config

import "html/template"

// application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
