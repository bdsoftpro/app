package profmodels

import (
	"database/sql"
)

// User Format
type Profile struct {
	Id		int64	`json:"id"`
	uId		int64	`json:"uid"`
	Name	string	`json:"name"`
	Detail	string	`json:"detail"`
}
// DB Format
type DB struct {
	Db *sql.DB
}
// Create Method Defination
func (prdb DB) Create(profile *Profile) error {	
	stmt, err := prdb.Db.Prepare("insert into profiles(uid, name, detail) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(profile.uId, profile.Name, profile.Detail)
	if err2 != nil {
		return err2
	}
	profile.Id, _ = result.LastInsertId()
	return nil
}

// Update Method Defination
func (prdb DB) Update(profile Profile) (int64, error) {
	stmt, err := prdb.Db.Prepare("update profiles set name = ?, detail = ? where uid = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(profile.Name, profile.Detail, profile.uId)
	if err2 != nil {
		return 0, err2
	}
	return result.RowsAffected()
}

// Find Profile Value Method Defination
func (prdb DB) Value(uid int64, name string) (string, error) {
	stmt, err := prdb.Db.Prepare("select detail from profiles where uid = ? and name = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var profile Profile
	err = stmt.QueryRow(uid, name).Scan(&profile.Detail)
	if err != nil {
		return "", err
	}
	return profile.Detail, nil
}

// Find Profile All Method Defination
func (prdb DB) FindAll(uid int64) ([]Profile, error) {
	stmt, err := prdb.Db.Prepare("SELECT id, name, detail FROM profiles where uid = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(uid)
	if err2 != nil {
		return nil, err2
	}
	profiles := []Profile{}
	for rows.Next() {
		var profile Profile
		err2 = rows.Scan(&profile.Id, &profile.Name, &profile.Detail)
		if err2 != nil {
			return nil, err2
		}
		profiles = append(profiles, profile)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return profiles, nil
}
