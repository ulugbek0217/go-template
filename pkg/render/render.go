package render

import (
	"bytes"
	"github.com/ulugbek0217/template/pkg/config"
	"github.com/ulugbek0217/template/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates *config.TemplatesConfig

// SetTemplateConfig sets template cache to the templates variable
func SetTemplateConfig(templateSet *config.TemplatesConfig) {
	templates = templateSet
}

func AddDefaultTD(td *models.TemplateData) *models.TemplateData {
	td.IntMap = map[string]int{
		"age": 18,
	}
	return td
}

// RenderTemplate renders templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// check if the mode is real time editing or using cache
	var tc map[string]*template.Template
	if templates.UseCache {
		tc = templates.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't get template config")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultTD(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache Creates and returns template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	functions := template.FuncMap{}
	// place to store template sets
	pagesCache := map[string]*template.Template{}

	// searching for files that ends with page.tmpl
	files, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return pagesCache, err
	}

	for _, page := range files {
		// getting base name of a file
		pageName := filepath.Base(page)

		templateSet, err := template.New(pageName).Funcs(functions).ParseFiles(page)
		if err != nil {
			return pagesCache, err
		}
		// searching for layouts to render pages based on them
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return pagesCache, err
		}
		// check if at least a layout exists
		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return pagesCache, err
			}
		}
		// store template set to page cache under the name of page file
		pagesCache[pageName] = templateSet

	}
	return pagesCache, nil

}
