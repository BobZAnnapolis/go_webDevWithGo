package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func main() {

	fmt.Println("\nbring up localhost:8080, enter /dog or /cat in url\n")

	var c hotcat
	var d hotdog

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)

}
