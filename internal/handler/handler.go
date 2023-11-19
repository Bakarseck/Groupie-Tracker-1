package handler

import (
	"Tracker/internal/app"
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderHandler(w, "home", app.Images)
}
func Infohandler(w http.ResponseWriter, r *http.Request) {

	RenderHandler(w, "info", nil)
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
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, Value)
	buffer.WriteTo(w)
}

func TemplateCacheHandler() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)
	pages, err := filepath.Glob("./template/*.page.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))
		layout, err := filepath.Glob("./template/layout/*.layout.tmpl")
		if err != nil {
			return nil, err
		}
		if len(layout) > 0 {
			tmpl.ParseGlob("./template/layout/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
