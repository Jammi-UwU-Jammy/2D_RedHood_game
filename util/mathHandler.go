package util

import (
	"github.com/co0p/tankism/lib/collision"
	_ "github.com/co0p/tankism/lib/collision"
	"math"
	"reflect"
	"time"
)

type Point struct {
	X int
	Y int
}

type Vector struct {
	X, Y float64
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

func IfCollided(box1, box2 collision.BoundingBox) bool {
	if collision.AABBCollision(box1, box2) {
		return true
	}
	return false
}

func VectorsDistance(p1x, p1y, p2x, p2y float64) float64 {
	distance := math.Sqrt(math.Pow(p1x-p2x, 2) + math.Pow(p1y-p2y, 2))
	return distance
}

func VectorLength(vx, vy float64) float64 {
	distance := math.Sqrt(math.Pow(vx, 2) + math.Pow(vy, 2))
	return distance
}

func UnitVectorFromTwoPoints(tx, ty, hx, hy float64) (float64, float64) {
	length := VectorsDistance(tx, ty, hx, hy)
	return (hx - tx) / length, (hy - ty) / length
}

func UnitVector(x, y float64) (float64, float64) {
	length := VectorLength(x, y)
	return x / length, y / length
}

func IsCDExceeded(CDinSec float64, since time.Time) bool {
	if time.Since(since).Seconds() > CDinSec {
		return true
	}
	return false
}
