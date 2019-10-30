package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	//login "github.com/Snow-HardWolf/Vulnerability-goapp/pkg/login"

	"./pkg/cookie"
	uploader "./pkg/image"
	"./pkg/login"
	"./pkg/logout"
	"./pkg/post"
	"./pkg/register"
	"./pkg/user"
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
	userName, sessionID, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		fmt.Println(err)
	}

	if cookie.CheckSessionsCount(userID, sessionID) {
		login.StoreSID(userID, sessionID)
	} else {
		fmt.Println("not register sessionID")
	}

	if sessionID == "" {
		fmt.Println("sid not exist")
		t, _ := template.ParseFiles("./views/public/error.gtpl")
		t.Execute(w, nil)
	} else {

		if r.Method == "GET" {

			if userID != 0 {
				uid := strconv.Itoa(userID)
				cookieUserID := &http.Cookie{
					Name:  "UserID",
					Value: uid,
				}

				http.SetCookie(w, cookieUserID)
				p := Person{UserName: userName}
				t, _ := template.ParseFiles("./views/public/top.gtpl")
				t.Execute(w, p)
			}

		} else {
			http.NotFound(w, r)
		}
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	userName, sid, err := cookie.CheckCookieOnlyLogin(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userName, sid)
	_, _, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		fmt.Println(err)
	}
	if cookie.CheckSessionsCount(userID, sid) {
		fmt.Println("session count true")
	} else {
		fmt.Println("session count false")
	}

	if cookie.CheckSessionID(r) {
		fmt.Println("checkCookie true")
	} else {
		fmt.Println("checkCookie false")
	}

}

func Bonus(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/public/bonus.gtpl")
	t.Execute(w, nil)
}

func Hints(w http.ResponseWriter, r *http.Request) {

	sql := "mysql -h mysql -u root -prootwolf -e 'select * from vulnapp.user;'"
	//res, err := exec.Command("sh","-c","mysql", "-hmysql", "-uroot", "-prootwolf","-e","`show databases;`").Output()
	res, err := exec.Command("sh", "-c", sql).Output()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Println("res : ", string(res))

	t, _ := template.ParseFiles("./views/hints/hints.gtpl")
	t.Execute(w, nil)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", sayYourName)
	http.HandleFunc("/test", test)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/logout", logout.Logout)
	http.HandleFunc("/new", register.NewUserRegister)
	http.HandleFunc("/top", showUserTopPage)
	http.HandleFunc("/profile", user.ShowUserProfile)
	http.HandleFunc("/profile/edit", user.ShowUserModifyPage)
	http.HandleFunc("/profile/edit/confirm", user.ShowEditConfirm)
	http.HandleFunc("/profile/edit/update", user.UpdateUserDetails)
	http.HandleFunc("/profile/changepasswd", user.PasswdChange)
	http.HandleFunc("/profile/compchangepasswd", user.ConfirmPasswdChange)
	http.HandleFunc("/profile/edit/image", uploader.ShowImageChangePage)
	http.HandleFunc("/profile/edit/upload", uploader.UploadImage)
	http.HandleFunc("/post", post.ShowAddPostPage)
	http.HandleFunc("/timeline", post.ShowTimeline)
	http.HandleFunc("/hints", Hints)
	http.HandleFunc("/bonus", Bonus)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
