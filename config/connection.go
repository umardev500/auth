package config

import (
	"auth/pb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct{}

func NewConn() *Connection {
	return &Connection{}
}

func (c *Connection) getConn(target string) (conn *grpc.ClientConn) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (c *Connection) UserConn() (client pb.UserServiceClient) {
	port := os.Getenv("USER_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewUserServiceClient(conn)

	return
}

func (c *Connection) CustomerConn() (client pb.CustomerServiceClient) {
	port := os.Getenv("CUSTOMER_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewCustomerServiceClient(conn)

	return
}
