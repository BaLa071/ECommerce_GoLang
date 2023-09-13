package grpcclient

import (
	"log"
	"sync"

	"e_commerce_site/ecommerce_config/constants"
	pb "e_commerce_site/ecommerce_proto/payment_proto"

	"google.golang.org/grpc"
)

var paymentOnce sync.Once

type GrpcPaymentServiceClient pb.PaymentServiceClient

var (
	paymentInstance GrpcPaymentServiceClient
)

func GetGrpcPaymnentService() (GrpcPaymentServiceClient,*grpc.ClientConn) {
	var conn *grpc.ClientConn
	paymentOnce.Do(func() { // <-- atomic, does not allow repeating
		conn, err := grpc.Dial("localhost"+constants.Payment_Port, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		//defer conn.Close()

		paymentInstance = pb.NewPaymentServiceClient(conn)
	})

	return paymentInstance,conn
}
