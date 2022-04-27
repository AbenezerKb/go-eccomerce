package main

import (
	"fmt"
	"gin-exercise/controller"
	"gin-exercise/middlewares"
	// "gin-exercise/Errors"
	"gin-exercise/service"
	"path"
	"time"

	"github.com/gin-contrib/timeout"

	//ginDumb "github.com/tpkeeper/gin-dumb"

	//"middlewares"
	"github.com/gin-gonic/gin"
)

var (
	services service.UserService       = service.NewUser()
	controll controller.UserController = controller.NewUserController(services)

	loginService    service.LoginService       = &service.UserLogin{}
	jwtService      service.JWTService         = service.JWTAuthService()
	loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	imgservices service.ImageService       = service.NewImage()
	imgcontroll controller.ImageController = controller.NewImage(imgservices)

	storeservices service.StoreService       = service.NewStore()
	storecontroll controller.StoreController = controller.NewStoreController(storeservices)

	itemservices service.ItemService       = service.NewItem()
	itemcontroll controller.ItemController = controller.NewItemController(itemservices)
	// orderservices service.OrderService       = service.NewOrder()
	// ordercontroll controller.OrderController = controller.NewOrderController(orderservices)
)

func main() {
	//setUptLogOutPut()
	server := gin.New()

	server.Static("/v1/image", "./")

	server.Use(gin.Recovery(), middlewares.Logger()) //, ginDumb.Dumb() , middlewares.AuthorizeJWT()

	if hasPolicy := service.Enforcer().HasPolicy("Customer", "Item", "read"); !hasPolicy {
		en, er := service.Enforcer().AddPolicy("Customer", "Item", "write")

		if er != nil {
			// Errors.Unable_to_save(er.Error())
			return
		}
		fmt.Println("enforecment result 1: ", en)
	}

	// if hasPolicy := service.Enforcer().HasPolicy("Owner", "StoreID", "Item", "write"); !hasPolicy {
	// 	en, er := service.Enforcer().AddPolicy("Owner", "StoreID", "Item", "write")

	// 	if er != nil {
	// 		fmt.Println("the error in first ", er)
	// 	}
	// 	fmt.Println("enforecment result 1: ", en)
	// }

	// if hasPolicy := service.Enforcer().HasPolicy("Owner", "StoreID", "Item", "read"); !hasPolicy {
	// 	en, er := service.Enforcer().AddPolicy("Owner", "StoreID", "Item", "read")

	// 	if er != nil {
	// 		fmt.Println("the error in first ", er)
	// 	}
	// 	fmt.Println("enforecment result 1: ", en)
	// }

	//server.ServeHTTP()
	//server.R
	//get users list
	///:size/:page

	//server.GET("/users", controll.FindAll)

	//Get all users
	server.GET("/users", timeout.New(
		timeout.WithTimeout(100*time.Second),
		timeout.WithHandler(controll.FindAll),
	))

	//search by email
	server.GET("/user", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(controll.Search),
	))

	server.PATCH("/user/:id", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(controll.Update),
	))

	//add new user
	server.POST("/users", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(controll.Save),
	))

	//login
	server.POST("/login", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(loginController.Login),
	))

	//image upload

	server.POST("/user/:userid", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(imgcontroll.ImageSave),
	))
	server.POST("/image/:userid", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(imgcontroll.ImageSave),
	))
	server.DELETE("/user/:userid", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(controll.Delete),
	))

	server.POST("/stripe", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(controll.Payment),
	))

	server.POST("/stores", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(storecontroll.Save),
	))

	server.GET("/stores", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(storecontroll.FindAll),
	))

	server.GET("/store", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(storecontroll.Search),
	))

	server.PATCH("/stores/:id", middlewares.Authorize("Item", "read", service.Enforcer()), timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(storecontroll.Update),
	))

	server.DELETE("/stores/:id", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(storecontroll.Delete),
	))

	server.POST("/stores/role", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(storecontroll.CreateStoreRole),
	))
	server.GET("/item", timeout.New(
		timeout.WithTimeout(1000*time.Second),
		timeout.WithHandler(itemcontroll.Search),
	))
	//access image
	//server.GET("/image/:userid", imgcontroll.Display_image)
	server.Static("/image", path.Join("v1", "image"))
	server.Run(":8080")
	//server.
}
