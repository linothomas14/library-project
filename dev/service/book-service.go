package service

import (
	"context"
	"library-project/dto"
	"library-project/entity"
	"library-project/repository"
)

type BookService interface {
	FetchByID(ctx context.Context, id string) (entity.Book, error)
	FetchByTitle(ctx context.Context, title string) ([]entity.Book, error)
	Fetch(ctx context.Context) ([]entity.Book, error)
	Create(ctx context.Context, book dto.BookCreateDTO) (entity.Book, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRep repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRep,
	}
}

func (service *bookService) Fetch(ctx context.Context) ([]entity.Book, error) {

	return service.bookRepository.Fetch(ctx)
}

func (service *bookService) FetchByID(ctx context.Context, id string) (entity.Book, error) {
	return service.bookRepository.FetchByID(ctx, id)
}

func (service *bookService) FetchByTitle(ctx context.Context, title string) ([]entity.Book, error) {
	return service.bookRepository.FetchByTitle(ctx, title)
}

func (service *bookService) Create(ctx context.Context, book dto.BookCreateDTO) (entity.Book, error) {
	return service.bookRepository.Create(ctx, book)
}

// func (service *bookService) Insert(m dto.BookCreateDTO) entity.Book {
// 	mhs := entity.Book{}
// 	err := smapping.FillStruct(&mhs, smapping.MapFields(&m))
// 	if err != nil {
// 		log.Fatalf("Failed map %v: ", err)
// 	}
// 	res := service.bookRepository.InsertBook(mhs)
// 	return res
// }
