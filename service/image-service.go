package service

import (
	"fmt"
	"log"

	"github.com/Alifanoveliasukma/golang_apl/dto"
	"github.com/Alifanoveliasukma/golang_apl/entity"
	"github.com/Alifanoveliasukma/golang_apl/repository"
	"github.com/mashingan/smapping"
)

// Service is a ....
type ImageService interface {
	Insert(b dto.ImageCreateDTO) entity.Image
	Update(b dto.ImageUpdateDTO) entity.Image
	Delete(b entity.Image)
	All() []entity.Image
	FindByID(imageID uint64) entity.Image
	IsAllowedToEdit(userID string, imageID uint64) bool
}

type imageService struct {
	imageRepository repository.ImageRepository
}

// NewService .....
func NewImageService(imageRepo repository.ImageRepository) ImageService {
	return &imageService{
		imageRepository: imageRepo,
	}
}

func (service *imageService) Insert(b dto.ImageCreateDTO) entity.Image {
	image := entity.Image{}
	err := smapping.FillStruct(&image, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.imageRepository.InsertImage(image)
	return res
}

func (service *imageService) Update(b dto.ImageUpdateDTO) entity.Image {
	image := entity.Image{}
	err := smapping.FillStruct(&image, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.imageRepository.UpdateImage(image)
	return res
}

func (service *imageService) Delete(b entity.Image) {
	service.imageRepository.DeleteImage(b)
}

func (service *imageService) All() []entity.Image {
	return service.imageRepository.AllImage()
}

func (service *imageService) FindByID(imageID uint64) entity.Image {
	return service.imageRepository.FindImageByID(imageID)
}

func (service *imageService) IsAllowedToEdit(userID string, imageID uint64) bool {
	b := service.imageRepository.FindImageByID(imageID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
