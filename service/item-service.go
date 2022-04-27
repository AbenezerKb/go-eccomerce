package service

import (
	"context"
	"fmt"
	"gin-exercise/Errors"
	"gin-exercise/db"
	"gin-exercise/entity"

	"github.com/joomcode/errorx"

	uuid "github.com/satori/go.uuid"
)

type ItemService interface {
	Save(context.Context, entity.Item) (*entity.Item, error)
	FindAll(context.Context, Pagination) (*[]entity.Item, error)
	Search(context.Context, string) (*entity.Item, error)
	SearchByID(context.Context, string) (*entity.Item, error)
	Delete(context.Context, string) (bool, error)
	FindAllByStoreId(context.Context, string, Pagination) (*[]entity.Item, error)
	ItemCount(context.Context) (int64, error)
	StoreItemCount(context.Context, string) (int64, error)
}

type itemService struct {
	users []entity.Item
}

type ItemName struct {
	Name string `json:"name" binding:"required"`
}

func NewItem() ItemService {
	return &itemService{}
}

func (i *itemService) Save(ctx context.Context, item entity.Item) (*entity.Item, error) {

	if !validateItemFields(item) {
		return nil, fmt.Errorf(Errors.UNABLE_TO_SAVE)
	}

	item.ID = fmt.Sprint(uuid.NewV4())
	item.StoreID = "6"

	db.ItemSave(ctx, item)
	it, err := db.ItemSearch(ctx, item.ID, "id")
	if err != nil {

		return nil, err
	}

	return it, nil

}

func (i *itemService) FindAll(ctx context.Context, page Pagination) (*[]entity.Item, error) {
	list, err := db.ItemList(page.Page, page.Size, ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (i *itemService) SearchByID(ctx context.Context, name string) (*entity.Item, error) {

	searchResult, err := db.ItemSearch(ctx, name, "name")
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

func (i *itemService) Search(ctx context.Context, name string) (*entity.Item, error) {
	searchResult, err := db.ItemSearch(ctx, name, "name")
	if err != nil {
		err = errorx.Decorate(err, "decorate")
		return nil, err

	}

	return searchResult, nil
}

func (i *itemService) Delete(ctx context.Context, id string) (bool, error) {

	searchResult, err := db.ItemDelete(ctx, id, "name")

	if err != nil {
		return searchResult, err
	}
	return searchResult, fmt.Errorf(Errors.UNABLE_TO_READ)
}

func validateItemFields(item entity.Item) bool {
	if len(item.Name) == 0 {
		return false
	}
	return true
}

func (i *itemService) FindAllByStoreId(ctx context.Context, id string, page Pagination) (*[]entity.Item, error) {

	list, err := db.StoreItemList(page.Page, page.Size, id, ctx)
	if err != nil {
		return nil, err
	}
	return list, nil

}

func (i *itemService) ItemCount(ctx context.Context) (int64, error) {
	return db.Count(ctx, entity.Item{})
}

func (i *itemService) StoreItemCount(ctx context.Context, id string) (int64, error) {
	return db.StoreItemCount(ctx, id)
}
