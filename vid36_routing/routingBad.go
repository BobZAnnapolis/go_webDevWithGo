package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {

	fmt.Println("\nbring up localhost:8080, enter /dog or /cat in url\n")

	var d hotdog
	http.ListenAndServe(":8080", d)

}
