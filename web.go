package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)                 // query string
	fmt.Println("path", r.URL.Path)     // path
	fmt.Println("scheme", r.URL.Scheme) // ??
	fmt.Println(r.Form["url_long"])     // query string url_long array value
	for k, v := range r.Form {          // loop query string
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // return this string as response
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
