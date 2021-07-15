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
	//defer Db.Close()
}

func main()  {

	r, err := Db.Exec("insert into user_account(id,name, money)values(?,?,?)",3,"mumu",200)

	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println("insert id =", id)
}
