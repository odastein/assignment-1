package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NeighbourUnisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequest2Handler(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusNotImplemented)
	}
}

func getRequest2Handler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")

	// check if the user has added 2-3 search words
	if len(urlParts) < 6 || urlParts[5] == "" || urlParts[4] == "" {
		http.Error(w, "Please enter 2 or 3 search words!", http.StatusBadRequest)
		return
	}

	name := urlParts[4]

	limitString := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit < 0 {
		http.Error(w, "Please enter a positive number", http.StatusBadRequest)
	}

	country, err2 := getCountryByName(w, name)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
	var outputInfo []UniInfo

	var listLength = len(country.Borders)
	for i := 0; i < listLength; i++ {
		alphaCode := country.Borders[i]
		borderCountry := getCountry(w, alphaCode)
		borderUniversities := getUniversities(w, borderCountry.Name.Common)
		var listLength2 = len(borderUniversities)
		for j := 0; j < limit && j < listLength2; j++ {
			outputInfo = append(outputInfo, UniInfo{Name: borderUniversities[j].Name,
				Country: borderUniversities[j].Country, Isocode: borderUniversities[j].AlphaTwoCode,
				Map: borderCountry.Maps.OpenStreetMaps, Webpages: borderUniversities[j].WebPages,
				Languages: borderCountry.Languages})
		}
	}
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err3 := encoder.Encode(outputInfo)
	if err3 != nil {
		http.Error(w, "Error during encoding: "+err3.Error(), http.StatusInternalServerError)
	}
}

func getCountryByName(w http.ResponseWriter, name string) (Country, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	response, err := client.Get(RestCountriesNamePath +
		name + RestCountriesTextPath)

	if err != nil {
		return Country{}, fmt.Errorf("error in response from service")
	}

	//decoding json
	decoder := json.NewDecoder(response.Body)
	var country Country
	err2 := decoder.Decode(&country)
	if err2 != nil {
		return country, fmt.Errorf("error when decoding response body from json to list")
	}

	return country, nil
}

func getUniByCountryAndName(w http.ResponseWriter, country string, universityName string) []University {
	foundCountry, found := countryNames[country]
	if found {
		country = foundCountry
	}

	country = url.QueryEscape(country)
	universityName = url.QueryEscape(universityName)

	// create new request
	request, err1 := http.NewRequest(http.MethodGet,
		UniversityURL+SearchNameURL+universityName+SearchCountryURL+country, nil)

	if err1 != nil {
		http.Error(w, "Error when creating request to dependency", http.StatusInternalServerError)
	}
	// give request header
	request.Header.Add("Content-Type", "application/json")

	// instantiate client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	//issue request
	response, err2 := client.Do(request)

	if err2 != nil {
		http.Error(w, "Error in response from service", http.StatusInternalServerError)
	}

	// decode json
	decoder := json.NewDecoder(response.Body)
	var universities []University
	err3 := decoder.Decode(&universities)
	if err3 != nil {
		http.Error(w, "Error when decoding response body from json to list", http.StatusInternalServerError)
	}
	return universities
}
