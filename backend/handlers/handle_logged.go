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

	// var user *UserRecord
	// user = foorum_dao.find_user_record_for_session(sessionID)
	// if my_post == "on" && nil != user {

	// 	filter.CreatorUserId = user.Id
	// }

	// my_like := r.FormValue("mylike")                  //3mai
	// fmt.Println("dashboard handle minulaik", my_like) //3mai

	// if my_like == "on" && nil != user { //3mai

	// 	filter.LikedByUserId = user.Id //3mai
	// 	fmt.Println("handledashboard filterliked", filter.LikedByUserId)

	// }

	// filter.Category = strings.Fields(r.FormValue("category"))

	// dashboardData := DashboardData{}
	// if nil != user {
	// 	dashboardData.LoggedIn = true
	// 	dashboardData.FirstName = user.FirstName
	// 	dashboardData.LastName = user.LastName
	// }
	// dashboardData.Posts = foorum_dao.list_posts(filter)
	// dashboardData.Filter = filter
	// tpl.ExecuteTemplate(w, "dashboard.html", dashboardData)

}
