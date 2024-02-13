package controller

import (
	"library-project/dto"
	"library-project/entity"
	"library-project/helper"
	"library-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	Fetch(ctx *gin.Context)
	FetchByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context)
}

type bookController struct {
	//
	bookService service.BookService
}

func NewBookController(bookServ service.BookService) BookController {
	return &bookController{
		bookService: bookServ,
	}
}

func (c *bookController) Create(ctx *gin.Context) {
	var bookDTO dto.BookCreateDTO

	if err := ctx.BindJSON(&bookDTO); err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	book, err := c.bookService.Create(ctx, bookDTO)
	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	} else {
		res := helper.BuildResponse("OK", book)
		ctx.JSON(http.StatusOK, res)
		return
	}
}

func (c *bookController) Fetch(ctx *gin.Context) {

	var books []entity.Book
	var err error
	var res helper.Response

	title := ctx.Query("title")

	if title != "" {
		books, err = c.bookService.FetchByTitle(ctx, title)
	} else {
		books, err = c.bookService.Fetch(ctx)
	}

	if err != nil {
		res = helper.BuildResponse(err.Error(), books)
	}

	res = helper.BuildResponse("OK", books)
	ctx.JSON(http.StatusOK, res)
}

func (c *bookController) FetchByID(ctx *gin.Context) {
	id := ctx.Param("id")

	book, err := c.bookService.FetchByID(ctx, id)
	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse("OK", book)
		ctx.JSON(http.StatusOK, res)
	}
}
