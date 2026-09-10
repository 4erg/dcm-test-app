package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "holaaaa prueba de desplege ")
	})

	http.ListenAndServe(":8080", nil)
}