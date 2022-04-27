package service

import (
	
	"gin-exercise/db"
	"gin-exercise/entity"
	"context"
	
)

type CategoryService interface {
	Save(context.Context,string) (*entity.Category, error)
	FindAll(context.Context, Pagination) (*[]entity.Category,error)
	CategoryCount(ctx context.Context) (int64, error)	
}

type  categoryService struct {
	category []entity.Category
}

func NewCategory() CategoryService {
	return &categoryService{}
}

func (i *categoryService) Save(ctx context.Context,category string) (*entity.Category, error) {

		cat,err:= db.CategorySave(ctx,category)	
		if err!=nil{
			return nil,err
		}
		return cat, nil
		
}

func (i *categoryService) FindAll(ctx context.Context, page Pagination) (*[]entity.Category, error) {
	page.Page = (page.Size * page.Page) - page.Size
	list,err:=db.CategoryList(page.Page, page.Size, ctx)
	if err!=nil{
		return nil,err
	}
	return list,nil
}



func (i *categoryService) CategoryCount(ctx context.Context) (int64, error){
	return db.Count(ctx,entity.Category{})
}


