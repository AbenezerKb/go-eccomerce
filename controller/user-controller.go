package controller

import (
	"context"	
	"gin-exercise/Errors"
	"gin-exercise/entity"
	"gin-exercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Search(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Payment(ctx *gin.Context)
}

type usercontroller struct {
	service service.UserService
}


func NewUserController(service service.UserService) UserController {
	return usercontroller{service: service}
}

func (c usercontroller) FindAll(ctx *gin.Context) {

	page := service.Pagination{}

	err := ctx.BindJSON(&page)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	list, err := c.service.FindAll(contx, page)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}
	total,err := c.service.UserCount(contx)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meta_data":total,"data":list})
	return

}

func (c usercontroller) Save(ctx *gin.Context) {
	var user entity.User

	err := ctx.ShouldBindJSON(&user)

	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	saveduser, err := c.service.Save(contx, user)

	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)

	}
	
	ctx.JSON(http.StatusOK, saveduser)
	return

}

func (c usercontroller) Delete(ctx *gin.Context) {

	id := ctx.Param("userid")
	contx := ctx.Request.Context()
	if id == "" {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return

	}

	res, err := c.service.Delete(contx, id)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}

func (c usercontroller) Update(ctx *gin.Context) {

	var user entity.UpdateUser
	var id string
	contx := ctx.Request.Context()

	err := ctx.ShouldBindJSON(&user)
	id = ctx.Param("id")

	user.Id = id

	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}

	if id == "" {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}

		updateduser, err := c.service.Update(contx, id, user)
		if err != nil {
			contx = context.WithValue(contx, Errors.UNABLE_TO_SAVE, err)
			customerror := Errors.Error(Errors.UNAUTHORIZED, contx, http.StatusInternalServerError)
			ctx.JSON(http.StatusInternalServerError, *customerror)

			return
		}

		ctx.JSON(http.StatusOK, updateduser)
		return

}

//Search
func (c usercontroller) Search(ctx *gin.Context) {

	var email service.Email

	err := ctx.ShouldBindJSON(&email)
	contx := ctx.Request.Context()


	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}


	result, err := c.service.Search(contx,email.Email)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusInternalServerError, *customerror)

		return
	}	

	

	ctx.JSON(http.StatusOK, result)
		
	return
	
}
