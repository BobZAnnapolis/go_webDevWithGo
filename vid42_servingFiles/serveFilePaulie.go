package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("\n\tbring up browser localhost:8080\n\t\tenter / or /pauley.jpg to see PauleyPerrette local jpg - file has to be in local dir\nServeFile()")
	http.HandleFunc("/", pauley)
	http.HandleFunc("/pauley.jpg", pauleypic)
	http.ListenAndServe(":8080", nil)
}

func pauley(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/pauley.jpg">`)
}

func pauleypic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pauley.jpg")
}
