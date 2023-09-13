package main

import (
	"context"
	"e_commerce_site/e_commerce_DAL/services"
	controllers "e_commerce_site/e_commerce_controllers/payment_controllers"
	"e_commerce_site/ecommerce_config/config"
	"e_commerce_site/ecommerce_config/constants"
	"fmt"
	"net"

	pro "e_commerce_site/ecommerce_proto/payment_proto"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDatabase(client *mongo.Client) {

	CustomerCollection := config.GetCollection(client, "pay", "payments")
	transactionCollection := config.GetCollection(client, "pay", "transactions")
	orderCollection := config.GetCollection(client, "pay", "order")

	controllers.TransactionService = services.NewTransactionServiceInit(client, CustomerCollection, transactionCollection, orderCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Payment_Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()

	pro.RegisterPaymentServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Payment_Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
