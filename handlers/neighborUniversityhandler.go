package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func NeighbourUnisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequest2Handler(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not supprted. Currently only '"+
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

	country := getCountryByName(w, name)

	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(country)
	if err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
	}

	var listLength int = len(country.Borders)
	//uniInfoOutput := make([]UniInfo, listLength)
	for i := 0; i < listLength; i++ {
		neighbourCountry := country.Borders[i]
		neighbourUniversities := getUniversities(w, neighbourCountry)
		var listLength2 int = len(neighbourCountry)
		for i := 0; i < listLength2; i++ {
			var outputUniversities []University
			outputUniversities = append(outputUniversities, neighbourUniversities...)
		}
	}
}

func getCountryByName(w http.ResponseWriter, name string) Country {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	response, err := client.Get(RestCountriesNamePath +
		name + RestCountriesTextPath)

	if err != nil {
		http.Error(w, "Error in response from service", http.StatusInternalServerError)
	}

	//decoding json
	decoder := json.NewDecoder(response.Body)
	var country Country
	err2 := decoder.Decode(&country)
	if err2 != nil {
		http.Error(w, "Error when decoding response body from json to list", http.StatusInternalServerError)
	}

	return country
}

//getLimit
