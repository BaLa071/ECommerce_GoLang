package routes

import (
	"e_commerce_site/cmd/client/order_client/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {

	r.POST("/createorder", controller.HandlerCreateOrder)
	r.POST("/updateorder", controller.HandlerUpdateOrder)
	r.POST("/Addorder", controller.HandlerAddOrder)
	r.GET("/Deleteorder", controller.HandlerDeleteOrder)
	r.GET("/Getorderbyid", controller.HandlerGetOrderById)
}
