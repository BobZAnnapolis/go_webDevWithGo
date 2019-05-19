package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("\n\tbring up browser against localhost:8080 - should see a puppy pic - file served from Wiki")
	http.HandleFunc("/", puppy)
	http.ListenAndServe(":8080", nil)
}

func puppy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
    <!--serve image from Wiki-->
    <img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
    `)
}
