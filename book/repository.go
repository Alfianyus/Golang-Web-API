package book

import (
	"log"

	"gorm.io/gorm"
)

// type Repository interface {
// 	FindAll() ([]Book, error)
// 	FindByID(ID int) (Book, error)
// 	Create(book Book) (Book, error)
// 	Update(book Book) (Book, error)
// 	DeleteBook(ID int) (Book, error)
// }

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	DeleteBook(ID int) (Book, error) // Menggunakan ID sebagai parameter
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *fileRepository) DeleteBook(ID int) (Book, error) {
	var book Book
	err := r.db.Delete(&book, ID).Error
	if err != nil {
		log.Println("Error saat query Delete di repository:", err)
		return Book{}, err
	}
	return book, nil
}
