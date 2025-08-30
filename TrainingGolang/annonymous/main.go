package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var e1, e2 Empty

	fmt.Printf("Address of e1:%p size of e1:%d\n", &e1, unsafe.Sizeof(e1))
	fmt.Printf("Address of e2:%p\n", &e2)

	// e1.Greet()
	// e2.Greet()

	var t1 T1
	fmt.Printf("Address of t1:%p size of t1:%d\n", &t1, unsafe.Sizeof(t1))
	var t2 T2
	fmt.Printf("Address of t2:%p size of t2:%d\n", &t2, unsafe.Sizeof(t2))

}

type Empty struct{} // missed call

func (e Empty) Greet() {
	println("Hey I am empty")
}

type T1 struct {
	ok1 bool // 1
	N1  int  // 8
	ok2 bool // 1
	N2  int  // 8
}

// Decending order
type T2 struct {
	N1  int  // 8
	N2  int  // 8
	ok1 bool // 1
	ok2 bool // 1
	N3  int32
}

//arena allocator
