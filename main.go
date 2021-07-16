package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core"
	"github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = core.NotFoundHandler("Main")
	err := src.ListenAndServe(":9999", router)
	if err != nil {
		log.Fatal("Error On Server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}