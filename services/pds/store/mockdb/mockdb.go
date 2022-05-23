package mockdb

import (
	"fmt"
	"log"
	"pgevangelidis/PortDomainService/services/pds/store"
	"sync"
)

type db struct {
	memory sync.Map
}

func New() store.IStore {
	return &db{
		memory: sync.Map{},
	}
}

func (s *db) Create(key string, value interface{}) error {
	return s.write(key, value)
}

func (s *db) Update(key string, value interface{}) error {
	return s.write(key, value)
}

func (s *db) Select(key string) (interface{}, error) {
	return s.read(key)
}

func (s *db) Delete(key string) error {
	return s.delete(key)
}

// unexported function write that locks the memory store in order to
// execute a write
func (s *db) write(key string, value interface{}) error {
	if s == nil {
		return fmt.Errorf("there is a problem with the memory store, the store is null")
	}
	_, loaded := s.memory.LoadOrStore(key, value)
	if loaded {
		log.Printf("the record already exists (key: %s)", key)
	} else {
		s.memory.Store(key, value)
	}
	return nil
}

// unexported function read that locks the memory store in order to read the object
// as an empty interface and return either the object or an error
func (s *db) read(key string) (interface{}, error) {
	if s == nil {
		return nil, fmt.Errorf("there is a problem with the memory store, the store is null")
	}
	val, ok := s.memory.Load(key)
	if !ok {
		return nil, fmt.Errorf("not found (key: %s)", key)
	}
	return val, nil
}

func (s *db) delete(key string) error {
	if s == nil {
		return fmt.Errorf("there is a problem with the memory store, the store is null")
	}
	_, existed := s.memory.LoadAndDelete(key)
	if !existed {
		return fmt.Errorf("not found (key: %s)", key)
	}
	return nil
}
