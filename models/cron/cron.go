package crmodels

import (
	"database/sql"
)

// CronData Format
type CronData struct {
	Id		int64	`json:"id"`
	Uid		int64	`json:"uid"`
	Link	string	`json:"link"`
	Name	string	`json:"name"`
	From	string	`json:"from_at"`
	To		string	`json:"to_at"`
	Time	string	`json:"time"`
	Repeat	int64	`json:"repeat"`
	Status	bool	`json:"status"`
}
// DB Format
type DB struct {
	Db *sql.DB
}
// Create Method Defination
func (crb DB) Create(crondata *CronData) error {
	stmt, err := crb.Db.Prepare("insert into crons(uid, link, name, from_at, to_at, time, repeat, status) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(crondata.Uid, crondata.Link, crondata.Name, crondata.From, crondata.To, crondata.Time, crondata.Repeat, crondata.Status)
	if err2 != nil {
		return err2
	}
	crondata.Id, _ = result.LastInsertId()
	return nil
}

// Update Method Defination
func (crb DB) Update(crondata CronData) (int64, error) {	
	stmt, err := crb.Db.Prepare("update crons set link = ?, name = ?, from_at = ?, to_at = ?, time = ?, status = ? where id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(crondata.Link, crondata.Name, crondata.From, crondata.To, crondata.Time, crondata.Status, crondata.Id)
	if err2 != nil {
		return 0, err2
	}
	return result.RowsAffected()
}

// Find Method Defination
func (crb DB) Find(id int64) (CronData, error) {
	stmt, err := crb.Db.Prepare("select id, uid, link, name, from_at, to_at, time, repeat, status from crons where id = ?")
	if err != nil {
		return CronData{}, err
	}
	defer stmt.Close()
	var crondata CronData
	err = stmt.QueryRow(id).Scan(&crondata.Id, &crondata.Uid, &crondata.Link, &crondata.Name, &crondata.From, &crondata.To, &crondata.Time, &crondata.Repeat, &crondata.Status)
	if err != nil {
		return CronData{}, err
	}
	return crondata, nil
}

// FindAll Method Defination
func (crb DB) FindAll() ([]CronData, error) {
	rows, err := crb.Db.Query("select id, uid, link, name, from_at, to_at, time, repeat, status from crons")
	if err != nil {
		return nil, err
	}
	crondatas := []CronData{}
	for rows.Next() {
		var crondata CronData
		err2 := rows.Scan(&crondata.Id, &crondata.Uid, &crondata.Link, &crondata.Name, &crondata.From, &crondata.To, &crondata.Time, &crondata.Repeat, &crondata.Status)
		if err2 != nil {
		panic(err2)
			return nil, err2
		}
		crondatas = append(crondatas, crondata)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return crondatas, nil
}
