package main

import (
	"fmt"
	"log"

	"tes-api-golang/book"
	"tes-api-golang/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:4306)/test_api_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	db.AutoMigrate(&book.Book{})

	repository := book.NewFileRepository(db)

	foundBook, err := repository.FindByID(1) // Gunakan foundBook
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Println("Book found:", foundBook)
	bookService := book.NewService(repository)

	bookHandler := handler.NewBookHandler(bookService)
	var books []book.Book
	err = db.Find(&books).Error
	if err != nil {
		fmt.Println("Error FindAll Book Record: ", err)
	} else {
		fmt.Println("Daftar buku dalam database:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Price: %d\n", book.ID, book.Title, book.Price)
		}
	}
	log.Println("Entering main function")

	if err != nil {
		fmt.Println("============================")
		fmt.Println("Error FindAll Book Record")
		fmt.Println("============================")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//main
	//service
	//repository
	//db
	//mysql

	v1 := router.Group("/v1")

	// Daftarkan handler

	v1.GET("/books/:id", bookHandler.GetBook)
	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	// v2 := router.Group("/v2")

	log.Println("Starting server on port 5007...")

	// Jalankan server pada port 5007
	if err := router.Run(":5007"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
