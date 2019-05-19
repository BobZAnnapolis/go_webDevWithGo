package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("\nvisit this page : localhost:8080/?q=aubreyKate, put anything after q= to pass a value")

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}
