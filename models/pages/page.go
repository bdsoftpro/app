package pgmodels

import (
	"database/sql"
)

// Page Format
type Page struct {
	Id		int64	`json:"id"`
	Name	string	`json:"name"`
	Price	float32	`json:"price"`
	Quantity int	`json:"quantity"`
	Status	bool	`json:"status"`
}
// DB Format
type DB struct {
	Db *sql.DB
}
// Create Method Defination
func (pgdb DB) Create(page *Page) error {
	stmt, err := pgdb.Db.Prepare("insert into pages(name, price, quantity, status) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(page.Name, page.Price, page.Quantity, page.Status)
	if err2 != nil {
		return err2
	}
	page.Id, _ = result.LastInsertId()
	return nil
}

// Update Methos Defination
func (pgdb DB) Update(page Page) (int64, error) {
	stmt, err := pgdb.Db.Prepare("update pages set name = ?, price = ?, quantity = ?, status = ? where id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(page.Name, page.Price, page.Quantity, page.Status, page.Id)
	if err2 != nil {
		return 0, err2
	}
	return result.RowsAffected()
}

// Find Method Defination
func (pgdb DB) Find(id int64) (Page, error) {
	stmt, err := pgdb.Db.Prepare("select id, name, price, quantity, status from pages where id = ?")
	if err != nil {
		return Page{}, err
	}
	defer stmt.Close()
	var page Page
	err = stmt.QueryRow(id).Scan(&page.Id, &page.Name, &page.Price, &page.Quantity, &page.Status)
	if err != nil {
		return Page{}, err
	}
	return page, nil
}

// FindAll Methos Defination
func (pgdb DB) FindAll() ([]Page, error) {
	rows, err := pgdb.Db.Query("select id, name, price, quantity, status from pages")
	if err != nil {
		return nil, err
	}
	pages := []Page{}
	for rows.Next() {
		var page Page
		err2 := rows.Scan(&page.Id, &page.Name, &page.Price, &page.Quantity, &page.Status)
		if err2 != nil {
			return nil, err2
		}
		pages = append(pages, page)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return pages, nil
}
