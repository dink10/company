package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func toJSON(w http.ResponseWriter, statusCode int, jsonObject interface{}) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if statusCode == http.StatusNoContent {
		return
	}
	err := json.NewEncoder(w).Encode(jsonObject)
	if err != nil {
		log.Printf("could not encode json return: %s", err)
	}
}

func toJSONError(w http.ResponseWriter, code int, err error) {
	toJSON(w, code, struct {
		Error string `json:"error"`
	}{
		Error: fmt.Sprintf("error on parse data: %v", err),
	})
}
