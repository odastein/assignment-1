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

	universityInput, err := getUniversities(urlParts[4])
	if err != nil {
		//todo
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
			//todo
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

	// create new request
	request, err1 := http.NewRequest(http.MethodGet,
		UniversityURL+SearchNameURL+
			strings.ReplaceAll(universityName, " ", "%20"), nil)

	if err1 != nil {
		return nil, err1
	}
	// give request header
	request.Header.Add("Content-Type", "application/json")

	// instantiate client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	//issue request
	response, err2 := client.Do(request)

	if err2 != nil {
		return nil, err2
	}

	// decode json
	decoder := json.NewDecoder(response.Body)
	var universities []University
	err3 := decoder.Decode(&universities)
	if err3 != nil {
		return nil, err3
	}
	return universities, nil
}

func getCountry(alphaCode string) (Country, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	response, err1 := client.Get(RestCountriesAlphaPath +
		alphaCode + RestCountriesFieldsPath)

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
