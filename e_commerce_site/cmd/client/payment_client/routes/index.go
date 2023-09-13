package routes

import (
	 "e_commerce_site/cmd/client/payment_client/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	r.POST("/createpayment", controller.HandleCreatePayment)
}
