package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
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

	fmt.Printf("\nuse cmd line - 1st parm gets filename, 2nd parm gets 1 word to move into created file, create file, write string into it - open new file in a browser %s\n", fname)

}
