package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github/ayonsaha2011/golang-gin-poc/entity"
	"github/ayonsaha2011/golang-gin-poc/service"
	"github/ayonsaha2011/golang-gin-poc/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type videoController struct {
	service service.VideoService
}

var  validate *validator.Validate

func New(service service.VideoService) VideoController  {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *videoController)  Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}