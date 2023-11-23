package handler

import (
	"Tracker/internal/app"
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderHandler(w, "home", app.Glob)
}
func Infohandler(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	if id < 0 || id > 53 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	app.Search(id, app.Artists, app.DonneRestant)
	RenderHandler(w, "info", app.Inf)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var information app.Listeinfo
	text := r.FormValue("browser")
	resultat, count := app.SearchBar(text, app.AllArtists)
	if count == 0 || text == "" {
		tmpl := template.Must(template.New("search").Parse(`<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/water.css@2/out/water.css">
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
			<link rel="stylesheet" href="/asset/css/style.css">
			<title>Groupie-Tracker</title>
		</head>
		<body>
			<header>
				<h1>Groupie-Tracker</h1>
				<img src="/asset/img/disk.png" id="image" height="600px" width="600px" alt="">
				<a href="/"><i class="fa fa-car" style="font-size:60px;color:red;"></i>Come back home page !!</a>
			</header>
			<div >
			     <p>Aucune Correspondance Trouver❌</p>
				 
            </div>
			<script src="https://code.jquery.com/jquery-3.6.4.min.js"></script>
			<script src="/asset/js/scripte.js"></script>
		</body>
		</html>
			`))
		tmpl.ExecuteTemplate(w, "search", information)
	}else{
	information = app.Trie(resultat)
	RenderHandler(w, "search", information)
	}
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
