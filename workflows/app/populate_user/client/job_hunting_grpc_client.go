package client

import (
	"log"

	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClient() *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(config.Koanf.String("jobhunt_service.grpc_address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
