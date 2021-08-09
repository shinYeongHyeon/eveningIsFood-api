package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core"
	createuserusecase "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user/application/CreateUserUseCase"
	createuserusecasedto "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user/application/CreateUserUseCase/dto"
	"net/http"
)

// Module is sub-router for user-domain
func Module() http.Handler {
	router := mux.NewRouter().PathPrefix("/users").Subrouter()

	router.NotFoundHandler = core.NotFoundHandler("User")
	router.HandleFunc("", Handle).Methods("GET", "POST")

	return router
}

// Handle User sub-router handler
func Handle(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode("GET-USER")
	case "POST":
		var createUserUseCaseApiRequest createuserusecasedto.Request
		json.NewDecoder(r.Body).Decode(&createUserUseCaseApiRequest)
		createuserusecase.Exec(createUserUseCaseApiRequest)
		rw.WriteHeader(http.StatusCreated)
	}
}
