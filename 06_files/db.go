package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func Connect(data map[string]string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		data["user"],
		data["pass"],
		data["host"],
		data["port"],
		data["scheme"],
		data["charset"],
	))
	if err != nil {
		log.Panicln(err)
	}

	return db
}
