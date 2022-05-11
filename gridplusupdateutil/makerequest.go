package main

import (
	"gridplusupdateutil/util"
	"io/ioutil"
	"net/http"
)

const url = "https://interview-release-catalog-api.staging-gridpl.us/update"

// makeRequest makes the request adn returns the response struct.
func makeRequest() (*util.Response, error) {
	payload, err := util.GetCurrent()
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", url, payload)
	req.SetBasicAuth("lattice1", "codetest")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return util.GetResponse(body)
}
