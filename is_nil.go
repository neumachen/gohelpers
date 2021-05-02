package paratils

import "reflect"

// IsNil checks if the given interface is nil.
func IsNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
