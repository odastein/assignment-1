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
		getRequest1Handler(w, r)
		http.Error(w, "Everything is ok", http.StatusOK)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// getRequest1Handler bla blabla
func getRequest1Handler(w http.ResponseWriter, r *http.Request) {

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
	client := &http.Client{}
	defer client.CloseIdleConnections()
	for i := 0; i < listLength; i++ {
		uniInfoOutput[i].Name = universityInput[i].Name
		uniInfoOutput[i].Country = universityInput[i].Country
		uniInfoOutput[i].Webpages = universityInput[i].WebPages
		uniInfoOutput[i].Isocode = universityInput[i].AlphaTwoCode
		alfaCodeLowerCase := strings.ToLower(universityInput[i].AlphaTwoCode)
		// Todo implement cache
		country := getCountry(w, alfaCodeLowerCase)
		uniInfoOutput[i].Languages = country.Languages
		uniInfoOutput[i].Map = country.Maps.OpenStreetMaps
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
	if len(urlParts) <= 5 && urlParts[4] == "" {
		http.Error(w, "Please enter a search word!", http.StatusBadRequest)
		return nil
	}

	// create new request
	request, err1 := http.NewRequest(http.MethodGet,
		UniversityURL+SearchURL+
			strings.ReplaceAll(urlParts[4], " ", "%20"), nil)

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

func getCountry(w http.ResponseWriter, alphacode string) Country {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	response, err := client.Get("https://restcountries.com/v3.1/alpha/" +
		alphacode + "?fields=name,maps,languages,borders")

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
