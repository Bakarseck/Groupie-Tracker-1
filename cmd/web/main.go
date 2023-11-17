package main

import (
	"Tracker/internal/app"
	bim "Tracker/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	app.Recuperation("https://groupietrackers.herokuapp.com/api")
	http.HandleFunc("/index", bim.NewHandler)
	http.HandleFunc("/", bim.Homehandler)
	fmt.Println("https://localhost:8080")
	http.ListenAndServe(":1111", nil)
}
