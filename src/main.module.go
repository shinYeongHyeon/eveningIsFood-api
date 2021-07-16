package src

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user"
	"net/http"
)

// ListenAndServe starts up the server
func ListenAndServe(address string, r *mux.Router) error {
	http.Handle("/", core.LoggingMiddleware(r))
	http.Handle("/user/", core.LoggingMiddleware(user.Module()))

	return http.ListenAndServe(
		address,
		nil,
	)
}
