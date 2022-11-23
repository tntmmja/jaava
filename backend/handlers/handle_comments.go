package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func addCommentHandler(w http.ResponseWriter, r *http.Request) {

	var user *UserRecord
	user = find_logged_in_user(r)
	if nil == user {
		fmt.Println("   user not logged in")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	//comment := r.URL.Query().Get("post_comment")
	comment := r.FormValue("post_comment")
	fmt.Println("kommentaar", comment)
	//post_id, err := strconv.Atoi(r.URL.Query().Get("post_id"))

	post_id, err := strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		fmt.Println("kommentierror", err)
		return
	}

	fmt.Println("komment", comment, post_id)

	if len(strings.TrimSpace(comment)) > 1 {
		// checks if comment field does not contain any bad requests
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "400 bad request", http.StatusBadRequest)
			return

		}

		foorum_dao.insert_db_comment(user.Id, post_id, comment)

	}
	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)

}

// kirjutab veebilehelt kommentaari andmebaasi
func (dao *FoorumDao) insert_db_comment(user_id int, post_id int, comment string) {

	stmt, err := dao.db.Prepare("INSERT INTO comments(user_id, post_id, text) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, post_id, comment)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// /// 20.04 lisatud

type Comment struct {
	CommentId        int
	PostId           int
	UserId           int
	UserEmail        string
	CreatedAt        time.Time
	Text             string
	NumberOfLikes    int
	NumberOfDislikes int
}

func (dao *FoorumDao) list_comments(postID int) []Comment {
	rows, err := dao.db.Query(`
	SELECT comments.id, comments.user_id, comments.post_id, comments.text, 
		comments.created_at,
		COALESCE(user.email, '') AS user_email,
		(SELECT COUNT (1) FROM likes WHERE likes.post_id = comments.post_id AND likes.comment_id = comments.id AND likes.likes = 'L') AS num_likes,
		(SELECT COUNT (1) FROM likes WHERE likes.post_id = comments.post_id AND likes.comment_id = comments.id AND likes.likes = 'D') AS num_dislikes
	
	FROM comments
		LEFT JOIN user ON comments.user_id = user.id
	WHERE post_id = ?
		
 ORDER BY comments.created_at DESC
		`, postID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		comment := Comment{}
		err := rows.Scan(
			&(comment.CommentId),
			&(comment.UserId),
			&(comment.PostId),
			&(comment.Text),
			&(comment.CreatedAt),
			&(comment.UserEmail),
			&(comment.NumberOfLikes),
			&(comment.NumberOfDislikes),
		)
		if err != nil {
			panic(err)
		}

		// Leia likeid ja dislikeid

		comments = append(comments, comment)
	}
	return comments
}
