package handlers

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(urlParts)
	fmt.Println(len(urlParts))

	// check if the user has added 2-3 search words
	if len(urlParts) < 6 || urlParts[5] == "" || urlParts[4] == "" {
		http.Error(w, "Please enter 2 or 3 search words!", http.StatusBadRequest)
		return
	}

	country := getCountry(w, name)
}

func getCountryName(w http.ResponseWriter, name string) Country {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	response, err := client.Get("https://restcountries.com/v3.1/alpha/" +
		name + "?fields=name,maps,languages,borders")

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

//getPartialOrCompleteUniversityName

//getLimit
