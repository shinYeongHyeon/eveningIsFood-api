package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src"
	"log"
	"net/http"
)

type User struct {
	ID 	 int `gorm:"primaryKey"`
	Name string
	Favor string
}

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = core.NotFoundHandler("Main")
	router.HandleFunc("/", handler)
	err := src.ListenAndServe(":9999", router)
	if err != nil {
		log.Fatal("Error On Server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	gorm := core.PostgresConnect()

	gorm.AutoMigrate(&User{})

	user := &User{Name: "Shin", Favor: "abcd"}
	gorm.Create(&user)

	fmt.Fprintf(w, "Hello")
	// fmt.Fprintf(w, strconv.Itoa(user.ID))
}