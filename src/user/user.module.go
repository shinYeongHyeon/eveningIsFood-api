package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core"
	"net/http"
)

func Module() http.Handler {
	router := mux.NewRouter().PathPrefix("/user/").Subrouter()

	router.NotFoundHandler = core.NotFoundHandler("User")
	router.HandleFunc("/", Handle)

	return router
}

func Handle(w http.ResponseWriter, r *http.Request	) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User: 3\n")
}