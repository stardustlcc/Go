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
	res , err := Db.Exec("update user_account set name=? where id=?", "翠翠",1)
	if err != nil {
		fmt.Println("err=", err)
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("res.RowsAffected", err)
	}

	fmt.Println("success row=", row)

}
