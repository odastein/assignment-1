package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// UniInfoHandler Takes in a request, and sends a response based on the request.
// The request is expected to be GET.
func UniInfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequest1Handler(w, r)
		http.Error(w, "Everything is ok", http.StatusOK)
		break
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// getRequest1Handler bla blabla
func getRequest1Handler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")

	// check if the user has added a search word
	if len(urlParts) <= 5 && urlParts[4] == "" {
		http.Error(w, "Please enter a search word!", http.StatusBadRequest)
		return
	}

	universityInput, err1 := getUniversities(urlParts[4])
	if err1 != nil {
		http.Error(w, "There was an error "+err1.Error(), http.StatusFailedDependency)
		return
	}

	var listLength int = len(universityInput)
	uniInfoOutput := make([]UniInfo, listLength)
	for i := 0; i < listLength; i++ {
		uniInfoOutput[i].Name = universityInput[i].Name
		uniInfoOutput[i].Country = universityInput[i].Country
		uniInfoOutput[i].Webpages = universityInput[i].WebPages
		uniInfoOutput[i].Isocode = universityInput[i].AlphaTwoCode
		alfaCodeLowerCase := strings.ToLower(universityInput[i].AlphaTwoCode)
		// Todo implement cache
		country, err2 := getCountry(alfaCodeLowerCase)
		if err2 != nil {
			http.Error(w, "There was an error "+err2.Error(), http.StatusFailedDependency)
			return
		}
		uniInfoOutput[i].Languages = country.Languages
		uniInfoOutput[i].Map = country.Maps.OpenStreetMaps
	}

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err3 := encoder.Encode(uniInfoOutput)
	if err3 != nil {
		http.Error(w, "Error during encoding: "+err3.Error(), http.StatusInternalServerError)
		return
	}
}

func getUniversities(universityName string) ([]University, error) {

	response, err1 := http.Get(UniversityURL + SearchNameURL + url.QueryEscape(universityName))

	if err1 != nil {
		return nil, err1
	}

	// decode json
	decoder := json.NewDecoder(response.Body)
	var universities []University
	err2 := decoder.Decode(&universities)
	if err2 != nil {
		return nil, err2
	}
	return universities, nil
}

func getCountry(alphaCode string) (Country, error) {
	response, err1 := http.Get(RestCountriesAlphaPath + alphaCode + RestCountriesFieldsPath)

	if err1 != nil {
		return Country{}, err1
	}

	//decoding json
	decoder := json.NewDecoder(response.Body)
	var country Country
	err2 := decoder.Decode(&country)
	if err2 != nil {
		return Country{}, err2
	}

	return country, nil
}
