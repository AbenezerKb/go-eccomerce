package service

import (
	// "context"
	"context"	
	"gin-exercise/db"
	"gin-exercise/entity"

	// "github.com/gin-gonic/gin"
)

type CartService interface {
	Save(context.Context,entity.Cart) (*entity.Cart, error)
	FindAll(context.Context,string,Pagination) (*[]entity.Cart, error)
	Search(context.Context, string) (*entity.Cart, error)
	CartItemCount(ctx context.Context,id string) (int64, error)	
	Delete(context.Context, string)(bool,error)
}

type cartService struct {
	carts []entity.Cart
}

func NewCart() CartService {
	return &cartService{}
}

func (i *cartService) Save(ctx context.Context, cart entity.Cart) (*entity.Cart, error) {

	res,err := db.CartSave(ctx, cart)
	
	if err!=nil {
		return nil, err
	}
	
	return res, nil
}


func (i *cartService) FindAll(ctx context.Context,id string, page Pagination) (*[]entity.Cart,error) {
	page.Page = (page.Size * page.Page) - page.Size
	list,err:=db.CartList(page.Page, page.Size,id, ctx)
	if err!=nil{
		return nil,err
	}
	return list,nil
}

// func (i *userService) FindAll(page Pagination) string {
// 	page.Page = (page.Size * page.Page) - page.Size
// 	return db.List(page.Page, page.Size, "User")
// }

func (i *cartService) Search(ctx context.Context, id string) (*entity.Cart, error) {
	searchResult, err := db.CartSearch(ctx, id, "id")
	if err != nil {
		
		return nil, err
	}
	
	return searchResult, nil
}

func (i *cartService) Delete(ctx context.Context, id string) (bool,error) {

res,err :=	db.Delete(ctx, id, "id")
if err!=nil{
	return res,err
}
return res,nil
}



func (i *cartService) CartItemCount(ctx context.Context,id string) (int64, error){
	return db.CartItemCount(ctx,id)
}