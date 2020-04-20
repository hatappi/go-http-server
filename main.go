package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("GET: /")
		io.WriteString(w, "Hello World!!\n")
	}

	http.HandleFunc("/", handler)

	log.Println("start http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
