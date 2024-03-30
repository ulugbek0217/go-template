package config

import "html/template"

type TemplatesConfig struct {
	TemplateCache map[string]*template.Template
}
