package repository

import (
	"context"
	"fmt"
	"library-project/config"
	"library-project/dto"
	"library-project/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	// InsertBook(book entity.Book) entity.Book
	Fetch(ctx context.Context) ([]entity.Book, error)
	FetchByID(ctx context.Context, id string) (entity.Book, error)
	FetchByTitle(ctx context.Context, title string) ([]entity.Book, error)
	Create(ctx context.Context, book dto.BookCreateDTO) (entity.Book, error)
}

type bookRepositoryImpl struct {
	coll *mongo.Collection
}

func NewBookRepository(client *mongo.Client) *bookRepositoryImpl {

	fmt.Printf("client: %v", client)
	db := client.Database(config.Configuration.DB.LibraryDatabase)
	coll := db.Collection(config.Configuration.DB.BookCollection)
	return &bookRepositoryImpl{
		coll: coll,
	}
}

func (b *bookRepositoryImpl) Create(ctx context.Context, bookDTO dto.BookCreateDTO) (book entity.Book, err error) {
	_, err = b.coll.InsertOne(
		context.TODO(), bookDTO,
	)
	if err != nil {
		return entity.Book{}, err // Mengembalikan error jika terjadi kesalahan saat menyimpan
	}

	err = b.coll.FindOne(ctx, bookDTO).Decode(&book)
	if err != nil {
		log.Printf("%s", err)
		return entity.Book{}, err
	}
	if err != nil {
		log.Printf("%s", err)
		return entity.Book{}, err // Mengembalikan error jika terjadi kesalahan saat menyimpan
	}
	return
}

func (b *bookRepositoryImpl) Fetch(ctx context.Context) (books []entity.Book, err error) {

	cursor, err := b.coll.Find(ctx, bson.D{})

	if err != nil {
		log.Printf("Failed to fetch books: %v\n", err)
		return nil, err
	}

	err = cursor.All(ctx, &books)

	if err != nil {
		log.Printf("Failed to decode books: %v\n", err)
		return nil, err
	}

	log.Println("Success get books from DB")
	return books, nil
}

func (b *bookRepositoryImpl) FetchByID(ctx context.Context, id string) (book entity.Book, err error) {

	idHex, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return book, err
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: idHex,
		},
	}
	log.Println(filter)
	mongoErr := b.coll.FindOne(ctx, filter).Decode(&book)
	if mongoErr != nil {
		if mongoErr == mongo.ErrNoDocuments {
			log.Printf("data not found %v", mongoErr)
			return entity.Book{}, mongoErr
		}
		log.Printf("%s", mongoErr)
		return entity.Book{}, mongoErr
	}
	log.Printf("Success get book %s from DB\n", id)
	return book, mongoErr
}

func (b *bookRepositoryImpl) FetchByTitle(ctx context.Context, title string) (books []entity.Book, err error) {

	filter := bson.D{
		{
			Key:   "title",
			Value: title,
		},
	}

	cursor, err := b.coll.Find(ctx, filter)

	if err != nil {
		log.Printf("Failed to fetch books: %v\n", err)
		return nil, err
	}

	err = cursor.All(ctx, &books)

	if err != nil {
		log.Printf("Failed to decode books: %v\n", err)
		return nil, err
	}

	log.Printf("Success get book %s from DB\n", title)
	return books, err
}
