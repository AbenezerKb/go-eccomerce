// categorycontroll

package controller

import (
	"context"
	"gin-exercise/Errors"
	"gin-exercise/entity"
	"gin-exercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)	
}

type categorycontroller struct {
	service service.CategoryService
}

// var (
// 	loginService         service.LoginService = &service.UserLogin{}
// 	jwtService           service.JWTService   = service.JWTAuthService()
// 	loginValidController LoginController      = LoginHandler(loginService, jwtService)
// )

func NewCategoryController(service service.CategoryService) CategoryController {
	return categorycontroller{service: service}
}

func (c categorycontroller) FindAll(ctx *gin.Context) {

	page := service.Pagination{}

	// err := ctx.BindJSON(&page)
	// contx := ctx.Request.Context()
	// if err != nil {
	// 	contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
	// 	customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
	// 	ctx.JSON(http.StatusBadRequest, *customerror)
	// 	return
	// }


	err := ctx.BindJSON(&page)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}


	list, err := c.service.FindAll(ctx, page)
	total,err := c.service.CategoryCount(contx)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}

	ctx.JSON(http.StatusOK,  gin.H{"meta_data":total,"data":list})
	return
	

}

func (c categorycontroller) Save(ctx *gin.Context) {
	var category entity.Category
	err := ctx.ShouldBindJSON(&category)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	savedcategory, err := c.service.Save(contx, category.Name)

	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)

	}
	
	ctx.JSON(http.StatusOK, savedcategory)
	return

}




