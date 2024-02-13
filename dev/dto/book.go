package dto

type BookCreateDTO struct {
	Title  string `json:"title" bson:"title" form:"title" binding:"required"`
	Author string `json:"author" bson:"author" form:"author" binding:"required"`
	Genre  string `json:"genre" bsion:"genre" form:"genre" binding:"required"`
}
