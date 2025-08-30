package main

import (
	"fmt"
	"math/rand/v2"
	"unsafe"
)

var Global int = 1000

func main() {

	// var num1 int = 100
	// var ptr1 *int = &num1

	// ptr1 = &Global
	// println(ptr1)

	// slice := make([]int, 1000000, 1000000) ///main.go:13:15: make([]int, 100, 100) does not escape not stored in heap stored in stack
	// println(slice)

	// ptr2 := new(int)
	// *ptr2 = 1000
	// var num2 uint8 = 127
	// var ptr3 *uint8=

	//void pointers
	// 	unsafe.Pointer

	// 	A pointer value of any type can be converted to a Pointer.
	// A Pointer can be converted to a pointer value of any type.
	// A uintptr can be converted to a Pointer.
	// A Pointer can be converted to a uintptr.

	// arr:=[5]int{1,2,3,4,5}
	// ptr:=&arr[0]

	slice1 := make([]int, 5, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.IntN(10)
	}

	ptr := (*[3]int)(unsafe.Pointer(&slice1))
	ptr[1] = 10
	fmt.Println(ptr)

}

//go tool nm pointer.exe |grep Global
//go build --gcflags="-m" main.go

//ldflags=linker flags

//new
//created dangling pointers can be moved to heap
//POINTER ARTHMETIC directly not used in go

func Square(num int) int {
	return num * num
}

//void pointers

//slice:=make([]int,5,10)
//fill the sice

//o ptr
//1 len -->10
//2->cap
