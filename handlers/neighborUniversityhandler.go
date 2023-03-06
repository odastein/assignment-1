package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//todo write comments for every function

// NeighbourUnisHandler Takes in a request, and sends a response based on the request.
// The request is expected to be GET.
func NeighbourUnisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequest2Handler(w, r)
		break
	default:
		http.Error(w, "REST method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusNotImplemented)
	}
}

// getRequest2Handler gets the input, and makes an output
func getRequest2Handler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")

	// check if the user has added 2-3 search words
	if len(urlParts) < 6 || urlParts[5] == "" || urlParts[4] == "" {
		http.Error(w, "Please enter 2 or 3 search words!", http.StatusBadRequest)
		return
	}

	countryName := urlParts[4]
	universityName := urlParts[5]

	var outputInfo []UniInfo
	country, err2 := getCountryByName(countryName)
	if err2 != nil {
		http.Error(w, "There was an error "+err2.Error(), http.StatusFailedDependency)
		return
	}
	var listLength = len(country.Borders)

	limitString := r.URL.Query().Get("limit")
	limit, err1 := strconv.Atoi(limitString)
	if err1 != nil {

		for i := 0; i < listLength; i++ {
			alphaCode := country.Borders[i]
			borderCountry, err3 := getCountry(alphaCode)
			if err3 != nil {
				http.Error(w, "There was an error "+err3.Error(), http.StatusFailedDependency)
				return
			}
			borderUniversities, err4 := getUniByCountryAndName(borderCountry.Name.Common, universityName)
			if err4 != nil {
				http.Error(w, "There was an error "+err4.Error(), http.StatusFailedDependency)
				return
			}
			var listLength2 = len(borderUniversities)
			for j := 0; j < limit && j < listLength2; j++ {
				outputInfo = append(outputInfo, UniInfo{Name: borderUniversities[j].Name,
					Country: borderUniversities[j].Country, Isocode: borderUniversities[j].AlphaTwoCode,
					Map: borderCountry.Maps.OpenStreetMaps, Webpages: borderUniversities[j].WebPages,
					Languages: borderCountry.Languages})
			}
		}

	} else if limit < 0 {
		http.Error(w, "Please enter a positive number", http.StatusBadRequest)
		return
	}

	for i := 0; i < listLength; i++ {
		alphaCode := country.Borders[i]
		borderCountry, err3 := getCountry(alphaCode)
		if err3 != nil {
			http.Error(w, "There was an error "+err3.Error(), http.StatusFailedDependency)
			return
		}
		borderUniversities, err4 := getUniByCountryAndName(borderCountry.Name.Common, universityName)
		if err4 != nil {
			http.Error(w, "There was an error "+err4.Error(), http.StatusFailedDependency)
			return
		}
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
	err5 := encoder.Encode(outputInfo)
	if err5 != nil {
		http.Error(w, "Error during encoding: "+err5.Error(), http.StatusInternalServerError)
		return
	}
}

// getCountryByName is a getter for the country-struct.
// Returns the country and an error, one of the will be nil
func getCountryByName(name string) (Country, error) {
	response, err1 := http.Get(RestCountriesNamePath + url.QueryEscape(name) + RestCountriesTextPath)

	if err1 != nil {
		return Country{}, err1
	}

	//decoding json
	decoder := json.NewDecoder(response.Body)
	var country []Country
	err2 := decoder.Decode(&country)
	if err2 != nil {
		return Country{}, err2
	}

	return country[0], nil
}

// getUniByCountryAndName is a getter for a list of elements of the university-struct.
// Returns a list and an error, one of them will be nil
func getUniByCountryAndName(country string, universityName string) ([]University, error) {
	foundCountry, found := countryNames[country]
	if found {
		country = foundCountry
	}

	//issue request
	response, err1 := http.Get(UniversityURL + SearchNameURL + url.QueryEscape(universityName) +
		SearchCountryURL + url.QueryEscape(country))

	if err1 != nil {
		return nil, err1
	}

	// decode json
	decoder := json.NewDecoder(response.Body)
	var universities []University
	err4 := decoder.Decode(&universities)
	if err4 != nil {
		return nil, err4
	}
	return universities, nil
}
