package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello gin!")
	})

	router.Run("localhost:8080")
}
func dbConnect() {
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

}
