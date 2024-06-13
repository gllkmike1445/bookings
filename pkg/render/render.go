package render

import (
	"bytes"
	"github.com/gllkmike1445/bookings/pkg/config"
	"github.com/gllkmike1445/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// Newtamplate sets the config for the template package
func Newtamplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//get the template cache from the app config
	var tc map[string]*template.Template
	if app.UsaCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Error loading template:", tmpl)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	td = AddDefaultData(td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error writing template:")
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all the files name *.page.tmpl from ./templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
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
