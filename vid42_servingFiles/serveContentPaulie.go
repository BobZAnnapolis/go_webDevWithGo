package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	fmt.Println("\n\tbring up browser localhost:8080\n\t\tenter / or /pauley.jpg to see PauleyPerrette local jpg - file has to be in local dir\nServeContent rarely used")
	http.HandleFunc("/", pauley)
	http.HandleFunc("/pauley.jpg", pauleypic)
	http.ListenAndServe(":8080", nil)
}

func pauley(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/pauley.jpg">`)
}

func pauleypic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("pauley.jpg")
	if err != nil {
		http.Error(w, "File not found,", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not found,", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)

}
