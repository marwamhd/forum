package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func (DataBase *DB) WhatUserLiked(uID, pID int) (int, error) {
	var liked int

	// Check if the user liked the post
	err := DataBase.DB.QueryRow("SELECT liked FROM author_liked_post WHERE p_id = ? AND u_id = ?", pID, uID).Scan(&liked)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 2, nil
		}
		return -1, err
	}
	return liked, nil
}

func (DataBase *DB) WhatUserLikedComment(uID, pID, cID int) (int, error) {
	var liked int

	// Check if the user liked the post
	err := DataBase.DB.QueryRow("SELECT liked FROM author_liked_comment WHERE p_id = ? AND u_id = ? AND c_id = ?", pID, uID, cID).Scan(&liked)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 2, nil
		}
		return -1, err
	}
	return liked, nil
}

func (DataBase *DB) WhatUserLikedPosts(uID int) ([]Post, error) {
	var posts []Post

	rows, err := DataBase.DB.Query("SELECT p_id FROM author_liked_post WHERE u_id = ? AND liked = 1", uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var postids int
		err := rows.Scan(&postids)
		if err != nil {
			return nil, err
		}
		prows, err := DataBase.DB.Query("SELECT * FROM posts WHERE id = ?", postids)
		if err != nil {
			return nil, err
		}
		defer prows.Close()

		for prows.Next() {
			var post Post
			err = prows.Scan(&post.ID, &post.U_ID, &post.Title, &post.Post, &post.Cat1, &post.Cat2, &post.Cat3)
			if err != nil {
				return nil, err
			}

			// Fetch username
			post.Username, err = getUsername(post.U_ID)
			if err != nil {
				return nil, err
			}

			// Query to get comments for the post
			crows, err := DataBase.DB.Query("SELECT * FROM comments WHERE p_ID = ?", post.ID)
			if err != nil {
				return nil, err
			}
			defer crows.Close()

			// Fetch comments
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

			likes, dislikes, err := DataBase.LikesDislikesTotal(strconv.Itoa(post.ID))
			if err != nil {
				return nil, err
			}

			post.Likes = likes
			post.Dislikes = dislikes

			posts = append(posts, post)
		}
	}
	return posts, nil
}
