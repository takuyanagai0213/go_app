package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// type Person struct {
// 	ID   int
// 	Name string
// }

func main() {
	db, err := sql.Open("mysql", "root@/go_app")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 1
	var name string
	err = db.QueryRow("SELECT name FROM sample_table WHERE id=?", id).Scan(&name)
	// defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(name)

	// for rows.Next() {
	// 	var person Person
	// 	fmt.Println(person.ID, person.Name)
	// }

}

