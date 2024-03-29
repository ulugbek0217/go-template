package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Renders template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// Creates and returns template cache
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

		fmt.Println("Current page:", pageName)

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
			templateSet, err = template.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return pagesCache, err
			}
		}
		// store template set to page cache under the name of page file
		pagesCache[pageName] = templateSet

	}
	return pagesCache, nil

}
