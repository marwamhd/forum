package database

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

//InsertCommentLike

func (DataBase *DB) InsertCommentLike(u_ID, p_ID, c_ID int, like int) error {

	lid, err := DataBase.LikeCommentExists(p_ID, c_ID, u_ID)
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
		statement, err := DataBase.DB.Prepare("insert into author_liked_comment (u_id, p_id,c_id, liked) VALUES (?, ?, ?, ?);")
		if err != nil {
			return err
		}

		_, err = statement.Exec(u_ID, p_ID, c_ID, like)
		if err != nil {
			return err
		}
	}

	return nil
}
