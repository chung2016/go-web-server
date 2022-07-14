package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// TODO: check token repeat submit
		} else {
			// TODO: no token exception
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //輸出到伺服器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println("----", t.Execute(w, nil))
	} else if r.Method == "POST" {
		r.ParseForm() // if no this line, can't read form value
		// fmt.Println("username:", r.Form["username"])
		// fmt.Println("password:", r.Form["password"])

		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintln(w, "username empty")
		}

		age := r.Form.Get("age")

		invalidAge := validateAgeReg(age)
		if !invalidAge {
			fmt.Fprintln(w, "age invalid")
		}
		username := r.Form.Get("username")
		if !isChineseName(username) {
			fmt.Fprintln(w, "is not chinese name")
		} else {
			fmt.Fprintln(w, "is chinese name")
		}
		t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))

		email := r.Form.Get("email")
		if !validEmail(email) {
			fmt.Fprintln(w, "is invalid email")
		} else {
			fmt.Fprintln(w, "valid email")
		}

		mobile := r.Form.Get("mobile")
		if m, _ := regexp.MatchString(`^([1-9]\d{7})$`, mobile); !m {
			fmt.Fprintln(w, "invalid mobile")
		} else {
			fmt.Fprintln(w, "valid mobile")
		}

		fruit := r.Form.Get("fruit")
		if validFruit(fruit) {
			fmt.Fprintln(w, "valid fruit")
		} else {
			fmt.Fprintln(w, "invalid fruit")
		}

		// t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		// fmt.Fprintf(w, "Go launched at %s\n", t.Local())

		v := r.Form
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		// fmt.Println("v.Get('name')", v.Get("name"))
		// fmt.Println("v.Get('friend')", v.Get("friend"))
		// fmt.Println("v['friend']", v["friend"])
		fmt.Fprintln(w, "form processed")
	} else {
		fmt.Fprintf(w, "unknown method")
	}
}

func validFruit(fruit string) bool {
	slice := []string{"apple", "pear", "banana"}
	for _, item := range slice {
		if item == fruit {
			return true
		}
	}
	return false
}

func validEmail(email string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	}
	return true
}

func isChineseName(name string) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", name); !m {
		return false
	}
	return true
}

func validateAge(age string) bool {
	if m, _ := regexp.MatchString("^[0-9]+$", age); !m {
		return false
	}
	return true
}
func validateAgeReg(age string) bool {
	_, err := strconv.Atoi(age)
	if err != nil {
		return false
	}
	return true
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/form", form)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
