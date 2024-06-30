package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// DB wraps around sql.DB to provide methods
type DB struct {
	*sql.DB
}

type Post struct {
	ID       int
	U_ID     int
	Title    string
	Post     string
	Username string
	Cat1     int
	Cat2     int
	Cat3     int
	Comments []Comment
	Likes    int
	Dislikes int
}

type Comment struct {
	ID      int
	U_ID    int
	P_ID    int
	Comment string
}

var DataBase DB

func init() {
	var err error
	DataBase.DB, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	// Attempt to ping the database to verify connectivity
	err = DataBase.DB.Ping()
	if err != nil {
		log.Fatal("error pinging database:", err)
	}

	_, err = DataBase.DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}

	DataBase.CreateTable()

	// If no error occurred, the database connection is successfully established
	log.Println("Database connection established successfully")
}

func (DataBase *DB) CreateTable() {
	_, err := DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		session_id TEXT NULL
	)`)

	if err != nil {
		log.Fatal(err, "As")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		p_ID INTEGER NOT NULL,
		comment TEXT NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id),
		FOREIGN KEY (p_ID) REFERENCES posts(id)
	)`)

	if err != nil {
		log.Fatal(err, "s")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		title TEXT NOT NULL,
		post TEXT NOT NULL,
		cat1 INTEGER NOT NULL,
		cat2 INTEGER NOT NULL,
		cat3 INTEGER NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id)

	)`)

	if err != nil {
		log.Fatal(err, "a")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS author_liked_post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		p_ID INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id),
		FOREIGN KEY (p_ID) REFERENCES posts(id)
	)`)

	if err != nil {
		log.Fatal(err, "a")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS author_liked_comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		c_ID INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id),
		FOREIGN KEY (p_ID) REFERENCES posts(id),
		FOREIGN KEY (c_ID) REFERENCES comments(id)
	)`)

	if err != nil {
		log.Fatal(err, "a")
	}
}

// InsertUser inserts a user into the database
func (DataBase *DB) InsertUser(email, username, password string) error {
	statement, err := DataBase.DB.Prepare("INSERT INTO users (email, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(email, username, password)
	if err != nil {
		return err
	}
	return nil
}

func (DataBase *DB) EmailQuery() error {
	rows, err := DataBase.DB.Query("select * from users")
	if err != nil {
		return err
	}
	var id int
	var email string
	var username string
	var password string

	for rows.Next() {
		err = rows.Scan(&id, &email, &username, &password)
		if err != nil {
			return err
		}
		fmt.Println(id)
		fmt.Println(email)
		fmt.Println(username)
		fmt.Println(password)
	}

	rows.Close()

	return nil
}

func (DataBase *DB) EmailExists(InputEmail string) (bool, error) {
	statement, err := DataBase.DB.Prepare("SELECT id FROM users WHERE email = ?")
	if err != nil {
		return false, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(InputEmail).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (DataBase *DB) SessionExists(session string) (bool, error) {
	statement, err := DataBase.DB.Prepare("SELECT id FROM users WHERE session_id = ?")
	if err != nil {
		return false, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(session).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (DataBase *DB) Login(InputEmail, inputPassword string, r *http.Request) (int, uuid.UUID, error) {
	statement, err := DataBase.DB.Prepare("SELECT id, email, username, password, session_id FROM users WHERE email = ?")
	if err != nil {
		return 0, uuid.Nil, err
	}
	defer statement.Close()

	var id int
	var email string
	var username string
	var password string
	var sessionID sql.NullString

	row := statement.QueryRow(InputEmail)
	err = row.Scan(&id, &email, &username, &password, &sessionID)
	if err != nil {
		return 0, uuid.Nil, err
	}

	if inputPassword != password {
		return 0, uuid.Nil, fmt.Errorf("invalid credentials")
	}

	cook, valid := r.Cookie("session_id")

	fmt.Println("hvnhhnfnf" + cook.Value)
	fmt.Println("cookie is ", cook.Value)
	fmt.Println("sessihtdhtdhon", sessionID.String)

	fmt.Println(sessionID.Valid)
	fmt.Println(sessionID.String)

	// Check if the user already has an active session
	if valid != nil || cook.Value != "" {
		return 0, uuid.Nil, fmt.Errorf("user already has an active session")
	}

	// Generate a new session ID (you can use a library like UUID for this)
	newSessionID := generateSessionID()

	// Update the database with the new session ID
	_, err = DataBase.DB.Exec("UPDATE users SET session_id = ? WHERE id = ?", newSessionID, id)
	if err != nil {
		return 0, uuid.Nil, err
	}

	return id, newSessionID, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}

func generateSessionID() uuid.UUID {
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u2)
	return u2
}

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

func (DataBase *DB) InsertPost(u_ID int, title, post string, cats []int) error {
	statement, err := DataBase.DB.Prepare("INSERT INTO posts (u_ID, title, post, cat1, cat2, cat3) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(u_ID, title, post, cats[0], cats[1], cats[2])
	if err != nil {
		return err
	}
	return nil
}

func (DataBase *DB) InsertComment(u_ID, p_ID int, comment string) error {
	statement, err := DataBase.DB.Prepare("INSERT INTO comments (u_ID, p_ID, comment) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(u_ID, p_ID, comment)
	if err != nil {
		return err
	}
	return nil
}

func (DataBase *DB) InsertLike(u_ID, p_ID int, like int) error {

	lid, err := DataBase.LikeExists(p_ID, u_ID)
	if err != nil {
		return err
	}

	if lid != 0 {
		statement, err := DataBase.DB.Prepare("update author_liked_post set liked = ? where id = ?")
		if err != nil {
			return err
		}

		_, err = statement.Exec(like, lid)
		if err != nil {
			return err
		}
	} else {
		statement, err := DataBase.DB.Prepare("insert into author_liked_post (u_id, p_id, liked) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}

		_, err = statement.Exec(u_ID, p_ID, like)
		if err != nil {
			return err
		}
	}

	return nil
}

func (DataBase *DB) InsertLikeComment(u_ID, c_ID int, like int) error {

	lid, err := DataBase.LikeCommentExists(c_ID, u_ID)
	if err != nil {
		return err
	}

	if lid != 0 {
		statement, err := DataBase.DB.Prepare("update author_liked_comment set liked = ? where id = ?")
		if err != nil {
			return err
		}

		_, err = statement.Exec(like, lid)
		if err != nil {
			return err
		}
	} else {
		statement, err := DataBase.DB.Prepare("insert into author_liked_comment (u_id, c_id, liked) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}

		_, err = statement.Exec(u_ID, c_ID, like)
		if err != nil {
			return err
		}
	}

	return nil
}

func (DataBase *DB) LikeCommentExists(cid, uid int) (int, error) {
	statement, err := DataBase.DB.Prepare("SELECT id FROM author_liked_comment WHERE u_id = ? and c_id = ?")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(uid, cid).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (DataBase *DB) LikeExists(pid, uid int) (int, error) {
	statement, err := DataBase.DB.Prepare("SELECT id FROM author_liked_post WHERE u_id = ? and p_id = ?")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(uid, pid).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (DataBase *DB) LikesDislikesCommentTotal(cid string) (int, int, error) {
	likedStatement, err := DataBase.DB.Prepare("SELECT COUNT(id) AS total_likes_comment FROM author_liked_comment WHERE c_id = ? AND liked = 1;")
	if err != nil {
		return 0, 0, err
	}
	defer likedStatement.Close()

	var totalLikes int
	err = likedStatement.QueryRow(cid).Scan(&totalLikes)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}

	DislikedStatement, err := DataBase.DB.Prepare("SELECT COUNT(id) AS total_likes_comment FROM author_liked_comment WHERE c_id = ? AND liked = 0;")
	if err != nil {
		return 0, 0, err
	}
	defer DislikedStatement.Close()

	var totalDisLikes int
	err = DislikedStatement.QueryRow(cid).Scan(&totalDisLikes)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}

	return totalLikes, totalDisLikes, nil
}

func (DataBase *DB) LikesDislikesTotal(pid string) (int, int, error) {
	likedStatement, err := DataBase.DB.Prepare("SELECT COUNT(id) AS total_likes FROM author_liked_post WHERE p_id = ? AND liked = 1;")
	if err != nil {
		return 0, 0, err
	}
	defer likedStatement.Close()

	var totalLikes int
	err = likedStatement.QueryRow(pid).Scan(&totalLikes)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}

	DislikedStatement, err := DataBase.DB.Prepare("SELECT COUNT(id) AS total_likes FROM author_liked_post WHERE p_id = ? AND liked = 0;")
	if err != nil {
		return 0, 0, err
	}
	defer DislikedStatement.Close()

	var totalDisLikes int
	err = DislikedStatement.QueryRow(pid).Scan(&totalDisLikes)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}

	return totalLikes, totalDisLikes, nil
}

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
