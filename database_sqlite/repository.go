package database_sqlite

import "database/sql"

func CreateUser(user *User) error {
	result, err := DB.Exec("insert into user values (null, ?, ?)", user.Name, user.Password)
	if err != nil {
		return err
	}

	user.Id, _ = result.LastInsertId()
	return nil
}

func QueryUser(userId int64) (*User, error) {
	row := DB.QueryRow("select * from user where id = ?", userId)
	user := &User{}
	var err error
	if err = row.Scan(&user.Id, &user.Name, &user.Password); err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	return user, err
}

func DeleteUser(userId int64) error {
	_, err := DB.Exec("delete from user where id = ?", userId)
	return err
}
