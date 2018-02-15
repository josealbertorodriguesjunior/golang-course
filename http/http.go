package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("serving on http://localhost:7777/hello")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "<html><body><h1>Hello With Golang</h1></body></html>")
}
