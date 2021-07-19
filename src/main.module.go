package src

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user"
	user_entities "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user/entities"
	"log"
	"net/http"
)

// ListenAndServe starts up the server
func ListenAndServe(address string, r *mux.Router) error {
	migratePostgres()

	http.Handle("/", core.LoggingMiddleware(r))
	http.Handle("/user/", core.LoggingMiddleware(user.Module()))

	return http.ListenAndServe(
		address,
		nil,
	)
}

func migratePostgres() {
	gorm := core.PostgresConnect()
	err := gorm.AutoMigrate(&user_entities.User{})

	if err != nil {
		log.Fatal(err)
	}
}
