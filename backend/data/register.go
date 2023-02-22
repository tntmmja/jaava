package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/tntmmja/jaava/config"
)

type User struct {
	ID        int
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	sessionID int
}

var db *sql.DB

func init() {
	config.DBConn()
	db = config.GetDB()
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registerhandler")
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// 	fmt.Println("registerhandleris on viga")
	// 	return
	// }
	fmt.Println("registerhandler2")

	// age, err := strconv.Atoi(r.FormValue("age"))
	// if err != nil {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }

	// ageStr := r.FormValue("age")
	// if _, err := strconv.Atoi(ageStr); err != nil {
	//   http.Error(w, "Bad Request", http.StatusBadRequest)
	//   return
	// }
	// age, err := strconv.Atoi(ageStr)

	// kui teha ageStr ja user:=User, siis ei toimi, tuleb ainult age 0 ja teised on tuhjad
	// ageStr := r.FormValue("age")
	// age, err := strconv.Atoi(ageStr)
	// if err != nil {
	// 	age = 0
	// }

	// user := User{
	// 	Nickname:  r.FormValue("nickname"),
	// 	Age:       age,
	// 	Gender:    r.FormValue("gender"),
	// 	FirstName: r.FormValue("firstName"),
	// 	LastName:  r.FormValue("lastName"),
	// 	Email:     r.FormValue("email"),
	// 	Password:  r.FormValue("password"),
	// }

	var user User

	log.Println(r.Body)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal([]byte(b), &user)

	fmt.Printf("%d, %s, %s, %s\n", user.Age, user.FirstName, user.LastName, user.Email)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	fmt.Println("krupteerimine")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	db := config.GetDB()
	defer db.Close()
	fmt.Println("hakka uuser tabelit prepeerima")
	stmt, err := db.Prepare("INSERT INTO user (nickname, age, gender, firstName, lastName, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	fmt.Println("hakka uuser tabelit ekskuutima")
	_, err = stmt.Exec(user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, hash)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// seda registreerimisel pole vaja, teeb logini
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "sessionID",
	// 	Value:   sessionID,
	// 	Expires: time.Now().Add(24 * time.Hour),
	// })

	//http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 	if r.Method == "POST" {
// 		firstName := r.FormValue("FirstName")
// 		lastName := r.FormValue("LastName")
// 		email := r.FormValue("email")
// 		fmt.Printf("%s, %s, %s\n", firstName, lastName, email)

// 		password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
// 		if err != nil {
// 			fmt.Println(err)
// 			// Handle error
// 		}

// 		_, err = db.Exec("INSERT INTO user(firstname, lastname,email,password) VALUES(?,?,?,?)", firstName, lastName, email, password)
// 		if err != nil {
// 			fmt.Println("Error when inserting: ", err.Error())
// 			if err.Error() == "UNIQUE constraint failed: user.email" {
// 				// Handle error: email already exists
// 			} else {
// 				http.Error(w, "500 internal server error", http.StatusInternalServerError)
// 				return
// 			}
// 		}
// 		log.Println("=> Inserted: First Name: " + firstName + " | Last Name: " + lastName)

// 		// Redirect to login page
// 		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
// 	} else {
// 		// Render registration template or handle method not allowed error
// 	}

// func generateSessionID() string {
// 	// implementation to generate a unique session ID
// }
