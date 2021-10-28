package http

import (
	"github.com/densus/pos_service/middleware"
	"github.com/densus/pos_service/pos/delivery/controller"
	"github.com/densus/pos_service/pos/service"
	"github.com/gin-gonic/gin"
)

type posHandler struct {
	authController controller.AuthController
	userController controller.UserController
	productController controller.ProductController
	outProductController controller.OutletProductController
	jwtService service.JWTService
}

func NewPosHandler(routes *gin.Engine, authCon controller.AuthController, userCon controller.UserController, productCon controller.ProductController, productController controller.OutletProductController, jwtServ service.JWTService)  {
	handler:= &posHandler{
		authController:            authCon,
		userController: 			userCon,
		productController:         productCon,
		outProductController: productController,
		jwtService:                jwtServ,
	}

	//handling request for authentication service
	authRoutes := routes.Group("api/auth")
	{
		authRoutes.POST("/login",  handler.authController.Login)
		authRoutes.POST("/register", handler.authController.Register)
	}

	//handling request for user service
	userRoutes := routes.Group("api/users", middleware.AuthorizeJWT(handler.jwtService))
	{
		userRoutes.GET("/", handler.userController.AllUser)
		userRoutes.GET("/profile", handler.userController.User)
		userRoutes.PUT("/update", handler.userController.Update)
		userRoutes.DELETE("/:id", handler.userController.Delete)
	}

	//handling request for product service
	productRoutes := routes.Group("api/products", middleware.AuthorizeJWT(handler.jwtService))
	{
		productRoutes.GET("/", handler.productController.All)
		productRoutes.POST("/create", handler.productController.Insert)
		productRoutes.PUT("/update/:id", handler.productController.Update)
		productRoutes.GET("/:id", handler.productController.FindByID)
		productRoutes.DELETE("/:id", handler.productController.Delete)
		productRoutes.POST("/:product_id/outlet-id/:outlet_id", handler.outProductController.SetPrice)
	}

}
