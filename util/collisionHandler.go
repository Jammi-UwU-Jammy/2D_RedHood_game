package util

import (
	"github.com/co0p/tankism/lib/collision"
	_ "github.com/co0p/tankism/lib/collision"
	"reflect"
)

type AABB struct {
	MinX, MinY, MaxX, MaxY float64
}

func ifObjectContainsField(obj interface{}, fieldName string) bool {
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

func IfBoxesCollided(box1, box2 AABB) bool {
	if box1.MaxX < box2.MinX || box1.MinX > box2.MaxX {
		return false
	}

	if box1.MaxY < box2.MinY || box1.MinY > box2.MaxY {
		return false
	}

	return true
}

func IfCollided(box1, box2 collision.BoundingBox) bool {
	if collision.AABBCollision(box1, box2) {
		return true
	}
	return false
}
