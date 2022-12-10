package handlers

import (
	"fmt"
	"net/http"
	"strings"

	//"text/template"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//var tpl = template.Must(template.ParseGlob("templates/*.html"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginhandler")
	if r.Method == "POST" {
		db := dbConn()
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Printf("%s, %s\n", email, password)

		if strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
			fmt.Println("Parameter's can't be empty")
			tpl.ExecuteTemplate(w, "loginwrong.html", nil)
			//http.Redirect(w, r, "/loginparametersempty", http.StatusMovedPermanently)
			return
		}

		checkUser, err := db.Query("SELECT id, password, firstname, lastname, email FROM user WHERE email=?", email)
		if err != nil {
			//tpl.ExecuteTemplate(w, "loginwrong.html", nil)
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
			//(err.Error())
		}

		fmt.Println("chekuser", checkUser)
		defer checkUser.Close()
		user := &User{}
		for checkUser.Next() {
			var id int
			var password, firstName, lastName, email string
			//var createdDate time.Time
			err = checkUser.Scan(&id, &password, &firstName, &lastName, &email)
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
				//panic(err.Error())
			}
			user.ID = id
			user.FirstName = firstName
			user.LastName = lastName
			user.Email = email
			user.Password = password
			//user.CreatedDate = createdDate
		}

		if user.ID == 0 {
			tpl.ExecuteTemplate(w, "login.html", nil)
		}
		fmt.Println("loginni cek", user.ID, user.Email)
		errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			fmt.Println(errf)

			//http.Redirect(w, r, "/loginpasswordnotmatch", http.StatusTemporaryRedirect)
			tpl.ExecuteTemplate(w, "loginwrong.html", nil)
		} else {
			sessionID := uuid.New().String() //generates random text
			fmt.Println("loginsessionid1", sessionID)

			upt, err := db.Prepare("update user set sessionID = ? where id = ?")

			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
				//panic(err.Error())
			}
			defer upt.Close()
			_, err = upt.Exec(sessionID, user.ID)
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
				//panic(err.Error())
			}

			w.Header().Add("Set-Cookie", "mycookie="+sessionID+"; Max-Age = 300")

			http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
			return
		}
	} else if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "login.html", nil)
	}
}
