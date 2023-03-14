package API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

	relationResp, err := http.Get(relationsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get dates: %v", err)
	}
	defer relationResp.Body.Close()

	var relations Relation
	err = json.NewDecoder(relationResp.Body).Decode(&relations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates response: %v", err)
	}
	artist.Dates = dates
	artist.Locations = locations
	artist.Relations = relations

	return &artist, nil
}
