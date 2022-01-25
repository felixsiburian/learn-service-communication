package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	protoA "learn-service-communication/proto-repo/protoA"
	"log"
	"time"
)

type Client struct {
	Host          string
	Port          string
	Connection    *grpc.ClientConn
	ClientService protoA.MessageServiceClient
}

var (
	host = flag.String("host", "localhost", "host")
	port = flag.String("port", "8000", "port")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("ServiceA as Client!")

	client := Client{
		Host: *host,
		Port: *port,
	}
	client.Connect()

	// send message here
	r, e := client.GetMessage("Mr. Smith")
	if e != nil {
		log.Println("Error GetMessage, ", e.Error())
	} else {
		log.Println("Message receive from ServiceB: ", r.Message)
	}

}

func (c *Client) Connect() {
	serverAddress := fmt.Sprintf("%s:%s", c.Host, c.Port)
	log.Println("Connecting to ServiceB in ", serverAddress)

	//	dialing to serviceB
	var err error
	c.Connection, err = grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Println("can't connect :  ", err)
		log.Println("try to reconnectiong after 5s")
		time.Sleep(5 * time.Second)
		defer c.Connect()
		return
	}

	c.ClientService = protoA.NewMessageServiceClient(c.Connection)
}

func (c *Client) GetMessage(name string) (*protoA.Response, error) {
	//	 create context with timeout
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10*time.Second))

	//	request
	response, err := c.ClientService.Get(ctx, &protoA.Request{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
