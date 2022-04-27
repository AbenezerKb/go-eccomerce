package controller

import (
	"context"	
	"gin-exercise/Errors"
	"gin-exercise/entity"
	"gin-exercise/service"
	"net/http"
	"github.com/joomcode/errorx"
	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Search(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type itemcontroller struct {
	service service.ItemService
}

func NewItemController(service service.ItemService) ItemController {
	return itemcontroller{service: service}
}

func (c itemcontroller) FindAll(ctx *gin.Context) {

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
	total, err := c.service.ItemCount(contx)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meta_data": total, "data": list})
	return

}

func (c itemcontroller) FindAllById(ctx *gin.Context) {

	page := service.Pagination{}

	err := ctx.BindJSON(&page)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	store_id := ctx.Param("store_id")

	if store_id == "" {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, Errors.UNABLE_TO_READ)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	list, err := c.service.FindAllByStoreId(contx, store_id, page)
	total, err := c.service.StoreItemCount(contx, store_id)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meta_data": total, "data": list})
	return


}

func (c itemcontroller) Delete(ctx *gin.Context) {

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

func (c itemcontroller) Save(ctx *gin.Context) {
	var item entity.Item
	err := ctx.ShouldBindJSON(&item)	
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	saveditem,err := c.service.Save(contx,item)

	if err != nil {

		contx = context.WithValue(contx, Errors.UNABLE_TO_SAVE, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_SAVE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusInternalServerError, *customerror)
		return
	}
	
	ctx.JSON(http.StatusOK, *saveditem)
	
}

//Search
func (c itemcontroller) Search(ctx *gin.Context) {

	var itemname service.ItemName

	err := ctx.ShouldBindJSON(&itemname)	
	
	contx := ctx.Request.Context()
	if err != nil {
		err = errorx.Decorate(err, "decorate")
		
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}
	result, err := c.service.Search(contx,itemname.Name)

	if err != nil {
	
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror :=Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}
	
	ctx.JSON(http.StatusOK, *result)

}
