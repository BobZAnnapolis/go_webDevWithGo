package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleDefault(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Enter /cat or /dog in url")
}

func handleCat(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func handleDog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func main() {

	fmt.Println("\nbring up localhost:8080, enter /dog or /cat in url\n")

	// using default http ListenAndServe server
	http.HandleFunc("/cat", handleCat)
	http.HandleFunc("/dog", handleDog)
	http.HandleFunc("/", handleDefault)

	http.ListenAndServe(":8080", nil)

}
