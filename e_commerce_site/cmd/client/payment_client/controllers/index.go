package controller

import (
	grpcclient "e_commerce_site/cmd/client/grpcClient"
	controllers "e_commerce_site/e_commerce_controllers/payment_controllers"
	"e_commerce_site/ecommerce_config/constants"
	"fmt"
	"net/http"

	pb "e_commerce_site/ecommerce_proto/payment_proto"

	"github.com/gin-gonic/gin"
)

func HandleCreatePayment(c *gin.Context) {

	grpcClient, _ := grpcclient.GetGrpcPaymnentService()
	var res pb.PaymentDetails
	token := c.GetHeader("Authorization")

	result1, err1 := controllers.ExtractCustomerID(token, constants.SecretKey)
	//var request pb.PaymentDetails
	res.CustomerID = result1
	fmt.Println(res.CustomerID)
	fmt.Println(err1)

	if err := c.ShouldBindJSON(&res); err != nil {
		fmt.Println("12")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(result1)
	response, err := grpcClient.CreatePayment(c.Request.Context(), &res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": response})
}
