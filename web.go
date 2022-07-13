package main

import (
	"fmt"
	"html/template"
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

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println("----", t.Execute(w, nil))
	} else if r.Method == "POST" {
		r.ParseForm() // if no this line, can't read form value
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		v := r.Form
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		fmt.Println("v.Get('name')", v.Get("name"))
		fmt.Println("v.Get('friend')", v.Get("friend"))
		fmt.Println("v['friend']", v["friend"])
	} else {
		fmt.Fprintf(w, "unknown method")
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
