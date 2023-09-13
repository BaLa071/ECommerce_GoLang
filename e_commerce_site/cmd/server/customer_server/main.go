package main

import (
	"context"
	"e_commerce_site/e_commerce_DAL/services"
	controllers "e_commerce_site/e_commerce_controllers/customer_controllers"
	"e_commerce_site/ecommerce_config/config"
	"e_commerce_site/ecommerce_config/constants"

	"fmt"
	"net"

	pro "e_commerce_site/ecommerce_proto/customer_proto"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func initDatabase(client *mongo.Client) {
	profileCollection := config.GetCollection(client, "Ecommerce", "CustomerProfile")
	tokenCollection := config.GetCollection(client, "Ecommerce", "Tokens")
	controllers.CustomerService = services.InitCustomerService(profileCollection, tokenCollection, context.Background())

}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)

	lis, err := net.Listen("tcp", constants.Customer_Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Customer_Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
