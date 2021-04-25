package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

func queryRow(db *sql.DB, name *string) error {
	err := db.QueryRow("select name from users where id =?", 1).Scan(name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		} else {
			return errors.Wrap(err,"QueryRow error")
		}
	}
	return nil
}

func main() {
	//连接数据库
	db, err := sql.Open("mysql", "root:csw1235@/myApp")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var name string
	err = queryRow(db, &name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name:", name)
}
