package rest

import (
	"net/http"
	"encoding/json"
)

func ItemNotFoundEndpoint(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func JsonEndpoint(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err := enc.Encode(result)
	if err != nil {
		panic(err)
	}
}
