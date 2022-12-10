package handlers

import (
	"fmt"
	"net/http"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logoutHandler")
	session, err := r.Cookie("mycookie")

	if err != nil {
		return
	}
	sessionID := session.Value
	fmt.Println("sessionid1:", sessionID)
	db := dbConn()
	upt, err := db.Prepare("update user set sessionID = null where sessionID = ?")
	fmt.Println("sessionid2:", sessionID)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
		//panic(err.Error())
	}
	defer upt.Close()
	_, err = upt.Exec(sessionID)

	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect) //oli "/login"
}
