package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/twinj/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}

func main() {

	fmt.Println("\ngo to localhost:8080 in browser")
	fmt.Println("\tuses cookie named session - use browser dev tools to check this cookie")

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// add route to service pictures
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		// create sha1 for the filename
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		path := filepath.Join(wd, "assets", "uploaded", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		//add filename to this
		c = appendValue(w, c, fname)

	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.html", xs)
}

// getCookie - handles geting / creating cookie on server
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	// append picnames to cookie
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
