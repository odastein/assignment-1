package handlers

import (
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

func getRequest2Handler(w http.ResponseWriter, r *http.Request) []Country {
	urlParts := strings.Split(r.URL.Path, "/")
	fmt.Println(urlParts)
	fmt.Println(len(urlParts))

	// check if the user has added 2-3 search word2
	if len(urlParts) <= 5 && urlParts[4] == "" {
		http.Error(w, "Please enter 2 or 3 search words!", http.StatusBadRequest)
		return nil
	}
}

//getCountryName

//getPartialOrCompleteUniversityName

//getLimit
