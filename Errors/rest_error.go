package Errors

import (
	"context"
	"fmt"
	// "net/http"
	"github.com/joomcode/errorx"

)

type RestErr struct {
	Message string	
	Describtion   string
}

var (
	UNABLE_TO_FIND_RESOURCE = "UNABLE_TO_FIND_RESOURCE"
	UNABLE_TO_READ          = "UNABLE_TO_READ"
	UNAUTHORIZED            = "UNAUTHORIZED"
	UNABLE_TO_SAVE          = "UNABLE_TO_SAVE"
)

var (
	//DBErrors        = errorx.NewNamespace("db")
	ErrInvalidToken    = errorx.NewType(errorx.CommonErrors, "invalid_token")
	ErrValidationError = errorx.NewType(errorx.CommonErrors, "invalid_input")
	ErrDuplicateData   = errorx.AssertionFailed.NewSubtype("duplicate_value", errorx.Duplicate())
	//ErrDBError = DBErrors.NewType("db_error")

)



func Error(message string, ctx context.Context, status int) *RestErr {
	return &RestErr{		
		Message: message,	
		Describtion:   fmt.Sprint(ctx.Value(message)),
	}
}


func ErrorHandling(msg string) *RestErr {

	var desc string
	var message string

	if msg == "invalid_token"{
	   message = "invalid token"
	   desc = "you have used invalid token"
	}
	if msg == "invalid_input"{
		message = "invalid token"
		desc = "you have used invalid token"
	}
	if msg == "duplicate_value"{
		message = "invalid token"
		desc = "you have used invalid token"
	}


	return &RestErr{
		Message: message,	
		Describtion:   desc,
	}
}





// func Unable_to_read(message string) *RestErr {
// 	return &RestErr{
// 		Message: message,
// 		Status:  http.StatusBadRequest,
// 		Error:   UNABLE_TO_READ,
// 	}
// }

// func Unauthorized(message string) *RestErr {
// 	return &RestErr{
// 		Message: message,
// 		Status:  http.StatusUnauthorized,
// 		Error:   UNAUTHORIZED,
// 	}
// }

// func Unable_to_find_resource(message string) *RestErr {
// 	return &RestErr{
// 		Message: message,
// 		Status:  http.StatusInternalServerError,
// 		Error:   UNABLE_TO_FIND_RESOURCE,
// 	}
// }
