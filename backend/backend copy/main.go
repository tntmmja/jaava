package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	//"real-time-forum/backend/handlers"

	//"strings"

	"text/template"

	//	"github.com/gorilla/context"

	_ "github.com/google/uuid"
	//	_ "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int
	FirstName string `json:"firstname" validate:"required, gte=3"`
	LastName  string `json:"lastname" validate:"required, gte=3"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	//CreatedDate time.Time `json:"createdDate"`
	Cookie string
}

func dbConn() (db *sql.DB) {
	

	db, err := sql.Open("sqlite3", "rltforum.db")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB Connected!!")
	return db
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("intekshandler ", r.RequestURI)

	if r.URL.Path != "/"{
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

func setRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/", backend.IndexHandler)
	mux.HandleFunc("/login", backend.LoginHandler)
	mux.HandleFunc("/reg", backend.RegHandler)
	mux.HandleFunc("/logout", backend.DeleteSession)
	mux.HandleFunc("/newpost", backend.NewPostHandler)
	mux.HandleFunc("/getcomments", backend.GetCommentsHandler)
	mux.HandleFunc("/getallusers", backend.GetAllUsersHandler)
	mux.HandleFunc("/getonline", backend.GetOnlineHandler)

	//mux.HandleFunc("/ws", backend.HandleConnections)
	mux.HandleFunc("/ws", backend.WebsocketHandler)

	// API
	mux.HandleFunc("/allposts", backend.GetAllPosts)
	mux.HandleFunc("/post", backend.GetPostAndComments)
	mux.HandleFunc("/commentpost", backend.CreatePostComment)
	mux.HandleFunc("/userauth", backend.IsUserAuthenticated)
	mux.HandleFunc("/getotherusers", backend.GetOtherUsers)
	mux.HandleFunc("/getlast10messages", backend.GetLast10Messages)

	fs := http.FileServer(http.Dir("./frontend"))
	//go backend.HandleMessages()
	mux.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	return mux
}


/////////////////////////////////////////////////////////////////////////
