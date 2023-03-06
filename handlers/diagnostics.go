package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

var StartTime time.Time

func upTime() string {
	return time.Since(StartTime).Round(time.Second).String()
}

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequest3Handler(w, r)
		break
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

func getRequest3Handler(w http.ResponseWriter, r *http.Request) {
	var universitiesApi int
	var countriesApi int

	responseUni, errUni := http.Get(UniversityURL)
	if errUni != nil {
		universitiesApi = http.StatusInternalServerError
	} else {
		universitiesApi = responseUni.StatusCode
	}

	responseCoun, errCoun := http.Get(RestCountriesPath)
	if errCoun != nil {
		countriesApi = http.StatusInternalServerError
	} else {
		countriesApi = responseCoun.StatusCode
	}

	diagnosticsInfoOutput := DiagnosticsInfo{UniversitiesApi: universitiesApi, CountriesApi: countriesApi,
		Version: versionPath, Uptime: upTime()}

	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err3 := encoder.Encode(diagnosticsInfoOutput)
	if err3 != nil {
		http.Error(w, "Error during encoding: "+err3.Error(), http.StatusInternalServerError)
		return
	}
}
