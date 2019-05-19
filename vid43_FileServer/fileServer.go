package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("\n\tbring up browser localhost:8080\n\t\tbrowser should appear as a browser-based File Browser\n\tadd or remove an index.html file to protect folder structure")

	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
