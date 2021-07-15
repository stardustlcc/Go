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
	res, err := Db.Exec("delete from user_account where id = ?", 3)
	if err != nil {
		fmt.Println("err=", err)
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("err=", err)
	}

	fmt.Println("delete id =", row)
}
