package config

import (
	"log"
	"os"
)

const (
	HOST = "HOST"
	PORT = "PORT"
)

type (
	conf struct {
		host string
		port string
	}
)

var Variables *conf

func init() {
	var found bool
	c := conf{}
	c.host, found = os.LookupEnv(HOST)
	if !found {
		log.Fatal("the server's env variable HOST is missing")
	}
	c.port, found = os.LookupEnv(PORT)
	if !found {
		log.Fatal("the server's env variable PORT is missing")
	}
	// instatiate the Variables
	Variables = &c
}

func (c *conf) Host() string {
	return c.host
}

func (c *conf) Port() string {
	return c.port
}
