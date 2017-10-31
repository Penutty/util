package util

import (
	"errors"
	"reflect"
)

func IsEmpty(i interface{}) bool {
	return !reflect.ValueOf(i).IsValid()
}

var (
	ErrorTypeNotImplemented = errors.New("Type not implemented by switch")
	ErrorDifferentTypes     = errors.New("Parameters are not equal in type")
)

func DeepEqual(i, j interface{}) (bool, error) {

	switch v := i.(type) {
	case []string:
		w, ok := j.([]string)
		if ok {
			return dEStringSlice(v, w), nil
		}
		return false, ErrorDifferentTypes

	default:
		return false, ErrorTypeNotImplemented
	}
}

func dEStringSlice(s, r []string) bool {
	if len(s) != len(r) {
		return false
	}

	for i, v := range r {
		if s[i] != v {
			return false
		}
	}
	return true
}

func IsTypeEqual(i, j interface{}) bool {
	return reflect.TypeOf(i) == reflect.TypeOf(j)
}
