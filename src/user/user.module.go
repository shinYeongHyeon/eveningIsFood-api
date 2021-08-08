package user

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core"
	"net/http"
)

// Module is sub-router for user-domain
func Module() http.Handler {
	router := mux.NewRouter().PathPrefix("/user/").Subrouter()

	router.NotFoundHandler = core.NotFoundHandler("User")
	router.HandleFunc("/", Handle)

	return router
}

// Handle User sub-router handler
func Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
