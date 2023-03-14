package API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Relations struct {
	Relations interface{}
}

func getArtistWithDatesAndLocations(id string) (*Artist, error) {
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)
	datesURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", id)
	locationsURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id)
	relationsURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)

	artistResp, err := http.Get(artistURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get artist: %v", err)
	}
	defer artistResp.Body.Close()

	var artist Artist
	err = json.NewDecoder(artistResp.Body).Decode(&artist)
	if err != nil {
		return nil, fmt.Errorf("failed to decode artist response: %v", err)
	}

	datesResp, err := http.Get(datesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get dates: %v", err)
	}
	defer datesResp.Body.Close()

	var dates interface{}
	err = json.NewDecoder(datesResp.Body).Decode(&dates)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates response: %v", err)
	}

	locationsResp, err := http.Get(locationsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get locations: %v", err)
	}
	defer locationsResp.Body.Close()

	var locations interface{}
	err = json.NewDecoder(locationsResp.Body).Decode(&locations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode locations response: %v", err)
	}

	relationsResp, err := http.Get(relationsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get dates: %v", err)
	}
	defer relationsResp.Body.Close()

	var relations interface{}
	err = json.NewDecoder(relationsResp.Body).Decode(&relations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates response: %v", err)
	}
	relationsStr := ""
	if relationsMap, ok := relations.(map[string]interface{}); ok {
		var relationsArr []interface{}
		for _, v := range relationsMap {
			relationsArr = append(relationsArr, v)
		}
		relationsStrBytes, err := json.Marshal(relationsArr)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal relations: %v", err)
		}
		relationsStr = string(relationsStrBytes)
		relationsStr = strings.Replace(relationsStr, "map", "", -1)
	}

	artist.Dates = dates
	artist.Locations = locations
	artist.Relations = relationsStr

	return &artist, nil
}
