package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	myName := "Bob Zee"

	tpl := `
  <!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>Web Dev with Go</title>
  </head>
  <body>
  <h1>` + myName + `</h1>
  </body>
  </html>
  `

	fname := "index.html"
	nf, err := os.Create(fname)
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
	fmt.Printf("\nusing defined string, creates a new file and writes string into it - open new file in a browser %s\n", fname)

}
