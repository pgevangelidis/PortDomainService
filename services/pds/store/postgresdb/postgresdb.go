package postgresdb

import (
	"database/sql"
	"fmt"
	"pgevangelidis/PortDomainService/services/pds/store"
)

type db struct {
	pool *sql.DB
}

// this can be extracted as a configuration and not be hard-coded
// const dsn = "postgres://user:password@localhost:5432/portsdb?sslmode=disable&pool_max_conns=5"

func New() store.IStore {
	return &db{}
}

func (s *db) Create(key string, value interface{}) error {
	return fmt.Errorf("not implemented")
}

func (s *db) Update(key string, value interface{}) error {
	return fmt.Errorf("not implemented")
}

func (s *db) Select(key string) (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *db) Delete(key string) error {
	return fmt.Errorf("not implemented")
}
