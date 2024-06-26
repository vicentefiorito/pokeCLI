package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	// if the page url is set
	// use it as the fullurl
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache here
	data, ok := c.cache.Get(fullURL)

	if ok {
		// cache hit
		fmt.Println("cache hit")
		// unmarshalling the json
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			fmt.Println(err)
		}

		return locationAreasResp, nil
	}

	//making the request to the API
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// closes the resp object before the function returns
	defer resp.Body.Close()

	// checking status code of the response
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// reading the response body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// unmarshalling the json
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		fmt.Println(err)
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}
