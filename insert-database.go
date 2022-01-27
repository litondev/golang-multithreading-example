package main

import (
	"database/sql"
	"fmt"
	"time"
	"sync"
	"os"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

var wg sync.WaitGroup

func insertData(thread int){
	fmt.Println(thread);

	database_my := os.Getenv("database_my")	

	fmt.Println(database_my)

	db, err := sql.Open("mysql", "root:root@tcp(" + database_my + ":3306)/crm_mediatama")
	
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 30)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Mysql Connnection Ready")

	sqlStr := "INSERT INTO visits(date, school_id, user_id, discription) VALUES "
	vals := []interface{}{}

	for i := 0;i < 10;i++{
	    sqlStr += "(?, ?, ?),"
	    vals = append(vals, "2020-01-01", 1, 1,"des-" + strconv.Itoa(i))
	}

	// trim the last ,
	sqlStr = sqlStr[0:len(sqlStr)-1]

	// prepare the statement
	stmt, errStmt := db.Prepare(sqlStr)

	if(errStmt != nil){
		panic(errStmt)
	}

	// format all vals at once
	res,errExc := stmt.Exec(vals...)

	if(errExc != nil){
		panic(errExc)
	}

	fmt.Println(res.RowsAffected());

	db.Close()

	wg.Done();
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go insertData(i)
	}

	wg.Wait()
}
