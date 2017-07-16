package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	UserId   int64
	Username string
	Password string
	Posts    []*Post
}

type Post struct {
	StatusId int64
	Username string
	Status   string
}

var router = mux.NewRouter()

func main() {
	db, err := sql.Open("mysql", "root@glootian@/status")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/home", homeHandler)
	router.HandleFunc("/home/post", statusHandler).Methods("POST")
	router.HandleFunc("/home/update", updateHandler).Methods("GET")
	router.HandleFunc("/home/delete", removeHandler)
	router.HandleFunc("/home/save", saveHandler)

	http.Handle("/", router)
	errors := http.ListenAndServe(":8080", nil)
	if errors != nil {
		panic(err)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", "Login", User{UserId: 1, Username: "Jiggaseige"})
}

func renderTemplate(w http.ResponseWriter, tmpl string, name string, user User) {
	t := template.Must(template.New("fb").ParseFiles("./templates/" + tmpl + ".html"))
	if err := t.ExecuteTemplate(w, name, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", "Home", User{UserId: 1, Username: "Jiggaseige"})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := User{}
	if len(username) > 0 && len(password) > 0 {
		user.Username = username
		user.Password = password
		//Set user cookie
		//h := sha256.New()
		//bytesWritten, _ := h.Write([]byte(password))
		//hashed_password := sha256.Sum256([]byte(password))
		cookie := http.Cookie{Name: "username", Value: username}
		http.SetCookie(w, &cookie)

	} else {
		fmt.Fprintf(w, "Username and password cannot be empty")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "update", "Update", User{UserId: 1, Username: "Jiggaseige"})
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "status", "Status", User{UserId: 1, Username: "Jiggaseige"})
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "status", "Status", User{UserId: 1, Username: "Jiggaseige"})
}
