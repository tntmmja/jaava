package handlers

import (
	"fmt"
	"net/http"

	"github.com/tntmmja/jaava/config"
	//"github.com/tntmmja/jaava/data"
	
	//"github.com/google/uuid"
	//"text/template"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
)

type DashboardData struct {
	LoggedIn  bool
	FirstName string
	LastName  string
	// Posts     []Post
	// Filter    PostFilter
}

func LoggedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loggedhandler")
	//http.ServeFile(w, r, "./dist/logged-user.html")
	session, err := r.Cookie("mycookie")
	var sessionID string
	if err == nil {
		sessionID = session.Value
	}
	fmt.Println("dhandler sessionid:", sessionID)

	//filter := PostFilter{}

	my_post := r.FormValue("mypost")
	fmt.Println("dashboard handle minupost", my_post)

	var id int
	db := config.GetDB()

	_, err = db.Query("select text from posts where id = ?", id)
	if err != nil {

		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	
}
