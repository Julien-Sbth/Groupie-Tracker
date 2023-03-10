package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

func geocode(address string) (float64, float64, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=YOUR_API_KEY", address)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var result GeocodeResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, 0, err
	}

	if len(result.Results) > 0 {
		return result.Results[0].Geometry.Location.Lat, result.Results[0].Geometry.Location.Lng, nil
	} else {
		return 0, 0, fmt.Errorf("no results found")
	}
}

func convertAddressesToCoordinates(addresses []string) ([][]float64, error) {
	var coordinates [][]float64
	for _, address := range addresses {
		lat, lng, err := geocode(address)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, []float64{lat, lng})
	}
	return coordinates, nil
}
