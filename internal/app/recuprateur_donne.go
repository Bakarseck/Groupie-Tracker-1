package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var Images PageData

func Recuperation(str string) {
	JsonFile, err := http.Get(str)
	if err != nil {
		panic(err)
	}
	defer JsonFile.Body.Close()
	if JsonFile.StatusCode != http.StatusOK {
		fmt.Println("La requête a retourné un code de statut non 200 OK:", JsonFile.Status)
		return
	}
	lire, err := io.ReadAll(JsonFile.Body)
	if err != nil {
		panic(err)
	}
	var Link Value
	err = json.Unmarshal(lire, &Link)
	if err != nil {
		panic(err)
	}
	JsonFile1, err := http.Get(Link.Artists)
	if err != nil {
		panic(err)
	}
	defer JsonFile1.Body.Close()
	if JsonFile1.StatusCode != http.StatusOK {
		fmt.Println("La requête a retourné un code de statut non 200 OK:", JsonFile1.Status)
		return
	}
	lir, err := io.ReadAll(JsonFile1.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(lir, &Images.Photo)
	if err != nil {
		log.Fatal(err)
	}
}
