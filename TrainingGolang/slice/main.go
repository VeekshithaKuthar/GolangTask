// // package main

// // import "fmt"

// // func main() {
// // 	// var slice []int
// // 	// if slice == nil {
// // 	// 	fmt.Println("slice 1 is anil")
// // 	// 	slice = make([]int, 10, 20)
// // 	// }
// // 	// fmt.Println(slice, "len:", len(slice), "capacity", cap(slice), "sliceptr", &slice[0], "address", &slice[0], &slice)
// // 	// // for i := range slice {
// // 	// // 	slice[i]
// // 	// // }

// // 	// fmt.Printf("slice:%v\n len:%v\n cap:%d\n ptr:%p\naddress:%p\n", slice, len(slice), cap(slice), &slice[0], &slice)
// // 	// println("sum:", sumOf())
// // 	// println("sum:", sumOf(10,20))

// // 	// slice1 := []int{1, 2, 3, 4, 5}

// // 	// slice2 := slice1

// // 	// slice2[1] = 9999

// // 	// fmt.Printf("slice:%v\n len:%v\n cap:%d\n ptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)

// // 	// fmt.Printf("slice:%v\n len:%v\n cap:%d\n ptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)

// // 	// var sl []int
// // 	// s2:=make([]int,0)
// // 	// s3:=[]int{}

// // 	slice :=[]int{1,2,3,4,5,6,7,8,9,10}
// // 	 slice1:=slice
// // 	 slice2:=

// // }

// // // variadic params... can only be used in function or method
// // // variadic params...  must be last params
// // func sumOf(nums ...int) int {
// // 	sum := 0
// // 	for _, v := range nums {
// // 		sum = sum + v
// // 	}
// // 	return sum
// // }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	// 	var slice1 []int

// 	// 	if slice1 == nil {
// 	// 		println("slice1 is a nil slice")
// 	// 		slice1 = make([]int, 5, 5)
// 	// 	}

// 	// 	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// 	for i := range slice1 {
// 	// 		slice1[i] = rand.IntN(100)
// 	// 	}
// 	// 	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// 	slice1 = append(slice1, 1111)
// 	// 	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// 	slice1 = append(slice1, 2222)
// 	// 	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// 	slice1 = append(slice1, 3333, 4444, 5555)
// 	// 	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)

// 	// 	slice1 = append(slice1, 6666)
// 	// 	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// 	println("Sum:", SumOfV())
// 	// 	println("Sum:", SumOfV(10, 20))
// 	// 	println("Sum:", SumOfV(12, 32, 34, 34, 4, 57, 57, 8, 7))
// 	// 	println("Sum:", SumOfV(slice1...))
// 	// 	//fmt.Println(true, 0, "helo", 12.213, "world", 1, 2, 3, false, true, false)

// 	// 	arr := [3]int{10, 20, 30}
// 	// 	fmt.Println(SumOf(arr[:]))
// 	// 	fmt.Println(SumOf(slice1))

// 	// slice1 := []int{1, 2, 3, 4, 5}
// 	// fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// slice1 = append(slice1, 5)
// 	// fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// slice2 := make([]int, 10)
// 	// copy(slice2, slice1)
// 	// fmt.Printf("slice2:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)

// 	// slice1 := []int{1, 2, 3}
// 	// slice2 := slice1[:2] // slice2 refers to first two elements of slice1

// 	// slice2[0] = 100

// 	// slice1 := []int{1, 2, 3}
// 	// slice2 := slice1[:2] // slice2 refers to first two elements of slice1
// 	// fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	// slice2[0] = 100

// 	slice1 := []int{1, 2, 3, 4}
// 	slice2 := slice1[1:3]
// 	slice2[0] = 99
// 	fmt.Println("a:", slice1)
// 	fmt.Println("b:", slice2)

// 	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)

// 	// fmt.Println("slice1:", slice1)
// 	// fmt.Println("slice2:", slice2)

// 	//slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	//slice2 := append(slice1[:4], slice1[5:]...)
// 	//fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
// 	fmt.Printf("slice2:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)

// 	// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	// //slice3:=arr[3:9]
// 	// slice3 := append(arr[:2], arr[3:]...)
// 	// fmt.Println(slice3)

// }

// // variadic params..
// // variadia param can only be used i5 function or method
// // variadic param must be the last param
// // func SumOfV(nums ...int) int {
// // 	sum := 0

// // 	for _, v := range nums {
// // 		sum += v
// // 	}
// // 	return sum
// // }

// // func SumOf(slice []int) int {
// // 	sum := 0
// // 	for _, v := range slice {
// // 		sum += v
// // 	}
// // 	return sum
// // }

package main

import "fmt"

func main() {

	slice1 := [5]int{1, 2, 3, 4, 5} //ptr: len:5 cap:5
	slice2 := slice1[:3]
	//slice1 := make([]int, 5, 5)
	// slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 1, 2, 3, 4, 5
	// fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)

	// slice2 := slice1 // The headers are copied
	// fmt.Printf("slice2:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)

	// slice2[0] = 9999

	// slice2 = append(slice2, 20)
	// fmt.Println("slice1", slice1)
	// fmt.Println("slice2", slice2)
	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	fmt.Printf("slice2:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)
	// slice2[1] = 8888
	// fmt.Println(slice1)
	// fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	// fmt.Printf("slice2:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)
	// sq(slice1)
	// fmt.Println(slice1)
	// AddSq(slice1, 6, 7, 8, 9, 10)
	// fmt.Println(slice1)
	// slice1 = AddSqR(slice1, 6, 7, 8, 9, 10)
	// fmt.Println(slice1)
}

func sq(slice []int) {
	for i, v := range slice {
		slice[i] = v * v
	}
}

func AddSq(slice []int, nums ...int) {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * v
	}
}

func AddSqR(slice []int, nums ...int) []int {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}
