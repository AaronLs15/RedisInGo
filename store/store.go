package store

import (
	"errors"
)

var ErrEmptyKey = errors.New("key is mandatory")
var ErrStoreFull = errors.New("Store is full")

type Storer interface {
	Get(key string) (string, error)
	Set(key, value string) error
	// Keys() []string
	Delete(key string)
	Len() int
}
