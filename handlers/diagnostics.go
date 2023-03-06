package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// DiagHandler Takes in a request, and sends a response based on the request.
// The request is expected to be GET.
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

// getRequest3Handler gets the input, and makes an output
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

var StartTime time.Time

// upTime returns the time from starttime to now
func upTime() string {
	return time.Since(StartTime).Round(time.Second).String()
}
