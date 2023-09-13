package controllers

import (
	"context"
	"e_commerce_site/e_commerce_DAL/interfaces"
	"e_commerce_site/e_commerce_DAL/models"
	pro "e_commerce_site/ecommerce_proto/payment_proto"
)

type RPCServer struct {
	pro.UnimplementedPaymentServiceServer
}

var (
	TransactionService interfaces.Ipayment
)

func (t *RPCServer) CreatePayment(ctx context.Context, req *pro.PaymentDetails) (*pro.PaymentResponse, error) {
	transactions := &models.Paymentscard{
		Cardno:          float32(req.Cardno),
		Brand:           "",
		PanLastFourNo:   "",
		ExpirationMonth: 0,
		ExpirationYear:  0,
		Cvvverified:     int32(req.Cvvverified),
		Balance:         float32(req.Amount),
	}

	newProfile, err := TransactionService.CreatePayment(float32(transactions.Cardno), int(transactions.Cvvverified), transactions.Balance,req.CustomerID)

	if err != nil {

		return nil, err
	} else {
		responsePayment := &pro.PaymentResponse{
			Status: newProfile,
		}
		return responsePayment, nil
	}

}
