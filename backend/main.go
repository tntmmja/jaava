package main

import (

	//	"github.com/gorilla/context"

	"fmt"
	"log"
	"net/http"

	"github.com/tntmmja/jaava/config"
	"github.com/tntmmja/jaava/data"
	"github.com/tntmmja/jaava/handlers"
	// _ "github.com/google/uuid"

	//	_ "github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	//"handlers"
)

// var tpl = template.Must(template.ParseGlob("templates/*.html"))
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("intekshandler ", r.RequestURI)

	// if r.URL.Path != "/" {
	// 	http.Error(w, "404 address NOT FOUND", http.StatusNotFound)
	// 	return
	// }
	fmt.Println("indexhandler")
	fmt.Fprintf(w, "testing backend to frontend")
	//tpl.ExecuteTemplate(w, "index.html", nil)
}
func LoggedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./dist/logged-user.html")
  }

func main() {
	
	// Create a new mux router. r on see router
	//https://www.youtube.com/watch?v=1E_YycpCsXw
	//seal on func main selline
	r := mux.NewRouter()
	SetRoutes(r)


	http.Handle("/", r) // Handle registers a new route with a matcher for the URL path. See Route.Path() and Route.Handler().
	config.DBConn()

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8083"},
	})
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe("localhost:8082", handler))
	/// miks siis kui config.DBConn on enne log.Fatal
	//siis tuleb "DB Connected" aga kui siin viimasena, siis ei tule?

}

// see voiks ka routes package all olla, aga praegu on siin
var SetRoutes = func(router *mux.Router) {
	//https://www.youtube.com/watch?v=1E_YycpCsXw
	//route teeb 16 minutil
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/register", data.RegisterHandler) // siia
	router.HandleFunc("/login", handlers.LoginHandler)
	// router.HandleFunc("/logouth", handlers.logoutHandler)
	router.HandleFunc("/loggedUser", handlers.LoggedHandler)
	// router.HandleFunc("/add_post", handlers.addPostHandler)
	// router.HandleFunc("/likes", handlers.likeHandler)
	// router.HandleFunc("/comment", handlers.addCommentHandler)
}

// https://levelup.gitconnected.com/experiment-golang-http-builtin-and-related-popular-packages-1d9a6dcb80d
// With gorilla/mux, we can declare complex routes with variables,
// constrain routes with methods, etc.
// r.HandleFunc("/users/", listUsers).Methods(http.MethodGet)  voi Methods("GET")
// r.HandleFunc("/users/", createUser).Methods(http.MethodPost)
// r.HandleFunc("/users/{userId}/", getUser).Methods(http.MethodGet)
// r.HandleFunc("/users/{userId}/", updateUser).Methods(http.MethodPut)
// r.HandleFunc("/users/{userId}/", deleteUser).Methods(Http.MethodDelete)

//Get captured variables
//We declared routes with variables, we should be able to
// capture the values of those variables in handlers.
// func handler(w http.ResponseWriter, r *http.Request) {
//      // mux.Vars(r) returns all values captured in the request URL.
//      vars := mux.Vars(r)
//      // vars is a dictionary whose key-value pairs are variables' name-value pairs.
//      fmt.Fprintf(w, "User %s\n", vars["userId"])
// }
