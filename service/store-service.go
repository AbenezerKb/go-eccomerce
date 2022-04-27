package service

import (
	// "context"
	
	"fmt"
	"gin-exercise/db"
	"gin-exercise/entity"
	"gin-exercise/Errors"

	"context"
)

type StoreService interface {
	Save(context.Context, entity.Store) (*entity.Store, error)
	FindAll(context.Context, Pagination) (*[]entity.Store,error)
	Search(context.Context, string) (*entity.Store, error)
	Update(context.Context, string, entity.UpdateStore) (*entity.Store, error) 
	UpdateStoreStatus(context.Context, string, string) (*entity.Store, error)
	Delete(context.Context, string)  (bool,error)  
	StoreCount(context.Context) (int64, error)  	
}

type Name struct {
	Name string `json:"name" binding:"required"`
}

type Status struct {
	Status string `json:"status" binding:"required"`
}

type ID struct {
	Id string `json:"id" binding:"required"`
}

type storeService struct {
	stores []entity.Store
}

func NewStore() StoreService {
	return &storeService{}
}

func (i *storeService) Save(ctx context.Context, store entity.Store) (*entity.Store, error) {

	if !FieldValidation(store) {
		// ctx.Error(gin.Error{})			
		return nil, fmt.Errorf(Errors.UNABLE_TO_SAVE)

	}

	res,err := db.StoreSave(ctx, store)
	if err!=nil{
		return nil,err
	}
	return res, nil

}

func (i *storeService) Update(ctx context.Context, id string, store entity.UpdateStore) (*entity.Store, error) { 

	searchResult, err := db.UpdateStore(ctx, id, store)

	if err != nil {

		return nil, err

	}
	return searchResult, nil
}

func (i *storeService) UpdateStoreStatus(ctx context.Context, id string, status string) (*entity.Store, error) { //([]byte, error)

	searchResult, err := db.UpdateStoreStatus(ctx, id, status)

	if err != nil {
		
		return nil, err
	}
	return searchResult, nil
}

func (i *storeService) FindAll(ctx context.Context, page Pagination) (*[]entity.Store,error) {
	
	list,err :=db.StoreList( page.Page, page.Size, ctx)
	
	
	if err!=nil{
		return nil,err
	}
		
	return list,nil

}

// func (i *userService) FindAll(page Pagination) string {
// 	page.Page = (page.Size * page.Page) - page.Size
// 	return db.List(page.Page, page.Size, "User")
// }

func (i *storeService) Search(ctx context.Context, name string) (*entity.Store, error) {
	searchResult, err := db.StoreSearch(ctx, name, "name")

	if err != nil {
	
		return searchResult, err
	}
		
	return searchResult, nil
}

func (i *storeService) Delete(ctx context.Context, id string)(bool,error) { //error {

	res,err :=db.StoreDelete(ctx, id, "id")
	if err!=nil{
		return res,err
	}
	return res,nil

}

func FieldValidation(store entity.Store) bool {
	if store.Location == "" || store.Name == "" || store.Image == "" {
		return false
	}
	return true
}


func (i *storeService)StoreCount(ctx context.Context) (int64, error){
	return db.Count(ctx,entity.Store{})
}

