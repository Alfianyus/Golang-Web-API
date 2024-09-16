package book

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error // Mengambil semua buku dari database
	if err != nil {
		log.Println("Error saat query FindAll di repository:", err) // Log error jika ada
		return nil, err
	}
	return books, nil

}

func (r *fileRepository) FindByID(ID int) (Book, error) {

	var book Book
	err := r.db.First(&book, ID).Error // Menggunakan First untuk mendapatkan satu entri berdasarkan ID
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Book{}, fmt.Errorf("book with ID %d not found", ID)
		}
		log.Println("Error saat query FindByID di repository:", err)
		return Book{}, err
	}
	return book, nil

}

func (r *fileRepository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		log.Println("Error saat query Create di repository:", err)
		return Book{}, err
	}
	return book, nil
}

func (r *fileRepository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		log.Println("Error saat query Update di repository:", err)
		return Book{}, err
	}
	return book, nil
}
func (r *fileRepository) Delete(ID int) error {
	// Hapus buku dari repository
	err := r.db.Delete(&Book{}, ID).Error
	return err
}
