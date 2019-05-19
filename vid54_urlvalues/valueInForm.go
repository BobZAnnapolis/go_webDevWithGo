package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("\nvisit this page : localhost:8080, or add /?q= in url, to see passing vals via GET or POST")

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post">
		<input type="text" name="q">
		<input type="submit">
		</form>
		<br />`+v)
}
