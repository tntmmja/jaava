package handlers

import (

	//"text/template"

	//"github.com/google/uuid"
	//"golang.org/x/crypto/bcrypt"

	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"github.com/tntmmja/jaava/config"
	"github.com/tntmmja/jaava/data"
	"github.com/google/uuid"

	//"text/template"

	//"github.com/dgrijalva/jwt-go"
	//"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//var tpl = template.Must(template.ParseGlob("templates/*.html"))

type Login struct {
	ID       int
	Username string `json:"username"`

	Password  string `json:"password"`
	sessionID string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginhandler")
	if r.Method == "POST" {
		db := config.GetDB()
		// username := r.FormValue("username")
		// password := r.FormValue("password")
		var login Login
		//log.Println(r.Body)
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal([]byte(b), &login)

		fmt.Printf("%s, %s\n", login.Username, login.Password)

		if strings.Trim(login.Username, " ") == "" || strings.Trim(login.Password, " ") == "" {
			fmt.Println("Parameter's can't be empty")
			//tpl.ExecuteTemplate(w, "loginwrong.html", nil)
			//http.Redirect(w, r, "/loginparametersempty", http.StatusMovedPermanently)
			return
		}
		fmt.Println("varchechuserini")
		var checkUser *sql.Rows
		//var err error

		if strings.Contains(login.Username, "@") {
			checkUser, err = db.Query("SELECT id, password, nickname, email FROM user WHERE email=?", login.Username)
		} else {
			checkUser, err = db.Query("SELECT id, password, nickname, email FROM user WHERE nickname=?", login.Username)
		}

		if err != nil {
			//tpl.ExecuteTemplate(w, "loginwrong.html", nil)
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
			//(err.Error())
		}

		fmt.Println("chekuser", checkUser)
		//fmt.Println("checking", fmt.Sprintf("%v", checkUser))
		//fmt.Printf("%s\n", checkUser)

		defer checkUser.Close()
		user := &data.User{}
		for checkUser.Next() {
			var id int
			var password, nickName, email string
			//var createdDate time.Time
			err = checkUser.Scan(&id, &password, &nickName, &email)
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
				//panic(err.Error())
			}
			user.ID = id
			user.Nickname = nickName
			user.Email = email
			user.Password = password
			//user.CreatedDate = createdDate
		}

		if user.ID == 0 {
			
		}
		fmt.Println("loginni cek", user.ID, user.Nickname, user.Email, user.Password)
		
		errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			fmt.Println("loginni password ei sobi")
			fmt.Println(errf)

			
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
			login.sessionID = sessionID
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
				//panic(err.Error())
			}

			//w.Header().Add("Set-Cookie", "mycookie="+sessionID+"; Max-Age = 300")
			fmt.Println("suunaloggeduser")
			
			return
		}
	} else if r.Method == "GET" {
		
	}
}
