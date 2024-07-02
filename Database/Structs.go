package database

import "database/sql"

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
	ID       int
	U_ID     int
	P_ID     int
	Username string
	Comment  string
	Likes    int
	Dislikes int
}

var DataBase DB
