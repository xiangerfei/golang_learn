package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "dba_yix:xiangerfer@tcp(124.222.238.134:5730)/gin_test")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func main()  {
	type User struct {
		u string
		h string
	}
	//users := make([]User, 5, 100)

	db := GetDB()



	if rows, err := db.Query("select user, host from mysql.user"); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}else{
		if columns, err := rows.Columns(); err != nil{
			fmt.Printf("some things error %s", err.Error())
		}else {
			for _, c := range columns{
				fmt.Printf("column name: %s\n", c)
			}
		}

	}

	fmt.Println(strings.Repeat("*", 100))
	slice_str := []string{
		"hello",
		"world",
		"cool",
		"man",
	}

	for k, v := range slice_str{
		fmt.Printf("%d %s\n", k, v)
	}
	//fmt.Printf("user: %s \t host: %s\n", user.u, user.h)
}