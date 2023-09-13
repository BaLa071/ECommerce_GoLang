package main

import (
	"context"
	"e_commerce_site/e_commerce_DAL/services"
	controllers "e_commerce_site/e_commerce_controllers/order_controllers"
	"e_commerce_site/ecommerce_config/config"
	"e_commerce_site/ecommerce_config/constants"
	"fmt"
	"net"

	pro "e_commerce_site/ecommerce_proto/order_proto"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	OrderCollection := config.GetCollection(client, "DataBase", "Order")
	InventoryCollection := config.GetCollection(client, "inventory_SKU", "items")
	controllers.OrderService = services.InitOrderService(client, OrderCollection, InventoryCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Order_Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterOrderServiceServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Order_Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
