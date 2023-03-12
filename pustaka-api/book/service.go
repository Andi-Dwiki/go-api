package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(Id int, bookRequest BookRequest) (Book, error)
	Delete(Id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FIndALl
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
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	NewBook, err := s.repository.Create(book)
	return NewBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
