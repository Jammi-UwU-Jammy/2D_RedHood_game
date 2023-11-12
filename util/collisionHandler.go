package util

import (
	"reflect"
)

func GetField(obj interface{}, fieldName string) bool {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr { // Dereference the pointer if obj is a pointer
		val = val.Elem()
	}

	field := val.FieldByName(fieldName)

	if field.IsValid() {
		//fmt.Printf("The struct has the field '%s' : %v.\n", fieldName, field)
		return true
	} else {
		//fmt.Printf("The struct does not have the field '%s'.\n", fieldName)
		return false
	}
}

func IsAColliderObject(obj interface{}) {
	if GetField(obj, "") {

	}
}
