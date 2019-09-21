package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"

	"./pkg/login"
	"./pkg/register"
)

type Person struct {
	UserName string
}

func sayYourName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.Form[name]", r.Form["name"])
	var Name string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		Name = strings.Join(v, ",")
	}
	fmt.Println(Name)
	fmt.Fprintf(w, Name)
}

func showUserTopPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println("Cookie :", err)
		return
	}

	if r.Method == "GET" {
		fmt.Println(cookie)
		userName, err := r.Cookie("UserName")
		if err != nil {
			fmt.Println(err)
		}

		p := Person{UserName: userName.Value}
		t, _ := template.ParseFiles("./views/top.gtpl")
		t.Execute(w, p)

	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", sayYourName)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/new", register.NewUserRegister)
	http.HandleFunc("/top", showUserTopPage)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
