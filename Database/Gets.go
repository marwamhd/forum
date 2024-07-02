package database

import (
	"fmt"
	"strconv"
)

func (Database DB) GetPosts() ([]Post, error) {
	rows, err := Database.DB.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.U_ID, &post.Title, &post.Post, &post.Cat1, &post.Cat2, &post.Cat3)
		if err != nil {
			return nil, err
		}

		post.Username, err = getUsername(post.U_ID)
		if err != nil {
			return nil, err
		}

		crows, err := Database.DB.Query("SELECT * FROM comments WHERE p_ID = ?", post.ID)
		if err != nil {
			return nil, err
		}
		defer crows.Close()

		var comments []Comment
		for crows.Next() {
			var comment Comment
			err = crows.Scan(&comment.ID, &comment.U_ID, &comment.P_ID, &comment.Comment)
			if err != nil {
				return nil, err
			}
			comments = append(comments, comment)
		}

		post.Comments = comments
		posts = append(posts, post)
	}

	return posts, nil
}

func (Database DB) GetFilteredPosts(str string) ([]Post, error) {
	rows, err := Database.DB.Query(str)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.U_ID, &post.Title, &post.Post, &post.Cat1, &post.Cat2, &post.Cat3)
		if err != nil {
			return nil, err
		}

		post.Username, err = getUsername(post.U_ID)
		if err != nil {
			return nil, err
		}

		crows, err := Database.DB.Query("SELECT * FROM comments WHERE p_ID = ?", post.ID)
		if err != nil {
			return nil, err
		}
		defer crows.Close()

		var comments []Comment
		for crows.Next() {
			var comment Comment
			err = crows.Scan(&comment.ID, &comment.U_ID, &comment.P_ID, &comment.Comment)
			if err != nil {
				return nil, err
			}
			likes, disliks, err := DataBase.CommentLikesDislikesTotal(strconv.Itoa(post.ID), strconv.Itoa(comment.ID))
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return nil, err
			}

			comment.Likes = likes
			comment.Dislikes = disliks
			comments = append(comments, comment)
		}

		post.Comments = comments

		likes, disliks, err := DataBase.LikesDislikesTotal(strconv.Itoa(post.ID))
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return nil, err
		}

		post.Likes = likes
		post.Dislikes = disliks

		posts = append(posts, post)
	}

	return posts, nil
}

func getUsername(id int) (string, error) {
	statement, err := DataBase.DB.Prepare("SELECT username FROM users WHERE id = ?")
	if err != nil {
		return "", err
	}
	defer statement.Close()

	var username string
	err = statement.QueryRow(id).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

func GetAuthor(session string) (int, error) {
	statement, err := DataBase.DB.Prepare("SELECT id FROM users WHERE session_id = ?")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(session).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
