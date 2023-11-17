package handler

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	RenderHandler(w, "home", nil)
}

func RenderHandler(w http.ResponseWriter, tmplname string, Value interface{}) {
	templatecache, err := TemplateCacheHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, ok := templatecache[tmplname+".page.tmpl"]
	if !ok {
		http.Error(w, "Not found", http.StatusInternalServerError)
	}
	buffw := new(bytes.Buffer)
	tmpl.Execute(buffw, Value)
	buffw.WriteTo(w)
}

func TemplateCacheHandler() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)
	pages, err := filepath.Glob("/template/*.page.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))
		layout, err := filepath.Glob("/template/*.layout.tmpl")
		if err != nil {
			return nil, err
		}
		if len(layout) > 0 {
			tmpl.ParseGlob("/template/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
