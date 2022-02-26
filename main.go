package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello EIF Team !"))
	})

	http.ListenAndServe(":9999", nil)
}
