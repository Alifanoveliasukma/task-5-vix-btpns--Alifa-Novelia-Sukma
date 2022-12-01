package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Alifanoveliasukma/golang_apl/dto"
	"github.com/Alifanoveliasukma/golang_apl/entity"
	"github.com/Alifanoveliasukma/golang_apl/helper"
	"github.com/Alifanoveliasukma/golang_apl/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Controller is a ...
type ImageController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type imageController struct {
	imageService service.ImageService
	jwtService   service.JWTService
}

// NewController create a new instances of BoookController
func NewImageController(imageServ service.ImageService, jwtServ service.JWTService) ImageController {
	return &imageController{
		imageService: imageServ,
		jwtService:   jwtServ,
	}
}

func (c *imageController) All(context *gin.Context) {
	var images []entity.Image = c.imageService.All()
	res := helper.BuildResponse(true, "OK", images)
	context.JSON(http.StatusOK, res)
}

func (c *imageController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var image entity.Image = c.imageService.FindByID(id)
	if (image == entity.Image{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", image)
		context.JSON(http.StatusOK, res)
	}
}

func (c *imageController) Insert(context *gin.Context) {
	var imageCreateDTO dto.ImageCreateDTO
	errDTO := context.ShouldBind(&imageCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			imageCreateDTO.UserID = convertedUserID
		}
		result := c.imageService.Insert(imageCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *imageController) Update(context *gin.Context) {
	var imageUpdateDTO dto.ImageUpdateDTO
	errDTO := context.ShouldBind(&imageUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.imageService.IsAllowedToEdit(userID, imageUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			imageUpdateDTO.UserID = id
		}
		result := c.imageService.Update(imageUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *imageController) Delete(context *gin.Context) {
	var image entity.Image
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	image.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.imageService.IsAllowedToEdit(userID, image.ID) {
		c.imageService.Delete(image)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
func (c *imageController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
