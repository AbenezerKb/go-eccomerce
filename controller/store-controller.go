package controller

import (
	"context"
	"gin-exercise/Errors"
	"gin-exercise/entity"
	"gin-exercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Search(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	CreateStoreRole(ctx *gin.Context)
	UpdateStoreStatus(ctx *gin.Context)
}

type storecontroller struct {
	service service.StoreService
}

var (
	loginService         service.LoginService = &service.UserLogin{}
	jwtService           service.JWTService   = service.JWTAuthService()
	loginValidController LoginController      = LoginHandler(loginService, jwtService)
)

func NewStoreController(service service.StoreService) StoreController {
	return storecontroller{service: service}
}

func (c storecontroller) FindAll(ctx *gin.Context) {

	page := service.Pagination{}

	err := ctx.BindJSON(&page)
	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	list, err := c.service.FindAll(ctx, page)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}
	total,err := c.service.StoreCount(contx)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_FIND_RESOURCE, err)
		customerror := Errors.Error(Errors.UNABLE_TO_FIND_RESOURCE, contx, http.StatusInternalServerError)
		ctx.JSON(http.StatusUnauthorized, *customerror)

		return
	}

	ctx.JSON(http.StatusOK,  gin.H{"meta_data":total,"data":list})
	return

}

func (c storecontroller) Save(ctx *gin.Context) {
	var store entity.Store

	err := ctx.ShouldBindJSON(&store)

	contx := ctx.Request.Context()
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	_U, er := c.service.Save(contx, store)
	ctx.Set("storeID", _U.ID)
	if er == nil {

		userid := ctx.GetString("userID")

		if hasPolicy := service.Enforcer().HasPolicy("Owner", _U.ID, "Item", "read"); !hasPolicy {
			_, err := service.Enforcer().AddPolicy("Owner", _U.ID, "Item", "read")

			if err != nil {
				contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
				customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
				ctx.JSON(http.StatusBadRequest, *customerror)
				return
			}

		}

		if hasPolicy := service.Enforcer().HasPolicy("Owner", _U.ID, "Item", "write"); !hasPolicy {
			_, err := service.Enforcer().AddPolicy("Owner", _U.ID, "Item", "write")

			if err != nil {
				contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
				customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
				ctx.JSON(http.StatusBadRequest, *customerror)
				return
			}

		}

		if hasPolicy := service.Enforcer().HasPolicy("Owner", _U.ID, "Item", "modify"); !hasPolicy {
			_, err := service.Enforcer().AddPolicy("Owner", _U.ID, "Item", "modify")

			if err != nil {
				contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
				customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
				ctx.JSON(http.StatusBadRequest, *customerror)
				return
			}

		}

		if hasPolicy := service.Enforcer().HasPolicy("Owner", _U.ID, "Item", "delete"); !hasPolicy {
			_, err := service.Enforcer().AddPolicy("Owner", _U.ID, "Item", "delete")

			if err != nil {
				contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
				customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
				ctx.JSON(http.StatusBadRequest, *customerror)
				return
			}

		}
		service.Enforcer().AddGroupingPolicy(userid, "Owner", _U.ID)

		ctx.JSON(http.StatusOK, _U)

		return
	}

}


func (c storecontroller) CreateStoreRole(ctx *gin.Context) {
	var storerole entity.StoreRole
	contx := ctx.Request.Context()
	err := ctx.ShouldBindJSON(&storerole)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}

	if hasPolicy := service.Enforcer().HasPolicy(storerole.Role, storerole.StoreID, storerole.Resource, storerole.Operation); !hasPolicy {
		_, err := service.Enforcer().AddPolicy(storerole.Role, storerole.StoreID, storerole.Resource, storerole.Operation)

		if err != nil {
			contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
			customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, *customerror)
			return
		}

	}

	service.Enforcer().AddGroupingPolicy(storerole.ClerkID, storerole.Role, storerole.StoreID)
	return

}





func (c storecontroller) UpdateRole(ctx *gin.Context) {
	var storerole entity.StoreRole
	contx := ctx.Request.Context()
	err := ctx.ShouldBindJSON(&storerole)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}

	if hasPolicy := service.Enforcer().HasPolicy(storerole.Role, storerole.StoreID, storerole.Resource, storerole.Operation); !hasPolicy {
		_, err := service.Enforcer().AddPolicy(storerole.Role, storerole.StoreID, storerole.Resource, storerole.Operation)

		if err != nil {
			contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
			customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, *customerror)
			return
		}

	}

	service.Enforcer().AddGroupingPolicy(storerole.ClerkID, storerole.Role, storerole.StoreID)
	return

}



func (c storecontroller) Update(ctx *gin.Context) {
	var store entity.UpdateStore
	err := ctx.ShouldBindJSON(&store)
	contx := ctx.Request.Context()
	name := ctx.Param("id")
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	_U, err := c.service.Update(contx, name, store)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)

		return
	}

	ctx.JSON(http.StatusOK, *_U)
}

func (c storecontroller) UpdateStoreStatus(ctx *gin.Context) {

	var status service.Status
	err := ctx.ShouldBindJSON(&status)
	contx := ctx.Request.Context()
	
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}
	
	
	
	id := ctx.Param("id")
	

	updatedstore, err := c.service.UpdateStoreStatus(contx, id, status.Status)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}
	
	ctx.JSON(http.StatusOK, *updatedstore)

}

//Search
func (c storecontroller) Search(ctx *gin.Context) {

	var name service.Name
		
	contx := ctx.Request.Context()
	err := ctx.ShouldBindJSON(&name)

	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	result, err := c.service.Search(ctx.Request.Context(), name.Name)
	if err != nil {
		contx = context.WithValue(contx, Errors.UNABLE_TO_READ, err)
		customerror := Errors.Error(Errors.UNABLE_TO_READ, contx, http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, *customerror)
		return
	}

	ctx.JSON(200, result)
}

func (c storecontroller) Delete(ctx *gin.Context) {

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
