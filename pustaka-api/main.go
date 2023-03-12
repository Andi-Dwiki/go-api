package main

import (
	"log"
	"pustaka-api/book"
	h "pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	bookService := book.NewService(bookRepository)

	bookHandler := h.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBooks)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8080")
}

// main
// handler
// service
// repository
// db
// mysql

// CRUD

// CREATE (C)
// book := book.Book{}
// book.Title = "stom"
// book.Price = 12000
// book.Discount = 15
// book.Rating = 4
// book.Description = "suiiiiiiiiiii"

// err = db.Create(&book).Error
// if err != nil {
// 	fmt.Println("=========================")
// 	fmt.Println("error creating book record")
// 	fmt.Println("============================")
// }

// gorm.io/docs/query.html<--dokumentasi read
// var books []book.Book

// err = db.Debug().Find(&books).Error
// if err != nil {
// 	fmt.Println("=========================")
// 	fmt.Println("error finding book record")
// 	fmt.Println("============================")
// }
// for _, b := range books {
// 	fmt.Println("Title: ", b.Title)
// 	fmt.Println("book object ", b)
// }

// Update

//Update
// var book book.Book

// err = db.Debug().Where("id = ?", 5).First(&book).Error
// if err != nil {
// 	fmt.Println("=========================")
// 	fmt.Println("error finding book record")
// 	fmt.Println("==========================")
// }

// book.Title = "manusia setengah macan (revised edition)"
// err = db.Save(&book).Error
// if err != nil {
// 	fmt.Println("=========================")
// 	fmt.Println("error updating book record")
// 	fmt.Println("==========================")
// }

// delete
// var book book.Book

// err = db.Debug().Where("id = ?", 5).First(&book).Error
// if err != nil {
// 	fmt.Println("=========================")
// 	fmt.Println("error finding book record")
// 	fmt.Println("==========================")
// }

// err = db.Delete(&book).Error
// if err != nil {
// 	fmt.Println("=========================")
// 	fmt.Println("error deleting book record")
// 	fmt.Println("==========================")
// }
