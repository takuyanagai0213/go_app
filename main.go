package main

import (
	"strconv"

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

	// fmt.Println(todos)

	// index
	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()

		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// new
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{
			"todo": todo,
		})
	})

	// Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})

	// 削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{
			"todo": todo,
		})
	})
	// Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		dbDelete(id)
		ctx.Redirect(302, "/")
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
func dbUpdate(id int, text string, status string) {
	db := dbConnect()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

func dbDelete(id int) {
	db := dbConnect()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}
