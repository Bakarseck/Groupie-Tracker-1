package app

import (
	"fmt"
	"strconv"
	"strings"
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

func Trie(Artist map[int]string) Listeinfo {
	mape := Listeinfo{}
	for id, _ := range Artist {
		for _, artist := range AllArtists {
			if artist.ID == id {
				mape.Artists = append(mape.Artists, artist)
			}
		}
	}
	return mape
}

func Remplissage(AllArtistsnotorder Global) []ArtistInfo {
	for _, art := range AllArtistsnotorder.Artist.Artist {
		var artist ArtistInfo
		artist.ID = art.ID
		artist.Image = art.Image
		artist.Name = art.Name
		artist.Members = art.Members
		artist.CreationDate = art.CreationDate
		artist.FirstAlbum = art.FirstAlbum

		for _, art := range AllArtistsnotorder.DonneRestant.Location {
			for _, c := range art {
				if c.Id == artist.ID {
					artist.Location = c.Location
				}
			}
		}
		for _, art := range AllArtistsnotorder.DonneRestant.Date {
			for _, c := range art {
				if c.Id == artist.ID {
					artist.Dates = c.Dates
				}
			}
		}
		for _, art := range AllArtistsnotorder.DonneRestant.Relation {
			for _, c := range art {
				if c.Id == artist.ID {
					artist.Relations = c.Relations
				}
			}
		}
		AllArtists = append(AllArtists, artist)
	}
	return AllArtists
}

func SearchBar(text string, AllInfo []ArtistInfo) (map[int]string, int) {
	result := make(map[int]string)
	count := 0
	for _, Artiste := range AllInfo {
		switch {
		case strings.Contains(strings.ToLower(Artiste.Name), strings.ToLower(text)):
			result[Artiste.ID] = Artiste.Name
			count++

		case strings.Contains(strings.ToLower(strconv.Itoa(Artiste.CreationDate)), strings.ToLower(text)):
			result[Artiste.ID] = strconv.Itoa(Artiste.CreationDate)
			count++
		case strings.Contains(strings.ToLower(Artiste.FirstAlbum), strings.ToLower(text)):
			result[Artiste.ID] = Artiste.FirstAlbum
			count++

		}
		for _, member := range Artiste.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(text)) {
				result[Artiste.ID] = member
				count++
			}
		}
		for _, Loca := range Artiste.Location {
			if strings.Contains(strings.ToLower(Loca), strings.ToLower(text)) {
				result[Artiste.ID] = Loca
				count++
			}
		}
		for _, Date := range Artiste.Dates {
			if strings.Contains(strings.ToLower(Date), strings.ToLower(text)) {
				result[Artiste.ID] = Date
				count++
			}
		}
		for index, Relat := range Artiste.Relations {
			for _, Rela := range Relat {
				if strings.Contains(strings.ToLower(Rela), strings.ToLower(text)) || strings.Contains(strings.ToLower(index), strings.ToLower(text)) {
					result[Artiste.ID] = Rela
					count++
				}
			}
		}
	}
	return result, count
}
