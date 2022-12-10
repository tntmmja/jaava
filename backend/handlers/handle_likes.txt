package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func likeHandler(w http.ResponseWriter, r *http.Request) {

	var user *UserRecord
	user = find_logged_in_user(r)
	if nil == user {
		fmt.Println("   user not logged in")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	like_param := r.URL.Query().Get("like")
	like := ""

	if like_param == "Like" {
		like = "L"
	} else if like_param == "Dislike" {
		like = "D"
	} else {
		return
	}

	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {

		http.Error(w, "500 internal server error", http.StatusInternalServerError)

		return
	}
	comment_id, err := strconv.Atoi(r.URL.Query().Get("id_comm"))

	if err != nil {
		comment_id = 0
		//return
	}

	fmt.Println("likehandler", like, post_id, comment_id)
	//foorum_dao.update_like(user.Id, post_id, -1, like)
	foorum_dao.update_like(user.Id, post_id, comment_id, like)
	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)

}

func (dao *FoorumDao) update_like(user_id int, post_id int, comment_id int, like string) {
	//dao.db.Query("INSERT INTO (title, text) VALUES (?, ?)")
	// kontrollin kas kasutaja on juba laikinud
	// kui kasutaja on postitust laikinud, siis uut like id-d ei tee,
	// kui on samasugune laik, nagu enne oli, siis kustutab laiki, kui teeb vastandlaiki, siis update
	checkLike, err := dao.db.Query("SELECT id, likes FROM likes WHERE user_id = ? AND post_id = ? AND comment_id = ?", user_id, post_id, comment_id)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var existingLike string
	existingId := -1
	{

		defer checkLike.Close()
		if checkLike.Next() {

			err = checkLike.Scan(&existingId, &existingLike)
			if err != nil {
				panic(err.Error())
			}
		}

	}
	checkLike.Close()
	fmt.Println("rida70, existinglike ja existingid", existingLike, existingId)
	if existingId >= 0 {

		if like == existingLike {
			// to do delete

			fmt.Println("enne deletestatementi")
			deletestmt, err := dao.db.Prepare("DELETE FROM likes WHERE id = ?")
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			defer deletestmt.Close()
			fmt.Println("enne ja after vahel deletestatementi")
			_, err = deletestmt.Exec(existingId)
			fmt.Println("after execit")
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			fmt.Println("after deletestatementi")
		} else {

			// to do update

			updtstmt, err := dao.db.Prepare("UPDATE likes SET likes = ? WHERE id = ?")
			if err != nil {

				log.Fatal(err)
				panic(err)
			}
			defer updtstmt.Close()

			_, err = updtstmt.Exec(like, existingId)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}

		}
		return
	}

	stmt, err := dao.db.Prepare("INSERT INTO likes(user_id, post_id, comment_id, likes) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, post_id, comment_id, like)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}

/*
//3mai
type UserLikes struct{
LikeId int
UserId int
PostId int
CommId int
Like string
}


func (dao *FoorumDao) my_likes(user *UserRecord) UserLikes {

	user = foorum_dao.find_user_record_for_session()
	likerows, err := dao.db.Query("SELECT id, user_id, post_id, comment_id, likes FROM likes WHERE user_id = ? AND likes = ?", user.id, Like)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("handlelikes mylikes", myLikes)

}
*/
