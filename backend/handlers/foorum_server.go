package handlers

import (
	"fmt"
	//"real-time-forum/backend/main"
	"net/http"
	"strings"
	//"encoding/json"
	//"fmt"
	//	"html"
	//	"io/ioutil"
	//	"log"
	//"net/http"
)

type Server struct {
	dao *FoorumDao
}

// server starts http server on specified port
func (s *Server) inithandlers() {

	//	http.Handle("/api/list_posts", s.handleListPost())
	//	http.Handle("/api/add_post", s.handleAddPost())
}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addPostHandler")
	fmt.Println("GET params were:", r.URL.Query())
	title := r.URL.Query().Get("post_title")
	post_text := r.URL.Query().Get("post_text")

	var user *UserRecord
	user = find_logged_in_user(r)
	if nil == user {
		fmt.Println("   user not logged in")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("   logged in user: " + user.EMail)
	fmt.Println("   title: " + title)
	fmt.Println("   post: " + post_text)
	if len(strings.TrimSpace(title)) > 1 && len(strings.TrimSpace(post_text)) > 5 {
		foorum_dao.insert_post(user.Id, title, post_text)
	}

	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
}

func find_session_cookie(r *http.Request) string {
	session, err := r.Cookie("mycookie")
	var sessionID string
	if err == nil {
		sessionID = session.Value
	}
	return sessionID
}

func find_logged_in_user(r *http.Request) *UserRecord {
	sessionID := find_session_cookie(r)
	return foorum_dao.find_user_record_for_session(sessionID)
}
