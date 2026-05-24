package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) GetLocationList() (PokeLocationResponse, error) {
	url := "location-area/"
	fullUrl := baseUrl+url

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return PokeLocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLocationResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return PokeLocationResponse{}, fmt.Errorf("Bad status code %v", resp.StatusCode)
	}

	bt, err := io.ReadAll(resp.Body)// this is using the standard io package to read the response body
	 if err != nil {
		return PokeLocationResponse{}, fmt.Errorf("Error reading file %v ", err)
	 }

	 locationResponse := PokeLocationResponse{}

	 err = json.Unmarshal(bt, &locationResponse)
	 if err != nil {
		return PokeLocationResponse{}, err
	}

	return  locationResponse, nil

	
}