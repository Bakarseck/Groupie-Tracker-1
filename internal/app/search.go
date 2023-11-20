package app

import (
	"fmt"
)

var Inf Info
var Glob Global

func Search(id int, artists PageData, restant LOcDatRel) {
	if len(restant.Location) > 0 || len(restant.Date) > 0 || len(restant.Relation) > 0 || len(Artists.Artist) > 0 {
		Inf.Artist = artists.Artist[id-1]
		Inf.Location = restant.Location["index"][id-1]
		Inf.Date = restant.Date["index"][id-1]
		Inf.Relation = restant.Relation["index"][id-1]
	} else {
		fmt.Printf("Donnes recuperer not valable")
	}
}
