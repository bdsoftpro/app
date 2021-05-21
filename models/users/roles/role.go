package rolemodels

import (
	"database/sql"
)

// Role Format
type Role struct {
	Id		int64	`json:"id"`
	Name	string	`json:"name"`
	uId		int64	`json:"uid"`
}
// DB Format
type DB struct {
	Db *sql.DB
}
// Create Method Defination
func (rdb DB) Create(role *Role) error {
	stmt, err := rdb.Db.Prepare("insert into roles(name, uid) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(role.Name, role.uId)
	if err2 != nil {
		return err2
	}
	role.Id, _ = result.LastInsertId()
	return nil
}

// Update Method Defination
func (rdb DB) Update(role Role) (int64, error) {
	stmt, err := rdb.Db.Prepare("update roles set name = ? where uid = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(role.Name, role.uId)
	if err2 != nil {
		return 0, err2
	}
	return result.RowsAffected()
}

// Find Method Defination
func (rdb DB) Find(uid int64) (Role, error) {
	stmt, err := rdb.Db.Prepare("select id, name, uid from roles where uid = ?")
	if err != nil {
		return Role{}, err
	}
	defer stmt.Close()
	var role Role
	err = stmt.QueryRow(uid).Scan(&role.Id, &role.Name, &role.uId)
	if err != nil {
		return Role{}, err
	}
	return role, nil
}

// FindAll Method Defination
func (rdb DB) FindAll() ([]Role, error) {
	rows, err := rdb.Db.Query("select id, name, uid from roles")
	if err != nil {
		return nil, err
	}
	roles := []Role{}
	for rows.Next() {
		var role Role
		err2 := rows.Scan(&role.Id, &role.Name, &role.uId)
		if err2 != nil {
			return nil, err2
		}
		roles = append(roles, role)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
