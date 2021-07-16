package core

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Message string
	Status int
}

// NotFoundHandler Handle 404
func NotFoundHandler(domain string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Request was not found On " + domain + " Domain",
		})
		if err != nil {
			log.Printf("Error On Json Encode : %s", err)
		}
	}
}