// package main

// // import (
// // 	"errors"
// // 	"fmt"
// // )

// func main() {
// 	// 	// var m1 map[string]string
// 	// 	// if m1 == nil {
// 	// 	// 	println("nil map")
// 	// 	// 	m1 = make(map[string]string)

// 	// 	// 	m1["5750017"] = "Mnagalore-1"
// 	// 	// 	m1["575002"] = "Mangalore-2"
// 	// 	// 	m1["575001"] = "Mangalore-2"
// 	// 	// 	m1["575001"] = "Mangalore-2"

// 	// 	// 	fmt.Println(m1)

// 	// 	// 	v := m1["5750017"]
// 	// 	// 	fmt.Println(v)

// 	// 	// 	for k, v := range m1 {
// 	// 	// 		println("key:", k, "value", v)
// 	// 	// 	}

// 	// 	// 	v1, ok := m1["561222"]
// 	// 	// 	if ok {
// 	// 	// 		fmt.Println(v1)
// 	// 	// 	} else {
// 	// 	// 		fmt.Println("key doesnt exist")
// 	// 	// 	}
// 	// 	// }

// 	// 	// err := Delete(m1, "5750017")
// 	// 	// if err != nil {

// 	// 	// }

// 	// 	//create 4 function add mul sub  div
// 	// 	//create function greet
// 	// 	//create function sqaure

// 	// 	//create a map with map[string]any

// 	// 	//keys:=add,sub,mul,div,greet,sq
// 	// 	//assign respectiove function to the map as value
// 	// 	//excute them in a range loop

// 	// 	var m1 map[string]any
// 	// 	m1["add"] = "add"
// 	// 	m1["sub"] = "sub"
// 	// 	m1["mul"] = "mul"
// 	// 	m1["div"] = "div"
// 	// 	m1["greet"] = "greet"

// 	// 	//a, b, n := 2, 3, 3

// 	// var any1 any
// 	// if any1 == nil {
// 	// 	fmt.Println(any1)
// 	// }

// 	// var any2 interface{}
// 	// fmt.Println(any2)

// 	// var f1 func()
// 	// f1 = greet/// Contains the pointer of Greet
// 	// f1()

// 	slice1 := make([]func(), 5)

// 	for i := 0; i < len(slice1); i++ {
// 		slice1[i] = func() {
// 			println(i)
// 		}
// 	}

// 	for _, fn := range slice1 {
// 		fn()
// 	}

// }

// // func sub() int {
// // 	return a - b
// // }

// func greet() {
// 	println("Hello World")
// }

// // func Delete(m map[string]string, key string) error {
// // 	if m == nil {
// // 		return errors.New("input map is nil")
// // 	}

// // 	if _, ok := m[key]; !ok {
// // 		return errors.New("input map is nil")
// // 	}

// // 	delete(m, key)
// // 	return nil
// // }

// // //create 4 function add mul sub  div
// // //create function greet
// // //create function sqaure

// // //create a map with map[string]any

// // //keys:=add,sub,mul,div,greet,sq
// // //assign respectiove function to the map as value
// // //excute them in a range loop

// // func add(a, b float64) any { //func(int,int)int
// // 	return a + b
// // }

// // func sub(a, b float64) any {
// // 	return a - b
// // }

// // func mul(a, b float64) any {
// // 	return a * b
// // }

// // func div(a, b float64) any {
// // 	return a / b
// // }

// // func greet(word string) {
// // 	fmt.Println("hello")
// // }

// // func square(n int) int {
// // 	return n * n
// // }
package main

import "fmt"

func main() {
	m1 := make(map[string]any)
	m1["add"] = add
	m1["sub"] = sub
	m1["mul"] = mul
	m1["div"] = div
	m1["greet"] = greet
	m1["sq"] = sq

	a, b, n := 10, 20, 5

	for k, v := range m1 {
		switch vf := v.(type) {
		case func(int, int) int:
			switch k {
			case "add":
				println("Add", vf(a, b))
			case "sub":
				println("Sub", vf(a, b))
			case "mul":
				println("Mul", vf(a, b))
			case "div":
				println("Div", vf(a, b))
			}
		case func():
			println("called greet")
			//v.(func())()
			vf()
		case func(int) int:
			println("Sq")
			println("Sq:", vf(n))
		default:
			println("not found")
		}

	}

	var m2 map[string]string = map[string]string{"name": "Jiten", "address": "Trv"}

	fmt.Println(m2)

	m3 := make(map[string]func(int, int) int)

	arr1 := [3]int{10, 20, 30}
	arr2 := [3]int{40, 50, 60}
	if arr1 == arr2 {

	}

	// slice1 := []int{10, 20, 30}
	// slice2 := []int{40, 50, 60}
	// if slice1 == slice2 {

	// }

	m3["add"] = add
	m3["sub"] = sub
	m3["mul"] = mul
	m3["div"] = div

	var fn func(int, int) int
	fn = add
	r := fn(100, 200)
	println(r)

	m4 := make(map[[3]int]string)
	m4[arr1] = "Some array with 3 elements"
	m4[arr2] = "another array"
}

func add(a, b int) int { // func(int,int)int
	return a + b
}

func sub(a, b int) int { // func(int,int)int
	return a - b
}

func mul(a, b int) int { // func(int,int)int
	return a * b
}

func div(a, b int) int { // func(int,int)int
	return a / b
}

func greet() { // func()
	println("Hello omneNEXT")
}

func sq(num int) int { // func(int)int
	return num * num
}
