package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

type DashboardData struct {
	LoggedIn  bool
	FirstName string
	LastName  string
	Posts     []Post
	Filter    PostFilter
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dasbordhandler")

	session, err := r.Cookie("mycookie")
	var sessionID string
	if err == nil {
		sessionID = session.Value
	}
	fmt.Println("dhandler sessionid:", sessionID)

	filter := PostFilter{}

	my_post := r.FormValue("mypost")
	fmt.Println("dashboard handle minupost", my_post)
	
	// if err != nil {
	// 	fmt.Println("cookie not found")
	// 	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	// 	return
	// }
	// sessionID := session.Value
	// db := dbConn()
	// check, err := db.Query("select id from user where sessionID = ?", sessionID)
	var id int
	db := dbConn()
	_, err = db.Query("select text from posts where id = ?", id)
	if err != nil {

		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	var user *UserRecord
	user = foorum_dao.find_user_record_for_session(sessionID)
	if my_post == "on" && nil != user {

		filter.CreatorUserId = user.Id
	}

	my_like := r.FormValue("mylike")                  //3mai
	fmt.Println("dashboard handle minulaik", my_like) //3mai

	//var likes *UserLikes //3mai

	if my_like == "on" && nil != user { //3mai
		//	fmt.Println("Testime", likes.Like)

		filter.LikedByUserId = user.Id //3mai
		fmt.Println("handledashboard filterliked", filter.LikedByUserId)

	}
	//likes = foorum_dao.my_likes(like_id, user_id, post_id, comment_id, like) //3mai

	filter.Category = strings.Fields(r.FormValue("category"))

	dashboardData := DashboardData{}
	if nil != user {
		dashboardData.LoggedIn = true
		dashboardData.FirstName = user.FirstName
		dashboardData.LastName = user.LastName
	}
	dashboardData.Posts = foorum_dao.list_posts(filter)
	dashboardData.Filter = filter
	tpl.ExecuteTemplate(w, "dashboard.html", dashboardData)

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer check.Close()
	// if check.Next() {
	// 	var id int
	// 	check.Scan(&id)
	// 	fmt.Println("id", id)

	// 	dashboardData := DashboardData{}
	// 	dashboardData.FirstName = "test"
	// 	dashboardData.LastName = "SN"
	// 	tpl.ExecuteTemplate(w, "dashboard.html", dashboardData)

	// } else {
	// 	fmt.Println("session not found")
	// 	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	// }

}
