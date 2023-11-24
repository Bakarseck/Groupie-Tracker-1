package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Geolocalisation(strg []string) {
	for _, str := range strg {
		url := "https://api.opencagedata.com/geocode/v1/json?q=URI-ENCODED-" + str + "&key=61856ea91e1340f0af761c46307ffa1e"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var Response Localisation

		err = json.Unmarshal(body, &Response)
		if err != nil {
			log.Fatal(err)
		}
		var Coordoones LatLng
		Coordoones.Lat =Response.Results[0].Annotations.Dms.Lat
		Coordoones.Lng =Response.Results[0].Annotations.Dms.Lng
		Inf.Coordoones= append(Inf.Coordoones, Coordoones)
	}
}
