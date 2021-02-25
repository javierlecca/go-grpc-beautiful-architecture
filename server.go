package main

import (
	"fmt"
	"log"
	"net"

	con_ "aimi-landing-back-go/config"
	"aimi-landing-back-go/controller"
)

func main() {

	config := con_.Config{}
	con, _ := config.NewConfig()

	fmt.Printf("Go gRPC server in port %v!", con.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", con.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	c := &controller.Controller{}
	g := c.NewController()

	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
