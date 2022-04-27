package controller

import (
	"gin-exercise/service"	
	"gin-exercise/Errors"	
	"net/http"
	"context"
	"github.com/gin-gonic/gin"
	//    "github.com/google/uuid" // To generate random file names
)

type ImageController interface {
	ImageSave(ctx *gin.Context)
	Display_image(ctx *gin.Context)
}

type imgcontroller struct {
	service service.ImageService
}

func NewImage(service service.ImageService) ImageController {
	return imgcontroller{service: service}
}

func (c imgcontroller) ImageSave(ctx *gin.Context) {    
	
	contx := ctx.Request.Context()	
	file,err :=ctx.FormFile("file")
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}
	// ctx.SaveUploadedFile()
	//ctx.JSON(200, c.service.Save(ctx))
	id := ctx.Param("id")	
	err =c.service.Save(file, id)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_SAVE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}
	ctx.JSON(http.StatusOK,nil)
}


func (c imgcontroller) Display_image(ctx *gin.Context) {

	c.service.Display(ctx)

	// var user entity.User

	// err := ctx.ShouldBindJSON(&user)
	// if err != nil {
	// 	rest_error.NewBadRequestError(("error, registration failed"))
	// }
	// c.service.Save(ctx)
	// ctx.JSON(http.StatusOK, gin.H{"Message": "user input is valid!"})

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Your file has been successfully uploaded.",
	// })

}
