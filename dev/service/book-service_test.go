package service

import (
	"context"
	"errors"
	"library-project/dto"
	"library-project/entity"
	mocks "library-project/mocks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBookService_Create(t *testing.T) {

	mockBookRepo := new(mocks.BookRepository)
	app := NewBookService(mockBookRepo)

	t.Run("Valid Request", func(t *testing.T) {

		expectedBookResp := entity.Book{
			ID:     primitive.NewObjectID(),
			Title:  "Title 1",
			Author: "Author 1",
			Genre:  "Genre 1",
		}

		ctx := context.Background()

		bookDTO := dto.BookCreateDTO{
			Title:  "Title 1",
			Author: "Author 1",
		}

		mockBookRepo.On("Create", ctx, bookDTO).Return(expectedBookResp, nil)

		actualBookResp, err := app.Create(ctx, bookDTO)

		assert.NoError(t, err)
		assert.NotEmpty(t, actualBookResp)
		assert.Equal(t, expectedBookResp, actualBookResp)
	})
	t.Run("Invalid Request", func(t *testing.T) {

		bookResp := entity.Book{}
		ctx := context.Background()
		book := dto.BookCreateDTO{}

		mockBookRepo.On("Create", ctx, book).Return(bookResp, errors.New("create error"))

		bookResp, err := app.Create(ctx, book)
		assert.Error(t, err)
	})
}

func TestBookService_Fetch(t *testing.T) {

	mockBookRepo := new(mocks.BookRepository)
	app := NewBookService(mockBookRepo)

	t.Run("Valid Request", func(t *testing.T) {

		ctx := context.Background()

		expectedBooks := []entity.Book{
			{
				Title:  "Book 1",
				Author: "Author 1",
			},
			{
				Title:  "Book 2",
				Author: "Author 2",
			},
		}

		mockBookRepo.On("Fetch", ctx).Return(expectedBooks, nil)

		books, err := app.Fetch(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, books)
		assert.Equal(t, expectedBooks, books)
	})

}

func TestBookService_FetchByID(t *testing.T) {

	mockBookRepo := new(mocks.BookRepository)
	app := NewBookService(mockBookRepo)

	t.Run("Valid Request", func(t *testing.T) {

		ctx := context.Background()

		expectedBooks := entity.Book{
			Title:  "Book 1",
			Author: "Author 1",
		}
		id := "65c9afbcaa41078886bf5bec" // <-- contoh id

		mockBookRepo.On("FetchByID", ctx, id).Return(expectedBooks, nil)

		book, err := app.FetchByID(ctx, id)
		assert.NoError(t, err)
		assert.NotEmpty(t, book)
		assert.Equal(t, expectedBooks, book)
	})

}

func TestBookService_FetchByTitle(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {

		mockBookRepo := new(mocks.BookRepository)

		app := NewBookService(mockBookRepo)

		ctx := context.Background()

		// Define the expected books and title
		expectedBooks := []entity.Book{
			{Title: "Book 1", Author: "Author 1"},
			{Title: "Book 2", Author: "Author 2"},
		}
		title := "Book 1"

		mockBookRepo.On("FetchByTitle", ctx, title).Return(expectedBooks, nil)

		books, err := app.FetchByTitle(ctx, title)

		assert.NoError(t, err)

		assert.NotEmpty(t, books)
		assert.Equal(t, expectedBooks, books)
	})

}
