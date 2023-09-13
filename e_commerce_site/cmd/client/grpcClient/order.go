package grpcclient

import (
	"log"
	"sync"

	"e_commerce_site/ecommerce_config/constants"
	pb "e_commerce_site/ecommerce_proto/order_proto"

	"google.golang.org/grpc"
)

var orderOnce sync.Once

type GrpcOrderServiceClient pb.OrderServiceClient

var (
	orderInstance GrpcOrderServiceClient
)

func GetGrpcOrderService() (GrpcOrderServiceClient,*grpc.ClientConn) {
	var conn *grpc.ClientConn
	orderOnce.Do(func() { // <-- atomic, does not allow repeating
		conn, err := grpc.Dial("localhost"+constants.Order_Port, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		//defer conn.Close()

		orderInstance = pb.NewOrderServiceClient(conn)
	})

	return orderInstance,conn
}
