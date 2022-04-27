package db

//TODO error import cycle created with entity.User

import (
	"context"
	// "encoding/json"
	// "errors"
	"fmt"
	"gin-exercise/entity"

	"gin-exercise/Errors"
	// "Errors"

	"reflect"
	"regexp"
	"strings"
	"github.com/joomcode/errorx"
	// "os"	
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/postgres"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//	"gorm.io/gorm/logger"
)

//GORM database connection
func ConnectDB() (db *gorm.DB, err error) {
	// db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	// .WithContext()
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=abeny dbname=Exersice port=5432",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true, //30% performance increases
		//Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {

		// Errors.Unable_to_read(err)
		return
	}
	// &entity.Item{},
	db.AutoMigrate(&entity.User{}, &entity.Item{}, &entity.Category{}, &entity.Store{}, &entity.Order{}, &entity.Charge{}, &entity.Cart{})

	return
}

//Lists

func ItemList(page, size int, ctx context.Context) (*[]entity.Item,error) {
	DbPool, err := ConnectDB()
	if err != nil {			
	return nil,err
	}

	var items []entity.Item
	DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&items)
	
	return &items,nil
}



func StoreItemList(page, size int, id string,ctx context.Context) (*[]entity.Item,error) {
	DbPool, err := ConnectDB()
	if err != nil {			
	return nil,err
	}

	var items []entity.Item
	DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Where("StoreID",id).Find(&items)
	
	return &items,nil
}







func UsersList(page, size int, ctx context.Context) (*[]entity.User, error) {
	DbPool, err := ConnectDB()
	if err != nil {				
		return nil,err
	}

	var users []entity.User
	DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&users)

	return &users, nil
}


func StoreList(page, size int, ctx context.Context) (*[]entity.Store, error) {
	DbPool, err := ConnectDB()
	if err != nil {				
		return nil,err
	}

	var stores []entity.Store
	DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&stores)

	return &stores, nil
}



func CategoryList(page, size int, ctx context.Context) (*[]entity.Category, error) {
	DbPool, err := ConnectDB()
	if err != nil {				
		return nil,err
	}

	var categories []entity.Category
	DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&categories)

	return &categories, nil
}



func OrderList(page, size int, ctx context.Context) (*[]entity.Order, error) {
	DbPool, err := ConnectDB()
	if err != nil {				
		return nil,err
	}

	var orders []entity.Order
	DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&orders)

	return &orders, nil
}



func CartList(page, size int,id string, ctx context.Context) (*[]entity.Cart, error) {
	DbPool, err := ConnectDB()
	if err != nil {				
		return nil,err
	}

	var carts []entity.Cart
	DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Where("OrderOwner",id).Find(&carts)

	return &carts, nil
}




// func List(page, size int, model string, ctx context.Context) *[]entity.User {

// 	DbPool, err := ConnectDB()
// 	if err != nil {
// 	//ctx *gin.Context,
// 		Errors.Unable_to_read(err.Error())
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}

// 	var List []byte
// 	var Err error

// 	switch model {
// 	case "User":
// 		var users []entity.User
// 		DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&users)
// 		List, Err = json.Marshal(users)
// 		if Err != nil {
// 		//ctx *gin.Context,
// 			Errors.Unable_to_read(err.Error())
// 			fmt.Println("marshaling users list error!")
// 			os.Exit(1)
// 		}
// 	case "Store":
// 		var stores []entity.Store
// 		DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&stores)
// 		fmt.Println(stores)
// 		List, Err = json.Marshal(stores)
// 		if Err != nil {
// 			//ctx.Error(err)
// 			Errors.Unable_to_read(err.Error())
// 			return ""
// 		}
// 	case "Category":
// 		var categories []entity.Category
// 		DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&categories)
// 		fmt.Println(categories)
// 		List, Err = json.Marshal(categories)
// 		if Err != nil {
// 			//ctx.Error(err)
// 			Errors.Unable_to_read(err.Error())
// 			return ""

// 		}
// 	case "Item":
// 		var items []entity.Item
// 		DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&items)
// 		n, _ := Count(ctx, entity.Item{})
// 		fmt.Println("the count", n)

// 		List, Err = json.Marshal(items)
// 		if Err != nil {
// 			//ctx.Error(err)
// 			Errors.Unable_to_read(err.Error())
// 			return ""

// 		}
// 	case "Order":
// 		var orders []entity.Order
// 		DbPool.WithContext(ctx).Offset((page - 1) * size).Limit(size).Find(&orders)
// 		fmt.Println(orders)
// 		List, Err = json.Marshal(orders)
// 		if Err != nil {
// 			//ctx.Error(err)
// 			Errors.Unable_to_read(err.Error())
// 			return ""

// 		}

// 	}
// 	return string(List)
// }

func UserSave(ctx context.Context, user entity.User) (*entity.User,error) {

	DbPool, err := ConnectDB()
	if err != nil {
		
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())		
		// os.Exit(1)
		return nil,err
	}


	user.Status = "pending"
	var evalid EmailValidator
	var pvalid PhoneValidator

	par1, _ := evalid.Validate(ctx, user.Email)

	par2, _ := pvalid.Validate(ctx, string(user.PhoneNumber))
	
	if par1 && par2 {

		DbPool.WithContext(ctx).Create(&user) 
		return &user,nil
	}

	return nil,fmt.Errorf(Errors.UNABLE_TO_SAVE)
}

func StoreSave(ctx context.Context, store entity.Store) (*entity.Store,error) {
	DbPool, err := ConnectDB()
	if err != nil {
		//ctx.Error(err)
		//Errors.Unable_to_read(err.Error())		
		return nil,err		
	}
	store.ID = fmt.Sprint(uuid.NewV4())

	DbPool.WithContext(ctx).Create(&store)
	return &store,nil
}

func OrderSave(ctx context.Context, order entity.Order) (*entity.Order,error) {
	DbPool, err := ConnectDB()
	if err != nil {
		
		return nil,err		
	}
	order.ID = fmt.Sprint(uuid.NewV4())
	DbPool.WithContext(ctx).Create(&order)
	return &order,nil
}

func CartSave(ctx context.Context, cart entity.Cart) (*entity.Cart,error) {
	DbPool, err := ConnectDB()
	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		// os.Exit(1)
		return nil,err
	}

	DbPool.WithContext(ctx).Create(&cart)
	return &cart,nil
}


func CategorySave(ctx context.Context, name string) (*entity.Category,error) {
	DbPool, err := ConnectDB()
	if err != nil {
	//ctx.Error(err)
	// Errors.Unable_to_read(err.Error())
	return nil,err	
	}

	var newcategory entity.Category
	newcategory.Name = name
	DbPool.WithContext(ctx).Create(&newcategory)	
	return &newcategory, nil
}

func ItemSave(ctx context.Context, item entity.Item)(*entity.Item,error) {

	DbPool, err := ConnectDB()
	if err != nil {	
		fmt.Println("the error: ",err)
		return nil,err
	}
//WithContext(ctx).
	err = DbPool.Create(&item).Error
	if err != nil {	
	
		return nil,err
	}
	
	return &item,nil
}

//User Registetion
// func SaveUser(user entity.User) bool {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		fmt.Println("saving failed")
// 		return false
// 	}
// 	user.Id = fmt.Sprint(uuid.NewV4())
// 	user.Status = "pending"
// 	var evalid EmailValidator
// 	var pvalid PhoneValidator
// 	par1, _ := evalid.Validate(user.Email)

// 	par2, _ := pvalid.Validate(string(user.PhoneNumber))

// 	if par1 && par2 {

// 		db.Create(&user) // pass pointer of data to Create
// 		return true
// 	}
// 	return false

// }

// func Save(store entity.Store) bool {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		fmt.Println("saving failed")
// 		return false
// 	}
// 	db.Create(&store) // pass pointer of data to Create
// 	return true
// }

//Login handler
func UserInfo(ctx context.Context, email string, password string) (string, bool) {

	// DbPool, err := ConnectDB()

	// var user entity.User

	// if err != nil {
	// 	fmt.Println("database connection error!")
	// 	return false
	// }
	// fmt.Println("before search")
	// DbPool.Where("email = ? ", email).Find(&user)
	// if user.FirstName == "" {
	// 	fmt.Println("empty user")
	// 	return false
	// }
	// fmt.Println("before comparison")
	// fmt.Println("password: ",password)
	// paswordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if paswordError != nil {
	// 	fmt.Println("password confirmation error")
	// 	return false
	// }

	// fmt.Println("after comparison")
	// return true

	DbPool, err := ConnectDB()

	var user entity.User

	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())		
		return "", false
	}

	DbPool.WithContext(ctx).Where("email = ? ", email).Find(&user)
	if user.FirstName == "" {
		//ctx.Error(errors.New(Errors.UNABLE_TO_FIND_RESOURCE))
	//	Errors.Unable_to_find_resource(err.Error())		
		return "", false
	}
	
	paswordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if paswordError != nil {
		//ctx.Error(paswordError)
	//	Errors.Unable_to_save(err.Error())		
		return "", false
	}

	return user.Id, true

}

func Update(ctx context.Context, id string, user entity.UpdateUser) (*entity.User, error) {

	DbPool, err := ConnectDB()
	var searchuser entity.User
	if err := DbPool.WithContext(ctx).Where("id = ?", id).First(&searchuser).Error; err != nil {
		touser := entity.User(user)
		return &touser, fmt.Errorf("user doesn't exist")
	}

	if err != nil {
		//ctx.Error(err)
	//	Errors.Unable_to_save(err.Error())		
		return nil, err
	}

	DbPool.WithContext(ctx).Model(&searchuser).Where("id = ?", id).Updates(&user)

	return nil, fmt.Errorf("user doesn't exist")
}

func UpdateStore(ctx context.Context, id string, store entity.UpdateStore) (*entity.Store, error) {

	DbPool, err := ConnectDB()

	if err != nil {
		//ctx.Error(err)
	//	Errors.Unable_to_save(err.Error())		
		return nil, err
	}
	var searchstore entity.Store
	thesearcherr := DbPool.WithContext(ctx).First(&searchstore, "id = ?", id).Error

	if thesearcherr != nil {
		touser := entity.Store(store)
		return &touser, fmt.Errorf(`store doesn't exist %v`, err)
	}

	DbPool.WithContext(ctx).Model(&searchstore).Where("id = ?", id).Updates(&store)
	DbPool.WithContext(ctx).First(&searchstore, "id = ?", id)
	return &searchstore, nil
}

func UpdateStoreStatus(ctx context.Context, id string, status string) (*entity.Store, error) {

	DbPool, err := ConnectDB()

	if err != nil {
		
		return nil, err
	}
	var searchstore entity.Store
	thesearcherr := DbPool.WithContext(ctx).First(&searchstore, "id = ?", id).Error

	if thesearcherr != nil {		

		return nil, thesearcherr
	}

	DbPool.WithContext(ctx).Model(&searchstore).Where("id = ?", id).Update("status", status)
	DbPool.WithContext(ctx).First(&searchstore, "id = ?", id)
	return &searchstore, nil

}

func UpdateOrder(ctx context.Context, id string, order entity.UpdateOrder) (*entity.Order, error) {

	DbPool, err := ConnectDB()

	if err != nil {
		
		return nil, err
	}

	var searchorder entity.Order
	thesearcherr := DbPool.WithContext(ctx).First(&searchorder, "id = ?", id).Error

	if thesearcherr != nil {
		
		return nil, thesearcherr
	}

	DbPool.WithContext(ctx).Model(&searchorder).Where("id = ?", id).Updates(&order)
	DbPool.WithContext(ctx).First(&searchorder, "id = ?", id)
	return &searchorder, nil

}

func Delete(ctx context.Context, field string, param string) (bool,error){ //bool {

	DbPool, err := ConnectDB()

	var user entity.User

	if err != nil {
		
		return false, err

	}
	DbPool.Where(param+" = ? ", field).Delete(&user)
	return true, nil

}

func ItemDelete(ctx context.Context, field string, param string) (bool,error) {

	DbPool, err := ConnectDB()

	var item entity.Item

	if err != nil {
	
		return false,err
	}
	DbPool.WithContext(ctx).Where(""+param+" = ? ", field).Delete(&item)
	_, notfound := ItemSearch(ctx, field, param)
	if notfound != nil {

		return true,nil
	}

	return false,fmt.Errorf(Errors.UNABLE_TO_READ)
}

func StoreDelete(ctx context.Context, field string, param string) (bool,error) {

	DbPool, err := ConnectDB()

	var store entity.Store

	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		return false,err
	}
	DbPool.WithContext(ctx).Where(param+" = ? ", field).Delete(&store)
	_, notfound := StoreSearch(ctx, field, param)
	if notfound != nil {

		return true,nil
	}

	return false,fmt.Errorf(Errors.UNABLE_TO_READ)
}

func OrderDelete(ctx context.Context, field string, param string) (bool,error) {

	DbPool, err := ConnectDB()

	var item entity.Order

	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		return false,err
	}
	DbPool.WithContext(ctx).Where(""+param+" = ? ", field).Delete(&item)
	_, notfound := OrderSearch(ctx, field, param)
	if notfound != nil {

		return true,nil
	}

	return false,fmt.Errorf(Errors.UNABLE_TO_READ)
}

//search user using email
func Search(ctx context.Context, field string, param string) (*entity.User, error) {

	DbPool, err := ConnectDB()

	var user entity.User

	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())		
		return nil, err
	}
	DbPool.WithContext(ctx).Where(""+param+" = ? ", field).Find(&user)
	if user.Id != "" {
		// ctx.Error(errors.New(Errors.UNABLE_TO_FIND_RESOURCE))
		// Errors.Unable_to_find_resource(err.Error())
		return &user, nil
	}
	return nil, fmt.Errorf(Errors.UNABLE_TO_FIND_RESOURCE)

}

func ItemSearch(ctx context.Context, field string, param string) (*entity.Item, error) {

	DbPool, err := ConnectDB()

	var item entity.Item

	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		return nil, err
	}
	DbPool.WithContext(ctx).Where(""+param+" = ? ", field).Find(&item)
	// DbPool.WithContext(ctx).First(&searchstore, "id = ?", id)
	
	if fmt.Sprint(item.ID) != "" {
		// ctx.Error(errors.New(Errors.UNABLE_TO_FIND_RESOURCE))
		// Errors.Unable_to_find_resource(err.Error())
		return &item, nil
	}
	return nil, DbPool.Error  //"item doesn't exist")

}
func OrderSearch(ctx context.Context, field string, param string) (*entity.Order, error) {

	DbPool, err := ConnectDB()

	var order entity.Order

	if err != nil {		
		err = errorx.Decorate(err, "decorate")
		return nil, err
	}
	DbPool.WithContext(ctx).Where(param+" = ? ", field).Find(&order)
	if fmt.Sprint(order.ID) != "" {	
		return &order, nil
	}
	return nil, fmt.Errorf(Errors.UNABLE_TO_FIND_RESOURCE)

}

func StoreSearch(ctx context.Context, field string, param string) (*entity.Store, error) {

	DbPool, err := ConnectDB()

	var store entity.Store

	if err != nil {
		//ctx.Error(err)
		// Errors.Error(err.Error())
		return nil, err
	}
	DbPool.WithContext(ctx).Where(param+" = ? ", field).Find(&store)
	if fmt.Sprint(store.ID) != "" {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		return& store, nil
	}

	return nil, fmt.Errorf("item doesn't exist")

}

func CartSearch(ctx context.Context, field string, param string) (*entity.Cart, error) {

	DbPool, err := ConnectDB()

	var cart entity.Cart

	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		return nil, err
	}
	DbPool.WithContext(ctx).Where(param+" = ? ", field).Find(&cart)
	if fmt.Sprint(cart.ID) != "" {
		// ctx.Error(errors.New(Errors.UNABLE_TO_FIND_RESOURCE))
		// Errors.Unable_to_find_resource(err.Error())
		return &cart, nil
	}
	return nil, fmt.Errorf("item doesn't exist")

}

const tagName = "validate"

var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

//var phoneRe = regexp.MustCompile("^\\+[1-9]?[0-9]{7,14}$")

type Validator interface {
	Validate(context.Context, interface{}) (bool, error)
}

type DefaultValidator struct {
}

func (v DefaultValidator) Validate(ctx context.Context, val interface{}) (bool, error) {
	return true, nil
}

type EmailValidator struct {
}

func (e EmailValidator) Validate(ctx context.Context, val interface{}) (bool, error) {
	if !mailRe.MatchString(val.(string)) {
		// ctx.Error(errors.New(Errors.UNABLE_TO_SAVE))
		// Errors.Unable_to_read(Errors.UNABLE_TO_SAVE)
		return false, fmt.Errorf("invalid email")
	}
	if result, err := Search(ctx, val.(string), "email"); err == nil || len(result.FirstName) != 0 {

		return false, fmt.Errorf("user by this email already exist")
	}

	return true, nil
}

//Phone validation

type PhoneValidator struct {
}

var digitRe = regexp.MustCompile(`\d{8}`)

//(p PhoneValidator)
func ValidateETHPhone(c context.Context, val string) (bool, error) {
	if val[:2] == "09" {
		if digitRe.MatchString(val[2:]) {
			return true, nil
		}
	}
	if val[:1] == "9" {
		if digitRe.MatchString(val[1:]) {
			return true, nil
		}
	}
	if val[:4] == "+251" {
		if string(val[4]) == "9" {
			if digitRe.MatchString(val[2:]) {
				return true, nil
			}
		}
	}
	if val[:3] == "251" {
		if string(val[3]) == "9" {
			if digitRe.MatchString(val[2:]) {
				return true, nil
			}
		}
	}
	return false, nil

}

func (p PhoneValidator) Validate(ctx context.Context, val interface{}) (bool, error) {

	_, err := ValidateETHPhone(ctx, val.(string))
	if err != nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		return false, nil
	}

	_, errr := Search(ctx, val.(string), "phone_number")
	if errr == nil {
		//ctx.Error(err)
		// Errors.Unable_to_read(err.Error())
		return false, nil
	}
	return true, nil

}


func StoreItemCount(ctx context.Context, id string) (int64, error) {
	var count int64
	var item entity.Item
	db, _ := ConnectDB()
	db.WithContext(ctx).Model(&item).Where("StoreID",id).Count(&count)
	return count, nil
}

func CartItemCount(ctx context.Context, id string) (int64, error) {
	var count int64
	var item entity.Cart
	db, _ := ConnectDB()
	db.WithContext(ctx).Model(&item).Where("OrderOwner",id).Count(&count)
	return count, nil
}






func Count(ctx context.Context, val interface{}) (int64, error) {
	switch x := val.(type) {
	case entity.User:
		var count int64
		db, _ := ConnectDB()
		db.WithContext(ctx).Model(&x).Count(&count)
		return count, nil
	case entity.Store:
		var count int64
		db, _ := ConnectDB()
		db.WithContext(ctx).Model(&x).Count(&count)
		return count, nil
	case entity.Item:
		var count int64
		db, _ := ConnectDB()
		db.WithContext(ctx).Model(&x).Count(&count)
		return count, nil
	default:
		return 0, fmt.Errorf("invalid type")
	}
}

//Adress validation; ADDRESS IS NOT REQUIRED FOR NOW
// type AddressValidator struct {
// }

// var addressRe = regexp.MustCompile(`\D*`)

// func (a AddressValidator) Validate(val interface{}) (bool, error) {

// 	if addressRe.MatchString(val.(string)) {
// 		return true, nil
// 	}
// 	return false, nil

// }

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")
	switch args[0] {
	case "email":
		//validator :=
		//fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return EmailValidator{}
	case "phone":
		//validator :=
		//fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return PhoneValidator{}
		// case "address":
		// 	return AddressValidator{}
	}
	return DefaultValidator{}
}

func ValidateStruct(ctx context.Context, s interface{}) []error {
	errs := []error{} // ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName) // Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		} // Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)                       // Perform validation
		valid, err := validator.Validate(ctx, v.Field(i).Interface()) // Append error to results
		if !valid && err != nil {						
			
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}
	return errs
}
func PaymentSave(ctx context.Context, charge entity.Charge) error {

	DbPool, err := ConnectDB()
	if err != nil {
		//ctx.Error(err)
		return err		
	}

	DbPool.WithContext(ctx).Create(&charge)
	return nil

}
