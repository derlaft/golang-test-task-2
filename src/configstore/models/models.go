package models

import (
	"fmt"
	"reflect"
)

type Model struct {
	Data string `gorm:"primary_key"`
}

// Database abstraction layer
type ConfigStore interface {
	Get(Type, ID string, result interface{}) error
	Set(Type, ID string, model interface{}) error
}

func NewConfig(Type string) (interface{}, error) {

	sample, found := Models[Type]
	if !found {
		return nil, fmt.Errorf("config model not present")
	}

	return NewStruct(sample), nil
}

func NewStruct(sample interface{}) interface{} {
	confType := reflect.TypeOf(sample)
	return reflect.New(confType).Interface()
}
