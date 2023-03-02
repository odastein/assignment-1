package handlers

import (
	"encoding/json"
	"net/http"
)

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
func getRequestHandler(w http.ResponseWriter) {

	uniInfo := UniInfo{
		Name:      "NTNU",
		Country:   "Norway",
		Isocode:   "NO",
		Webpages:  []string{"http://www.ntnu.no/"},
		Languages: map[string]string{"nno": "Norwegian Nynorsk", "nob": "Norwegian Bokmål", "smi": "Sami"},
		Map:       "https://openstreetmap.org/relation/2978650"}

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(uniInfo)
	if err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
	}
}
