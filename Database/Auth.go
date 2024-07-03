package database

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
)

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

func (DataBase *DB) UsernameExists(InputEmail string) (bool, error) {
	statement, err := DataBase.DB.Prepare("SELECT id FROM users WHERE username = ?")
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

	if !CheckPasswordHash(inputPassword, password) {
		return 0, uuid.Nil, fmt.Errorf("invalid credentials")
	}

	cook, valid := r.Cookie("session_id")


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

func (DataBase *DB) LikeCommentExists(pid, cid, uid int) (int, error) {
	statement, err := DataBase.DB.Prepare("SELECT id FROM author_liked_comment WHERE u_id = ? and p_id = ? and c_id = ?")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(uid, pid, cid).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return id, nil
}
