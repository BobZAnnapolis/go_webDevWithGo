package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	fname := "newlyCreated.gohtml"
	nf, err := os.Create(fname)
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nread text from templated file, write into a new GOHTML file \n\t%s\n", fname)

}
