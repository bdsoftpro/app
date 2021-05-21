package pdmodels

import (
	"database/sql"
)

// Product Format
type Product struct {
	Id		int64
	Name	string
	Price	float32
	Quantity int
	Status	bool
}
// DB Format
type DB struct {
	Db *sql.DB
}
// Create Method Defination
func (pdb DB) Create(product *Product) error {
	stmt, err := pdb.Db.Prepare("insert into products(name, price, quantity, status) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(product.Name, product.Price, product.Quantity, product.Status)
	if err2 != nil {
		return err2
	}
	product.Id, _ = result.LastInsertId()
	return nil
}

// Update Method Defination
func (pdb DB) Update(product Product) (int64, error) {	
	stmt, err := pdb.Db.Prepare("update products set name = ?, price = ?, quantity = ?, status = ? where id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(product.Name, product.Price, product.Quantity, product.Status, product.Id)
	if err2 != nil {
		return 0, err2
	}
	return result.RowsAffected()
}

// Find Method Defination
func (pdb DB) Find(id int64) (Product, error) {
	stmt, err := pdb.Db.Prepare("select * from products where id = ?")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	var product Product
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Status)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// FindAll Method Defination
func (pdb DB) FindAll() ([]Product, error) {
	rows, err := pdb.Db.Query("select * from products")
	if err != nil {
		return nil, err
	}
	products := []Product{}
	for rows.Next() {
		var product Product
		err2 := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Status)
		if err2 != nil {
			return nil, err2
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}
