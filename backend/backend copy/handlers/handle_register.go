package handlers

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registerhandler")
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

		//dt := time.Now()

		//createdDateString := dt.Format("2006-01-02 15:04:05")

		// Convert the time before inserting into the database
		//createdDate, err := time.Parse("2006-01-02 15:04:05", createdDateString)
		//if err != nil {
		//	log.Fatal("Error converting the time:", err)
		//	}

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

			//fmt.Println("Error when inserting: ", err.Error())
			//tpl.ExecuteTemplate(w, "registertaken.html", nil)
			//panic(err.Error())
		}
		log.Println("=> Inserted: First Name: " + firstName + " | Last Name: " + lastName)

		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	} else if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "register.html", nil)
	}
}
