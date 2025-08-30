package main

import (
	"errors"
	"fmt"

	"strconv"
)

func main() {
	var a any
	a = "5.5"

	if sqr, err := square(a); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Square: %v\n", sqr)
	}
}

func IsNumber(n any) bool {
	//type switch
	switch n.(type) {
	case uint, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, string, bool:
		return true
	}
	return false
}

func square(a any) (any, error) {
	if !IsNumber(a) {
		return 0, errors.New("input argument is not a number")
	}

	switch v := a.(type) {
	case int:
		return (a.(int) * a.(int)), nil
	case uint:
		return a.(uint) * a.(uint), nil
	case float32:
		return v * v, nil
	case float64:
		return v * v, nil
	case string:
		b, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid numeric string: %w", err)
		}
		return b * b, nil
	case bool:
		n := 0
		if v {
			n = 1
		}
		return n * n, nil
	default:
		return 0, errors.New("unsupported type")
	}
}
