package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("\n\tbring up browser localhost:8080\n\t\tenter /, /abby.png to see pic - using stripPrefix")

	http.HandleFunc("/", abby)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func abby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/abby.png">`)
}
