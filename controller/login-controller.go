package controller

import (
	// "fmt"
	"fmt"
	"gin-exercise/db"
	"gin-exercise/entity"

	"context"
	"gin-exercise/Errors"
	"gin-exercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context)
	ValidateToken(context.Context, string) bool
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

var c *gin.Context
var u *entity.User	

func (controller loginController) Login(ctx *gin.Context) {
	var credential entity.LoginInfo
	contx := ctx.Request.Context()
	err := ctx.ShouldBind(&credential)

	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror:=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	isUserAuthenticated := controller.loginService.Login(contx, credential.Email, credential.Password)
	
	if isUserAuthenticated {

		token, err := controller.jWtService.GenerateToken(credential.Email)
	
		if err != nil {
	
			contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
			customerror:=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, *customerror)

			return

		} else {
			fmt.Println("the token: ",*token)
		ctx.JSON(http.StatusOK, gin.H{
			
			"token": *token,
		})
		userId, err := db.Search(contx, credential.Email, "email")

		
		if err != nil {
			contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
			customerror:=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, *customerror)
			return
		}

		fmt.Println("user ID:", userId.Id)
		ctx.Set("userID", userId.Id)
		c = ctx
		u = userId

		

	}
}
	ctx.JSON(http.StatusUnauthorized, nil)

}

func (controller loginController) ValidateToken(contx context.Context,token string) bool {
	newToken, err := controller.jWtService.ValidateToken(contx,token)
	if err != nil {				
			//contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
			//Errors.ErrInvalidToken
			// customerror:=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
			// ctx.JSON(http.StatusBadRequest, *customerror)
			// return 
		
		return false
	}

	if newToken != nil {
		c.Set("userID", u.Id)
		return true
	}
	return false
}
