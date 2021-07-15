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

type Account struct {
	Id int 		`db:"id"`
	Name string	`db:"name"`
	Money int	`db:"money"`
}

func main()  {
	var account []Account
	err := Db.Select(&account, "select id,name,money from user_account where id=?", 2)
	if err != nil {
		fmt.Println("err=", err)
	}

	fmt.Println(account)
}
