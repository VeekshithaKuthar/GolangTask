package main

import (
	"errors"
	"reflect"
	"strconv"
)

func main() {

}

// func IsNumber(n any) bool {
// 	switch n.(type) {
// 	case int16, int32, int64, uint, uint8, uint32, uint64, float32, float64:
// 		return true
// 	default:
// 		return false
// 	}

// }
func add(a, b any) (float64, error) {
	//a and b belongs to same type
	//sum := float64(0)

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return 0, errors.New("a and b are different types")
	}

	if !IsNumber(a) {
		return 0, errors.New("a is not number")
	}
	switch a.(type) {
	case int:
		return float64(a.(int) + b.(int)), nil
	case int32:
		return float64(a.(int32) + b.(int32)), nil
	case int64:
		return float64(a.(int64) + b.(int64)), nil

	}
	return 0, errors.New("failed")
}

// generic type
// func addG[T uint | int | uint8 | uint32 | uint64](a T, b T) {
// 	return a + b
// }

//monomorphic
//static dispatch and dynamic itables

func IsNumber(n any) bool {
	switch n.(type) {
	case int16, int32, int64, uint, uint8, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}

}
func SquareNumber(n any) (float64, error) {

	// if !IsNumber(n) {
	// 	return 0, errors.New("a is not number")
	// }
	if reflect.TypeOf(n) != reflect.TypeOf(b) {
		return 0, errors.New("a and b are different types")
	}

	if reflect.TypeOf(n).Kind() == reflect.String {
		a, err := strconv.Atoi(n.(string))
		if err != nil {
			return 0, err
		}
		return float64(a.(int) * a.(int)), nil
	}
	switch n.(type) {
	case int:
		return float64(n.(int) * n.(int)), nil
	case int32:
		return float64(n.(int32) + n.(int32)), nil
	case int64:
		return float64(n.(int64) + n.(int64)), nil
	case float32:
		return float64(n.(int64) + n.(int64)), nil
	case float64:
		return float64(n.(int64) + n.(int64)), nil
	case bool:
		n, err := strconv.ParseBool(n)
		if err != nil {
			return "", err
		}
		return n, nil
	case string:

	}

	return 0, errors.New("failed")
}
