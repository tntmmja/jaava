package main

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/backend/handlers"
	"real-time-forum/backend/config"

	//"strings"
	"text/template"
	//	"github.com/gorilla/context"
	_ "github.com/google/uuid"
	//	_ "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	//"handlers"
)

type User struct {
	ID        int
	Nickname  string `json:"nickname" validate:"required, gte=3"`
	Age       int    `json: age`
	Gender    string `json:"gender" validate:"required, gte=3"`
	FirstName string `json:"firstname" validate:"required, gte=3"`
	LastName  string `json:"lastname" validate:"required, gte=3"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	//CreatedDate time.Time `json:"createdDate"`
	Cookie string // cookie was also used as session id in forum
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("intekshandler ", r.RequestURI)

	if r.URL.Path != "/" {
		http.Error(w, "404 address NOT FOUND", http.StatusNotFound)
		return
	}

	tpl.ExecuteTemplate(w, "index.html", nil)
}

var foorum_dao *FoorumDao

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", handlers.loginHandler)
	http.HandleFunc("/logouth", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/add_post", addPostHandler)
	http.HandleFunc("/likes", likeHandler)
	http.HandleFunc("/comment", addCommentHandler)
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("static"))))

	foorum_dao = &FoorumDao{dbConn()}
	server := Server{foorum_dao}
	server.inithandlers()

	log.Println("Server started on: http://localhost:8000")
	//err := http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)) // context to prevent memory leak
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("500 Internal server error", http.StatusInternalServerError) // internal server error
		return
	}
}


/////////////////////////////////////////////////////////////////////////
