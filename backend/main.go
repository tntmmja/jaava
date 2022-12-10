package main

import (

	// "real-time-forum/backend/handlers"
	// "real-time-forum/backend/config"

	//"strings"

	//	"github.com/gorilla/context"

	"fmt"
	"log"
	"net/http"

	// "01.kood.tech/git/hr.tauno/real-time-forum/data"
	// _ "github.com/google/uuid"

	//	_ "github.com/joho/godotenv"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	//"handlers"
)

//var tpl = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("intekshandler ", r.RequestURI)

	// if r.URL.Path != "/" {
	// 	http.Error(w, "404 address NOT FOUND", http.StatusNotFound)
	// 	return
	// }
	fmt.Fprintf(w, "testing backend to frontend")
	//tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/login", handlers.loginHandler)
	// http.HandleFunc("/logouth", logoutHandler)
	// http.HandleFunc("/register", registerHandler)
	// http.HandleFunc("/dashboard", dashboardHandler)
	// http.HandleFunc("/add_post", addPostHandler)
	// http.HandleFunc("/likes", likeHandler)
	// http.HandleFunc("/comment", addCommentHandler)
	//http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("static"))))
	r := mux.NewRouter()
	
	// config.DBConn()
	// http.HandleFunc("/", indexHandler)
	data.RegisterBookstoreRoutes()

	// data.registerHandler(r)


	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}

//https://www.youtube.com/watch?v=1E_YycpCsXw
//seal on func main selline
/*func main() {
	r := mux.NewRouter() //github.com/gorilla/mux on imporditud
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r) //HandleFunc-id olid juba package routes all uhe variabli RegisterBookstoreRoutes alla koondatud
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}*/
