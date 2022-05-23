package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"pgevangelidis/PortDomainService/services/client/api/models"
	"pgevangelidis/PortDomainService/services/client/api/parser"
	"pgevangelidis/PortDomainService/services/pds"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	client pds.PortServiceClient
}

var (
	host       = flag.String("host", "localhost", "The client host")
	port       = flag.String("port", "8080", "The client port")
	sampleFile = flag.String("file", "", "The sample json test file, contains ~1600 records")
)

func (s Server) Start() {
	cl := NewPortServiceClient(fmt.Sprintf("%s:%s", *host, *port))
	s.client = cl
	router := mux.NewRouter()

	router.HandleFunc("/api/ports/testfile", s.AddPorts).Methods(http.MethodPost)
}

func (s Server) AddPorts(w http.ResponseWriter, r *http.Request) {
	if *sampleFile == "" {

	}
	stream := make(chan (models.Port))
	done := make(chan (bool), 1)

	var err error
	var counter = 0
	go func() {
		for p := range stream {
			var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			// waste of resources here while marshalling the struct again to bytes
			// if the gRPC server accepts a define struct then we can omit the Marshal
			// otherwise there must be a way to simple return bytes from the stream channel
			payload, _ := json.Marshal(p)
			_, err = s.client.InsertOrUpdate(ctx, &pds.Port{ID: p.ID, Content: payload})
			if err != nil {
				close(stream)
				cancel()
				break
			}
			cancel()
			counter++
		}
		done <- true
	}()

	// var sample = "/Users/pavlos/go/src/pgevangelidis/PortDomainService/ports.json"
	err = parser.StreamJson(*sampleFile, stream)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/text")
		w.Write([]byte(fmt.Sprintf("encountered an error while storing the %d record", counter)))
	}

	close(stream)
	<-done
	close(done)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d records of ports have been successfully inserted or updated", counter)))
}
