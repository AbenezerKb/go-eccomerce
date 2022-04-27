package middlewares

import (
	//"JWT-auth/src/service"
	"context"
	"fmt"
	"gin-exercise/Errors"
	"gin-exercise/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	loginService service.LoginService = &service.UserLogin{}
	jwtService   service.JWTService   = service.JWTAuthService()
)

func AuthorizeJWT() gin.HandlerFunc {
	
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]		
		contx := ctx.Request.Context()
		token, err := service.JWTAuthService().ValidateToken(contx,tokenString)	


		if err != nil {

			contx = context.WithValue(contx, Errors.UNAUTHORIZED, err)
			Errors.Error(Errors.UNAUTHORIZED, contx, http.StatusUnauthorized)
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}



		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)			

			ctx.Set("userID", claims["issuser"])
			return
			//  ctx.Keys["userID"]= claims.ID
		} else {
			
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

// loginValidController.ValidateToken(ctx.GetHeader("token"))

func IsLogged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contx:=	ctx.Request.Context()
		fmt.Println("the token ", ctx.GetHeader("token"))
		newToken, err := jwtService.ValidateToken(contx,ctx.GetHeader("token"))
		if err != nil {

			contx = context.WithValue(contx, Errors.UNAUTHORIZED, err)
			Errors.Error(Errors.UNAUTHORIZED, contx, http.StatusUnauthorized)
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		if newToken == nil {

			contx = context.WithValue(contx, Errors.UNAUTHORIZED, err)
			Errors.Error(Errors.UNAUTHORIZED, contx, http.StatusUnauthorized)
			ctx.JSON(http.StatusUnauthorized, nil)
			return

		}

		ctx.Next()

	}
}

// func LogginChecker() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	// token string
// 	// token :=ctx.GetHeader("token")
// 	// var ser = service.
// 	// newToken, err := service.JWTService.ValidateToken(token)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// ctx.Next()

// 	// ctx

// }
// }

// func ValidateToken(token string) bool {
// 	newToken, err := controller.jWtService.ValidateToken(token)
// 	if err != nil {
// 		return false
// 	}

// 	if newToken != nil {
// 		c.Set("userID", u.Id)
// 	}
// 	return true
// }
