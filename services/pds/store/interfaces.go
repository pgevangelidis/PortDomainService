package store

type (
	// This interface extracts the common crud logic from the database implementation.
	// The accepted parameters are an assumption for the time being. In later iterations and
	// improvements, we can pass structs that will match any store implementation, sql or nosql.
	IStore interface {
		Create(key string, value interface{}) error
		Update(key string, value interface{}) error
		Select(key string) (interface{}, error)
		Delete(key string) error
	}
)
