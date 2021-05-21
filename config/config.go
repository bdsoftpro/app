package config

import (
	"database/sql"
	_ "github.com/bdsoftpro/sqlite3"
	"net/url"
	"fmt"
)

func GetSqlDB() (db *sql.DB, err error) {
	key := url.QueryEscape("del!@12sha")
	dbname := fmt.Sprintf("file:shop.dll?_pragma_key=%s&_pragma_cipher_page_size=4096", key)
	db, err = sql.Open("sqlite3", dbname)
	return 
}