package data

import (
	"fmt"
	"log"
	"net/http"

	"01.kood.tech/git/hr.tauno/real-time-forum/config"
	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"
)

// see db on db_connect.go-s ka
var db *sql.DB

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

// is that the right place or should it be in db_connect.go
// vt https://www.youtube.com/watch?v=1E_YycpCsXw
// 28:46
func init() {
	
	db = config.dbConn()
}

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


// mux-i jaoks tuleb ikka responsewriteriga ka midagi teha
var RegisterHandler = func(router *mux.Route) {
	fmt.Println("registerhandler")
	router.HandleFunc("/register/", controllers.)Methods("POST")
	if r.Method == "POST" {
		db := dbConn()
		firstName := r.FormValue("FirstName")
		lastName := r.FormValue("LastName")
		email := r.FormValue("email")
		fmt.Printf("%s, %s, %s\n", firstName, lastName, email)

		password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			tpl.ExecuteTemplate(w, "Register", err)
		}

		_, err = db.Exec("INSERT INTO user(firstname, lastname,email,password) VALUES(?,?,?,?)", firstName, lastName, email, password)
		if err != nil {
			fmt.Println("Error when inserting: ", err.Error())
			if err.Error() == "UNIQUE constraint failed: user.email" {
				tpl.ExecuteTemplate(w, "registertaken.html", nil)
				//panic(err.Error())
			} else {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
			}

		}
		log.Println("=> Inserted: First Name: " + firstName + " | Last Name: " + lastName)

		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	} else if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "register.html", nil)
		
	}
}

