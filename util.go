package util

import (
	"reflect"
)

func IsEmpty(i interface{}) bool {
	return !reflect.ValueOf(i).IsValid()
}
