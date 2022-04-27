package service

import (
	
	"gin-exercise/db"
	"gin-exercise/entity"
	
	"context"
)

type OrderService interface {
	Save(context.Context, entity.Order) (*entity.Order, error)
	FindAll(context.Context, Pagination) (*[]entity.Order,error)
	Search(context.Context, string) (*entity.Order, error)
	Update(context.Context, string, entity.UpdateOrder) (*entity.Order, error) 
	Delete(context.Context, string)  (bool,error)    
	OrderCount(context.Context) (int64, error)                           
}

type orderService struct {
	orders []entity.Order
}

func NewOrder() OrderService {
	return &orderService{}
}

func (i *orderService) Save(ctx context.Context,order entity.Order) (*entity.Order, error) {

	res,err := db.OrderSave(ctx, order)	
	if err != nil {
	
		return nil, err
	}	

	return res, nil
}

func (i *orderService) Update(ctx context.Context, id string, order entity.UpdateOrder) (*entity.Order, error) { 

	searchResult, err := db.UpdateOrder(ctx,id, order)

	if err != nil {
		return nil, err
	}
	return searchResult, nil
}

func (i *orderService) FindAll(ctx context.Context,page Pagination) (*[]entity.Order,error) {
	page.Page = (page.Size * page.Page) - page.Size
	list,err := db.OrderList(page.Page, page.Size,ctx)
	if err!=nil{
		return nil,err
	}

	return list,nil
}


func (i *orderService) Search(ctx context.Context, id string) (*entity.Order, error) {

	searchResult, err := db.OrderSearch(ctx,id, "id")
	if err != nil {
		
		return nil, err
	}
	
	return searchResult, nil
}

func (i *orderService) Delete(ctx context.Context, id string)(bool, error) {
	
	res,err := db.Delete(ctx, id, "id")
	if err !=nil{
		return res,err
	}

	return res,nil

}


func (i *orderService)OrderCount(ctx context.Context) (int64, error){
	return db.Count(ctx,entity.Order{})
}
