package main

import (
	bim "Tracker/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", bim.Homehandler)
	fmt.Println("https://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
