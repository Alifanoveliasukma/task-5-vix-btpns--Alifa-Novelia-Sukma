package repository

import (
	"github.com/Alifanoveliasukma/golang_apl/entity"
	"gorm.io/gorm"
)

// BookRepository is a ....
type ImageRepository interface {
	InsertImage(b entity.Image) entity.Image
	UpdateImage(b entity.Image) entity.Image
	DeleteImage(b entity.Image)
	AllImage() []entity.Image
	FindImageByID(imageID uint64) entity.Image
}

type imageConnection struct {
	connection *gorm.DB
}

// NewBookRepository creates an instance BookRepository
func NewImageRepository(dbConn *gorm.DB) ImageRepository {
	return &imageConnection{
		connection: dbConn,
	}
}

func (db *imageConnection) InsertImage(b entity.Image) entity.Image {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *imageConnection) UpdateImage(b entity.Image) entity.Image {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *imageConnection) DeleteImage(b entity.Image) {
	db.connection.Delete(&b)
}

func (db *imageConnection) FindImageByID(imageID uint64) entity.Image {
	var image entity.Image
	db.connection.Preload("User").Find(&image, imageID)
	return image
}

func (db *imageConnection) AllImage() []entity.Image {
	var images []entity.Image
	db.connection.Preload("User").Find(&images)
	return images
}
