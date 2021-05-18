package paratils

import (
	"fmt"
	"reflect"
)

type ArrayLengther interface {
	GetLength() int
}

// ArrayIsEmpty ...
func ArrayIsEmpty(arr ArrayLengther) bool {
	return arr == nil || arr.GetLength() > 0
}

// IsNil checks if the given interface is nil.
func IsNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}

// AreNil checks if the given interface is nil.
func AreNil(v ...interface{}) bool {
	for i := range v {
		if !IsNil(v[i]) {
			return false
		}
	}
	return true
}

type StringGetter interface {
	GetString() string
}

func EqualStrings(a, b StringGetter) (bool, string) {
	equal := a.GetString() == b.GetString()
	return equal, fmt.Sprintf("%s, %s, equal: %v\n", a.GetString(), b.GetString(), equal)
}
