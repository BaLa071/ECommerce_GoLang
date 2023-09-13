package routes

import (
	"github.com/gin-gonic/gin"
 "e_commerce_site/cmd/client/customer_client/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CustomerRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	user := v1.Group("/users")
	{
	user.POST("/signup", customer_controller.HandlerSignUp)
	user.POST("/signin", customer_controller.HandlerSignIn)
	user.POST("/delete", customer_controller.HandlerDeleteCustomer)
	user.POST("/update", customer_controller.HandlerUpdateCustomer)
	user.GET("/getbyid", customer_controller.HandlerGetById)
	user.POST("/reset", customer_controller.HandlerResetPassword)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
