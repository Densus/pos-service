package main

import (
	"github.com/densus/pos_service/config"
	"github.com/densus/pos_service/pos/delivery/controller"
	"github.com/densus/pos_service/pos/delivery/http"
	"github.com/densus/pos_service/pos/repository"
	"github.com/densus/pos_service/pos/service"
	"github.com/gin-gonic/gin"
)

func main()  {
	db := config.SetupDBConnection()
	defer config.CloseDBConnection(db)

	r := gin.Default()

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	outletProductRepository := repository.NewOutletProductRepository(db)

	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository)
	productService := service.NewProductService(productRepository)
	authService := service.NewAuthService(userRepository)
	outletProductService := service.NewOutletProductService(outletProductRepository)

	authController := controller.NewAuthController(authService, jwtService)
	userController := controller.NewUserController(userService, jwtService)
	productController:= controller.NewArticleController(productService, jwtService)
	outletProductController := controller.NewOutletProductController(outletProductService, jwtService)

	http.NewPosHandler(r, authController, userController, productController, outletProductController, jwtService)

	r.Run()
}
