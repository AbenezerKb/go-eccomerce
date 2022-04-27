package middlewares

import (
	"gin-exercise/Errors"
	//"context"
	//"gin-exercise/Errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
)


func Error_Handling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		
		for _,err := range ctx.Errors{
			var msg string
		status :=http.StatusInternalServerError
		if 	errorx.HasTrait(err,errorx.Timeout()){
			status = http.StatusRequestTimeout
			msg = "request_timeout"
		}
		
		if 	errorx.HasTrait(err,errorx.Duplicate()){
			status = http.StatusBadRequest
			msg = "duplicate_value_found"
		}
				//TODO ERROR HANDLING FOR INVALID INPUT
			//if 	errorx.HasTrait(err,errorx.){
			//	status = http.StatusBadRequest
			//	msg = "invalid_input"
			//}


			if 	errorx.HasTrait(err,errorx.RegisterTrait(Errors.UNABLE_TO_SAVE)){
			status = http.StatusInternalServerError
			msg = "unable_to_save"
		}


			ctx.JSON(status,Errors.ErrorHandling(msg))

		}



	}
}