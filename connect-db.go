package main

import (
	"database/sql"
	"time"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:root@tcp(database_my:3306)/crm_mediatama")
	
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Mysql Connnection Ready")

	defer db.Close()
}