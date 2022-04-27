package service

import (
	//	"fmt"
	// "context"
	"encoding/json"
	"fmt"	
	"gin-exercise/Errors"
	"gin-exercise/db"
	"gin-exercise/entity"

	uuid "github.com/satori/go.uuid"

	"io"
	"strings"

	"log"
	"os"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Save(context.Context, entity.User) (*entity.User, error)
	FindAll(context.Context, Pagination) (*[]entity.User,error)
	Search(context.Context, string) ([]byte, error)
	Update(context.Context, string, entity.UpdateUser) (*entity.User, error)
	Delete(context.Context, string) (bool,error)    
	UserCount(context.Context) (int64, error)
}

type userService struct {
	users []entity.User
}

type Pagination struct {
	Page int `json:"page" binding:"required"`
	Size int `json:"size" binding:"required"`
	Total int `json:"total"`
	
}

type Email struct {
	Email string `json:"email" binding:"email"`
}

func NewUser() UserService {
	return &userService{}
}

func (i *userService) Save(ctx context.Context, user entity.User) (*entity.User, error) {

	if !validateUserFields(user) {
		
		return nil, fmt.Errorf(Errors.UNABLE_TO_READ)
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		
		return nil, fmt.Errorf(Errors.UNABLE_TO_SAVE)
	}
	user.Password = string(pw[:])	
	user.Id = fmt.Sprint(uuid.NewV4())
	user.ImageURL = "movefile(user.ImageURL)"	
	_,err = Enforcer().AddGroupingPolicy(user.Id, "Customer", "read")	

	res,err := db.UserSave(ctx, user)
	
	if err!=nil {
		return nil, err
	}

	return res, fmt.Errorf(Errors.UNABLE_TO_SAVE)
}

func (i *userService) FindAll(ctx context.Context, page Pagination) (*[]entity.User,error) {
	page.Page = (page.Size * page.Page) - page.Size
	list, err:=db.UsersList(page.Page, page.Size,ctx)
	return list,err
}

func (i *userService) Search(ctx context.Context, email string) ([]byte, error) {

	searchResult, err := db.Search(ctx, email, "email")

	if err != nil {
		fmt.Println("error in search result")
		return nil, err
	}
	jsonUser, err := json.Marshal(searchResult)

	if err != nil {
		fmt.Println("error in marshaling search result")
		return nil, err
	}

	return jsonUser, nil
}

func (i *userService) Delete(ctx context.Context, id string) (bool,error)    { //bool { //([]byte, error)

	res,err:=db.Delete(ctx, id, "id")
	if err!=nil{
	return res,err
	}
return res,nil

}

func (i *userService) Update(ctx context.Context, id string, user entity.UpdateUser) (*entity.User, error) { //([]byte, error)

	searchResult, err := db.Update(ctx, id, user)

	if err != nil {
		return nil, err
	}
	return searchResult, nil

}

func validateUserFields(user entity.User) bool {
	if len(user.FirstName) == 0 || len(user.SecondName) == 0 || len(user.LastName) == 0 || len(user.Email) == 0 || len(user.PhoneNumber) == 0 || len(user.Password) == 0 {

		fmt.Println("invalid user fields")
		return false
	}
	return true
}

func movefile(filepath string) string {
	original, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()
	// fmt.Println("the image path", filepath)
	// fmt.Println("the last index",filepath[strings.LastIndex(filepath,`\`)+1:])

	PERMANENT_DIR := `C:\Users\Administrator\Documents\New folder (3)\abenezer-dev-prep\perm`
	// Create new file
	new, err := os.Create(PERMANENT_DIR + `\` + filepath[strings.LastIndex(filepath, `\`)+1:])
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	//This will copy
	_, err = io.Copy(new, original)
	if err != nil {
		log.Fatal(err)
	}
	return PERMANENT_DIR + `\` + filepath[strings.LastIndex(filepath, `\`)+1:]
}


func  (i *userService) UserCount(ctx context.Context) (int64, error){
	return db.Count(ctx,entity.User{})
}




