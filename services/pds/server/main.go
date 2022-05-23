package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"pgevangelidis/PortDomainService/services/pds"
	"pgevangelidis/PortDomainService/services/pds/config"
	"pgevangelidis/PortDomainService/services/pds/store"
	"pgevangelidis/PortDomainService/services/pds/store/mockdb"
	"syscall"

	"google.golang.org/grpc"
)

type (
	ps struct {
		memory store.IStore
	}
)

func NewPDService() pds.PortServiceServer {
	return &ps{memory: mockdb.New()}
}

func (p *ps) InsertOrUpdate(ctx context.Context, req *pds.Port) (*pds.Empty, error) {
	return nil, p.memory.Create(req.ID, req.Content)
}

func main() {
	list, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Variables.Host(), config.Variables.Port()))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	srv := NewPDService()
	pds.RegisterPortServiceServer(server, srv)

	log.Printf("server listening at %v", list.Addr())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		log.Println("service is shutting down after signal %v", <-stop)
		server.GracefulStop()
	}()
	err = server.Serve(list)
	if err != nil {
		log.Fatal(err)
	}
}
