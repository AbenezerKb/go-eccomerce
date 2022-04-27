package controller

import (	
	"gin-exercise/entity"
	"gin-exercise/Errors"
	"gin-exercise/service"
	"net/http"
	"context"
	"github.com/gin-gonic/gin"
)

type CartController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	
	
}

type cartcontroller struct {
	service service.CartService
}

// var (
// 	loginService         service.LoginService = &service.UserLogin{}
// 	jwtService           service.JWTService   = service.JWTAuthService()
// 	loginValidController LoginController      = LoginHandler(loginService, jwtService)
// )

func NewCartController(service service.CartService) CartController {
	return cartcontroller{service: service}
}

func (c cartcontroller) FindAll(ctx *gin.Context) {


	page := service.Pagination{}

	err := ctx.BindJSON(&page)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	user_id := ctx.GetString("userID")

	if user_id == "" {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	list, err := c.service.FindAll(contx, user_id, page)
	total, err := c.service.CartItemCount(contx, user_id)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meta_data": total, "data": list})
	return

}

func (c cartcontroller) Save(ctx *gin.Context) {
	
	var cart entity.Cart	

	err := ctx.ShouldBindJSON(&cart)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}


	savedcart, err := c.service.Save(contx,cart)


	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)

	}
	
	ctx.JSON(http.StatusOK, *savedcart)
	return

}



// func (c cartcontroller) Update(ctx *gin.Context) {
// 	var cart entity.UpdateOrder
// 	err := ctx.ShouldBindJSON(&cart)
// 	name := ctx.Param("name")
// 	if err != nil {
// 		rest_error.NewBadRequestError(("error, registration failed"))
// 	}
// 	_U, er := c.service.Update(name,cart)
// 	if er == nil {
// 		//res, _ := c.service.Search(user.Email)
// 		ctx.JSON(http.StatusOK,  _U)
// 		return
// 	}
// 	ctx.JSON(http.StatusBadRequest, gin.H{"error:": "bad request "})
// }


//Search
func (c cartcontroller) Search(ctx *gin.Context) {

	var email service.Email

	err := ctx.ShouldBindJSON(&email)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

		result, err := c.service.Search(contx,email.Email)
		

		if err != nil {
		
				contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
				customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
				ctx.JSON(http.StatusBadRequest, *customerror)
				return
		}		
		ctx.JSON(http.StatusOK, *result)
	
}

func (c cartcontroller) Delete(ctx *gin.Context) {
		

	id := ctx.Param("id")
	contx := ctx.Request.Context()
	
	if id == "" {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	_,err :=c.service.Delete(ctx.Request.Context(), id)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
	return

}

