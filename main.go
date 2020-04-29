package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "HHello World!!!!\n")
	}

	http.HandleFunc("/", handler)

	port := ":8080"

	log.Println("start http server. port is " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
