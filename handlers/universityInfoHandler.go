package handlers

import (
	"encoding/json"
	"net/http"
)

// UniInfoHandler Takes in a request, and sends a respond based on the request.
// The request is expected to be GET.
func UniInfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequestHandler(w)
		http.Error(w, "Everything is ok", http.StatusOK)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// getRequestHandler bla blabla
func getRequestHandler(w http.ResponseWriter) {

	universityInput := getUniversities(w)
	/*
		uniInfo := UniInfo{
			Name:      "NTNU",
			Country:   "Norway",
			Isocode:   "NO",
			Webpages:  []string{"http://www.ntnu.no/"},
			Languages: map[string]string{"nno": "Norwegian Nynorsk", "nob": "Norwegian Bokmål", "smi": "Sami"},
			Map:       "https://openstreetmap.org/relation/2978650"}
	*/

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(universityInput)
	if err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
	}
}

func getUniversities(w http.ResponseWriter) []University {
	// create new request
	request, err1 := http.NewRequest(http.MethodGet,
		"http://universities.hipolabs.com/search?name=norwegian", nil)

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
