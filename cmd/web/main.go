package main

import (
	"Tracker/internal/app"
	bim "Tracker/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	app.Recuperation("https://groupietrackers.herokuapp.com/api")
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))
	http.HandleFunc("/", bim.HomeHandler)
	http.HandleFunc("/info", bim.Infohandler)
	http.HandleFunc("/search", bim.SearchHandler)
	fmt.Printf("127.0.0.1:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
		return
	}
}
