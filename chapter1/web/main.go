package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, world!")
	if err != nil {
		return
	}
}

func handlerAllData(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "data")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/go", handlerAllData)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
