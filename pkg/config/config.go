package config

import "html/template"

type TemplatesConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
