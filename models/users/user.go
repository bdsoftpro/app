package usrmodels

import (
	"errors"
	"database/sql"
	"github.com/bdsoftpro/app/system"
)

// User Format
type User struct {
	Id		int64	`json:"id"`
	Fname	string	`json:"fname"`
	Lname	string	`json:"lname"`
	Uname	string	`json:"uname"`
	Passwd	string	`json:"passwd"`
	Email	string	`json:"email"`
	Gender	string	`json:"gender"`
	RememberToken	string	`json:"remember_token"`
	Status	int8	`json:"status"`
	Created_at	string	`json:"created_at"`
	Updated_at	string	`json:"updated_at"`
}
// DB Format
type DB struct {
	Db *sql.DB
}
// Create Method Defination
func (udb DB) Create(user *User) error {
	stmt, err := udb.Db.Prepare("insert into users(fname, lname, uname, passwd, email, gender, status, created_at) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(user.Fname, user.Lname, user.Uname, user.Passwd, user.Email, user.Gender, user.Status, user.Created_at)
	if err2 != nil {
		return err2
	}
	user.Id, _ = result.LastInsertId()
	return nil
}

// Login Method Defination
func (udb DB) Login(usr User) (User, error) {
	var user User
	stmt, err := udb.Db.Prepare("select id, fname, lname, uname, email, passwd, gender, status from users where uname = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(usr.Uname).Scan(&user.Id, &user.Fname, &user.Lname, &user.Uname, &user.Email, &user.Passwd, &user.Gender, &user.Status)
	if err != nil {
		return user, err
	}
	if System.Decrypt(user.Passwd) != usr.Passwd {
		return user, errors.New("the password not match")
	}
	return user, nil
	
}

// Update Method Defination
func (udb DB) Update(user User) (int64, error) {
	stmt, err := udb.Db.Prepare("update users set fname = ?, lname = ?, email = ?, status = ?, updated_at = ? where id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(user.Fname, user.Lname, user.Email, user.Status, user.Updated_at, user.Id)
	if err2 != nil {
		return 0, err2
	}
	return result.RowsAffected()
}

// Find Method Defination
func (udb DB) Find(id int64) (User, error) {
	stmt, err := udb.Db.Prepare("select id, fname, lname, uname, email, gender, status from users where id = ?")
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()
	var user User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Fname, &user.Lname, &user.Uname, &user.Email, &user.Gender, &user.Status)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// FindAll Method Defination
func (udb DB) FindAll() ([]User, error) {
	rows, err := udb.Db.Query("select id, fname, lname, uname, email, gender, status from users")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var user User
		err2 := rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Uname, &user.Email, &user.Gender, &user.Status)
		if err2 != nil {
			return nil, err2
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}
