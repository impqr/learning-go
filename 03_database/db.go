package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	host    = "localhost"
	port    = "3306"
	user    = "root"
	pass    = "pang1993"
	scheme  = "test"
	charset = "utf8"
)

// Account sqlx中对于允许为Null的字段，需要用sql.NullXXX填充
type Account struct {
	ID         sql.NullInt64  `db:"id"`
	Username   sql.NullString `db:"username"`
	Password   sql.NullString `db:"password"`
	Name       sql.NullString `db:"name"`
	Remark     sql.NullString `db:"remark"`
	Status     sql.NullInt64  `db:"status"`
	CreateTime sql.NullInt64  `db:"create_time"`
	UpdateTime sql.NullInt64  `db:"update_time"`
}

func main() {
	db := connect()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(db)

	_ = db.Ping()
	log.Printf("database connect successfully")

	var accounts []Account
	err := db.Select(&accounts, "select * from account")
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(len(accounts))

	for _, account := range accounts {
		fmt.Println(account.ID.Int64)
		fmt.Println(account)
	}
}

func connect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		user,
		pass,
		host,
		port,
		scheme,
		charset,
	))
	if err != nil {
		log.Panicln(err)
	}

	return db
}
