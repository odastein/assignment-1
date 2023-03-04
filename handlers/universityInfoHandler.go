package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

// UniInfoHandler Takes in a request, and sends a response based on the request.
// The request is expected to be GET.
func UniInfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequestHandler(w, r)
		http.Error(w, "Everything is ok", http.StatusOK)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// getRequestHandler bla blabla
func getRequestHandler(w http.ResponseWriter, r *http.Request) {

	universityInput := getUniversities(w, r)
	/*
		uniInfo := UniInfo{
			Name:      "NTNU",
			Country:   "Norway",
			Isocode:   "NO",
			Webpages:  []string{"http://www.ntnu.no/"},
			Languages: map[string]string{"nno": "Norwegian Nynorsk", "nob": "Norwegian Bokmål", "smi": "Sami"},
			Map:       "https://openstreetmap.org/relation/2978650"}
	*/

	var listLength int = len(universityInput)
	uniInfoOutput := make([]UniInfo, listLength)
	for i := 0; i < listLength; i++ {
		uniInfoOutput[i].Name = universityInput[i].Name
		uniInfoOutput[i].Country = universityInput[i].Country
		uniInfoOutput[i].Webpages = universityInput[i].WebPages
		uniInfoOutput[i].Isocode = universityInput[i].AlphaTwoCode
	}

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(uniInfoOutput)
	if err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
	}
}

func getUniversities(w http.ResponseWriter, r *http.Request) []University {
	urlParts := strings.Split(r.URL.Path, "/")

	// check if the user has added a search word
	if len(urlParts) > 5 {
		http.Error(w, "Too many search words!", http.StatusBadRequest)
		return nil
	}
	if len(urlParts) <= 5 && urlParts[4] == "" {
		http.Error(w, "Please enter a search word!", http.StatusBadRequest)
		return nil
	}

	// create new request
	request, err1 := http.NewRequest(http.MethodGet,
		UniversityURL+SearchURL+urlParts[4], nil)

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
