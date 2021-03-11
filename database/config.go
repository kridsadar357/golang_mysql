package database

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

func InitialDB() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "goAdmin"
	dbPass := "goTar35700"
	dbName := "db_erp"

	addr, err := net.LookupIP("mylab861.ddns.me")
	if err != nil {
		fmt.Println("Unknown host")
	}
	ipdatabase := addr[0].String()

	db, err2 := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+ipdatabase+":3306)/"+dbName)
	if err2 != nil {
		panic(err2.Error())
	}
	return db
}
