package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("\nuse cmd line - 1st parm is a filename, 2nd parm is 1 word to move into a newly created html, open new file in a browser")

	fname := os.Args[1]
	ftext := os.Args[2]

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	tpl := `
  <!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>Web Dev with Go</title>
  </head>
  <body>
  <h1>` + ftext + `</h1>
  </body>
  </html>
  `

	nf, err := os.Create(fname)
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))

	fmt.Printf("\nopen new file %s in a browser\n", fname)

}
