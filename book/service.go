package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, BookRequest BookRequest) (Book, error)
	DeleteBook(ID int) (Book, error)
}

type service struct {
	repository Repository
}

// // DeleteBook implements Service.
// func (s *service) DeleteBook(ID int) (Book, error) {
// 	panic("unimplemented")
// }

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err

}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: *bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)
	if err != nil {
		return Book{}, err
	}
	return newBook, err

}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Book{}, err
	}

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = *bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	if err != nil {
		return Book{}, err
	}
	return newBook, err
}
func (s *service) DeleteBook(ID int) (Book, error) {
	// Cek apakah buku dengan ID tersebut ada
	_, err := s.repository.FindByID(ID)
	if err != nil {
		return Book{}, err
	}

	// Hapus buku dari repository
	deletedBook, err := s.repository.DeleteBook(ID)
	if err != nil {
		return Book{}, err
	}

	return deletedBook, nil
}
