package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nfrom the command line, direct output to an HTML file - then open in a browser\n")

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

	fmt.Println(tpl)

}
