package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Zawislak-key", "this is from bz")
	w.Header().Set("Content-Type", "text/html;charset=utf-8") // toggle html-plain to see diff output
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {

	fmt.Println("\njust execute from cmd line, bring up localhost:8080, demos http.ResponseWriter\n")

	var d hotdog
	http.ListenAndServe(":8080", d)

}
