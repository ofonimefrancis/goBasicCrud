package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
var store = sessions.NewCookieStore([]byte("youhadgot2bekiddingme"))
var err error
var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root@glootian@/status")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router.HandleFunc("/", indexHandler)
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

//TODO: Analyze for the appropriate arguments to this function
func InsertPost(post Post) {
	err := db.Ping() //make sure that the database is available
	if err != nil {
		//Do something here
	}
}

//TODO: Analyze for the appropriate arguments to this function
func DeletePost(post Post) {
	err := db.Ping() //make sure that the database is available
	if err != nil {
		//Do something here
	}
}

func UpdatePost(postId int64) {
	err := db.Ping()
	if err != nil {
		//Do something here
	}
}

func ListAllPost() {
	err := db.Ping()
	if err != nil {
		//Do something here
	}
}

func getPostByID() {
	err := db.Ping()
	if err != nil {
		//Do something here
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func setSession(w http.ResponseWriter, r *http.Request, username string, password string) {
	session, err := store.Get(r, "logged-in")
	handleError(w, r, err)

	if !session.IsNew {
		return
	}
	bytePassword := []byte(password)
	hasher := sha1.New()
	hasher.Write(bytePassword)
	encryptedPassword := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	session.Values["username"] = username
	session.Values["password"] = encryptedPassword //Do not set password session if u wont atleast hash it
	session.Values["loggedin"] = true
	session.Save(r, w)
}

func invalidateSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "logged-in")
	handleError(w, r, err)
	if !session.IsNew {
		session.Values["username"] = " "
		session.Values["password"] = " "
		session.Values["loggedin"] = false
		session.Save(r, w)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", "Login", User{UserId: 1, Username: "Jiggaseige"})
}

func renderTemplate(w http.ResponseWriter, tmpl string, name string, user interface{}) {
	t := template.Must(template.New("fb").ParseFiles("./templates/" + tmpl + ".html"))
	if err := t.ExecuteTemplate(w, name, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "logged-in")

	//username := session.Values["username"]

	//We want to show the
	//1. Welcome Username here
	//2. Loop through all the post added in the database
	renderTemplate(w, "home", "Home", User{Username: "googlehead"})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	invalidateSession(w, r)
	http.Redirect(w, r, "/", 302)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := User{}
	if len(username) > 0 && len(password) > 0 {
		user.Username = username
		user.Password = password
		cookie := http.Cookie{Name: "username", Value: username}
		setSession(w, r, username, password)
		http.SetCookie(w, &cookie)
		//TODO: save to the database
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	} else {
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
