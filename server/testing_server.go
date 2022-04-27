package server

import (
	"fmt"
	"gin-exercise/controller"
	"gin-exercise/middlewares"
	"gin-exercise/service"

	//ginDumb "github.com/tpkeeper/gin-dumb"

	//"middlewares"
	"github.com/gin-gonic/gin"
)

var (
	services service.UserService       = service.NewUser()
	controll controller.UserController = controller.NewUserController(services)

	catservices service.CategoryService       = service.NewCategory()
	catcontroll controller.CategoryController = controller.NewCategoryController(catservices)

	itemservices service.ItemService       = service.NewItem()
	itemcontroll controller.ItemController = controller.NewItemController(itemservices)

	loginService    service.LoginService       = &service.UserLogin{}
	jwtService      service.JWTService         = service.JWTAuthService()
	loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	imgservices service.ImageService       = service.NewImage()
	imgcontroll controller.ImageController = controller.NewImage(imgservices)

	storeservices service.StoreService       = service.NewStore()
	storecontroll controller.StoreController = controller.NewStoreController(storeservices)

	orderservices service.OrderService       = service.NewOrder()
	ordercontroll controller.OrderController = controller.NewOrderController(orderservices)

	// orderservices service.OrderService       = service.NewOrder()
	// ordercontroll controller.OrderController = controller.NewOrderController(orderservices)

	// paymentcontroll controller.PaymentController = controller.NewPaymentController(storeservices)
)

func TestingServer() *gin.Engine {
	//setUptLogOutPut()
	server := gin.New()

	server.Static("/v1/image", "./")

	server.Use(gin.Recovery(), middlewares.Logger() )

	

	//	server.POST("/users", controll.Save)

	// if hasPolicy := service.Enforcer().HasPolicy("ROOT", "StoreID", "Item", "read"); !hasPolicy {
	// 	_, er := service.Enforcer().AddPolicy("ROOT", "StoreID", "Item", "read")
	// 	if er != nil {
	// 		fmt.Println("", er)
	// 	}
	// 	// fmt.Println(" ",en)
	// }

	if hasPolicy := service.Enforcer().HasPolicy("Customer", "StoreID", "Item", "read"); !hasPolicy {
		en, er := service.Enforcer().AddPolicy("Customer", "StoreID", "Item", "read")

		if er != nil {
			fmt.Println("the error in first ", er)
		}
		fmt.Println("enforecment result 1: ", en)
	}

	if hasPolicy := service.Enforcer().HasPolicy("Owner", "StoreID", "Item", "write"); !hasPolicy {
		en, er := service.Enforcer().AddPolicy("Owner", "StoreID", "Item", "write")

		if er != nil {
			fmt.Println("the error in first ", er)
		}
		fmt.Println("enforecment result 1: ", en)
	}

	if hasPolicy := service.Enforcer().HasPolicy("Owner", "StoreID", "Item", "read"); !hasPolicy {
		en, er := service.Enforcer().AddPolicy("Owner", "StoreID", "Item", "read")

		if er != nil {
			fmt.Println("the error in first ", er)
		}
		fmt.Println("enforecment result 1: ", en)
	}

	//add new user

	server.PATCH("/user/:id", middlewares.IsLogged(), controll.Update)

	server.POST("/users", controll.Save)

	server.GET("/user", middlewares.IsLogged(), controll.Search)

	server.DELETE("/user/:userid", middlewares.IsLogged(), controll.Delete)

	server.POST("/image", imgcontroll.ImageSave)

	server.POST("/categories", middlewares.IsLogged(), catcontroll.Save)

	server.POST("/items", middlewares.IsLogged(), itemcontroll.Save)

	server.GET("/categories", catcontroll.FindAll)

	server.GET("/item", itemcontroll.Search)

	server.GET("/items", itemcontroll.FindAll)

	//login
	server.POST("/login", loginController.Login)

	//image upload
	server.POST("/user/:userid",imgcontroll.ImageSave)

	server.POST("/stripe", middlewares.IsLogged(),controll.Payment)

	server.POST("/stores", middlewares.IsLogged(),storecontroll.Save)

	server.GET("/stores", storecontroll.FindAll)

	server.GET("/store", storecontroll.Search)

	server.PATCH("/stores/:id", middlewares.IsLogged(), middlewares.Authorize("Item", "read", service.Enforcer()), storecontroll.Update)

	server.PATCH("/stores/:id/status", middlewares.IsLogged(), middlewares.Authorize("Store", "write", service.Enforcer()), storecontroll.UpdateStoreStatus)

	server.DELETE("/stores/:id", middlewares.IsLogged(), storecontroll.Delete)

	server.DELETE("/item", middlewares.IsLogged(),itemcontroll.Delete)

	server.POST("/orders",middlewares.IsLogged(), ordercontroll.Save)

	server.GET("/orders", middlewares.IsLogged(), ordercontroll.FindAll)

	server.GET("/orders/:orderid", middlewares.IsLogged(), ordercontroll.Search)

	server.PATCH("/orders/:id", middlewares.IsLogged(), ordercontroll.Update)
	
	return server
}
