package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var Artists PageData
var DonneRestant LOcDatRel

func Recuperation(str string) {
	lire, err := MiniRecup(str)
	if err != nil {
		panic(err)
	}
	var Link Value
	err = json.Unmarshal(lire, &Link)
	if err != nil {
		panic(err)
	}
	lir, err := MiniRecup(Link.Artists)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(lir, &Artists.Artist)
	if err != nil {
		log.Fatal(err)
	}
	Location, err := MiniRecup(Link.Locations)
	if err != nil {
		panic(err)
	}
	var loc = make(map[string][]Locations)
	err = json.Unmarshal(Location, &loc)
	if err != nil {
		log.Fatal(err)
	}
	DonneRestant.Location = loc
	Date, err := MiniRecup(Link.Dates)
	if err != nil {
		panic(err)
	}
	var date = make(map[string][]Dates)
	err = json.Unmarshal(Date, &date)
	if err != nil {
		log.Fatal(err)
	}
	DonneRestant.Date = date
	Relatio, err := MiniRecup(Link.Relation)
	if err != nil {
		panic(err)
	}
	var rel = make(map[string][]Relations)
	err = json.Unmarshal(Relatio, &rel)
	if err != nil {
		log.Fatal(err)
	}
	DonneRestant.Relation = rel
	Glob.Artist = Artists
	Glob.DonneRestant = DonneRestant
}



func MiniRecup(str string) ([]byte, error) {
	jfile, erreur := http.Get(str)
	if erreur != nil {
		panic(erreur)
	}
	defer jfile.Body.Close()
	if jfile.StatusCode != http.StatusOK {
		fmt.Println("La requête a retourné un code de statut non 200 OK:", jfile.Status)
		return nil, erreur
	}
	lire, erreur := io.ReadAll(jfile.Body)
	if erreur != nil {
		panic(erreur)
	}
	return lire, nil
}
