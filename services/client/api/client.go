package main

import (
	"log"
	"pgevangelidis/PortDomainService/services/pds"

	"google.golang.org/grpc"
)

func NewPortServiceClient(address string) pds.PortServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	return pds.NewPortServiceClient(conn)
}
