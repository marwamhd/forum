package database

import (
	"database/sql"
	"log"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

//CommentLikesDislikesTotal

func (DataBase *DB) CommentLikesDislikesTotal(pid, cid string) (int, int, error) {
	likedStatement, err := DataBase.DB.Prepare("SELECT COUNT(id) AS total_likes FROM author_liked_comment WHERE p_id = ? AND c_id = ? AND liked = 1;")
	if err != nil {
		return 0, 0, err
	}
	defer likedStatement.Close()

	var totalLikes int
	err = likedStatement.QueryRow(pid, cid).Scan(&totalLikes)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}

	DislikedStatement, err := DataBase.DB.Prepare("SELECT COUNT(id) AS total_likes FROM author_liked_comment WHERE p_id = ? and c_id = ? AND liked = 0;")
	if err != nil {
		return 0, 0, err
	}
	defer DislikedStatement.Close()

	var totalDisLikes int
	err = DislikedStatement.QueryRow(pid, cid).Scan(&totalDisLikes)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}

	return totalLikes, totalDisLikes, nil
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
