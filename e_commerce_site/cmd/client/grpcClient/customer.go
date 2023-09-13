package grpcclient

import (
	"log"
	"sync"

	"e_commerce_site/ecommerce_config/constants"
	pb "e_commerce_site/ecommerce_proto/customer_proto"

	"google.golang.org/grpc"
)

var once sync.Once

type GrpcCustomerServiceClient pb.CustomerServiceClient

var (
	instance GrpcCustomerServiceClient
)

func GetGrpcCustomerServiceClient() (GrpcCustomerServiceClient,*grpc.ClientConn) {
	var conn *grpc.ClientConn
	once.Do(func() { // <-- atomic, does not allow repeating
		conn, err := grpc.Dial("localhost"+constants.Customer_Port, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		//defer conn.Close()

		instance = pb.NewCustomerServiceClient(conn)
	})

	return instance,conn
}
