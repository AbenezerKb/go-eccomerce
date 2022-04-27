package controller

import (
	"context"	
	"gin-exercise/Errors"
	"gin-exercise/entity"
	"gin-exercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Search(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ordercontroller struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return ordercontroller{service: service}
}

func (c ordercontroller) FindAll(ctx *gin.Context) {

	page := service.Pagination{}

	err := ctx.BindJSON(&page)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	
	list, err := c.service.FindAll(contx, page)
	total,err := c.service.OrderCount(contx)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}

	ctx.JSON(http.StatusOK,  gin.H{"meta_data":total,"data":list})
	return

}

func (c ordercontroller) Save(ctx *gin.Context) {

	var order entity.Order	

	err := ctx.BindJSON(&order)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}


	savedorder, err := c.service.Save(contx,order)


	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)

	}
	
	ctx.JSON(http.StatusOK, *savedorder)
	return
}

func (c ordercontroller) Update(ctx *gin.Context) {




	var order entity.UpdateOrder

	var id string
	contx := ctx.Request.Context()

	err := ctx.ShouldBindJSON(&order)

	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}

	id = ctx.Param("id")

	order.ID = id
	
	if id == "" {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror )

		return
	}

	updatedorder, err := c.service.Update(contx,id, order)



	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_SAVE, err)
		customerror :=Errors.Error(Errors.UNAUTHORIZED, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusInternalServerError, *customerror)

		return
	}

	ctx.JSON(http.StatusOK, *updatedorder)
	return


}

//Search
func (c ordercontroller) Search(ctx *gin.Context) {

	orderID := ctx.Param("orderid")
	contx := ctx.Request.Context()
	if orderID != "" {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}


	result, err := c.service.Search(contx,orderID)
	
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusInternalServerError, *customerror)
		return
	
	}
	
	
	if result.ID == "" {
		ctx.JSON(204, entity.Order{})		
		return
	}
	ctx.JSON(200, *result)
	
}

func (c ordercontroller) Delete(ctx *gin.Context) {

	contx := ctx.Request.Context()
	id := ctx.Param("id")

	if id == "" {

		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror:=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}

	c.service.Delete(contx,id)
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
