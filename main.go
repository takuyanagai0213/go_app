package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	dbInit()

	todos := dbGetAll()

	fmt.Println(todos)
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.Run("localhost:8080")
}

func dbConnect() *gorm.DB {
	dsn := "root@tcp(localhost:3306)/go_app?parseTime=true"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func dbInit() {
	db := dbConnect()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

func dbInsert(text string, status string) {
	db := dbConnect()
	todo := Todo{Text: text, Status: status}
	db.Create(&todo)
	defer db.Close()
}

func dbGetAll() []Todo {
	db := dbConnect()
	var todos []Todo
	db.Find(&todos)
	db.Close()
	return todos
}
func dbGetOne(id int) Todo {
	db := dbConnect()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
