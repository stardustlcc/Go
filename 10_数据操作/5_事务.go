package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init()  {
	database, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/lcc")
	if err != nil {
		fmt.Println("err=", err)
	}
	Db = database
}

func main()  {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("err=", err)
	}

	r, err := conn.Exec("insert into user_account(id,name,money) values(?,?,?)", 3,"nita",300)
	if err != nil {
		conn.Rollback()
		fmt.Println("err=", err)
	}

	id, err := r.LastInsertId()
	if err != nil {
		conn.Rollback()
		fmt.Println("err=", err)
	}
	fmt.Println("id=", id)
	conn.Commit()




}
