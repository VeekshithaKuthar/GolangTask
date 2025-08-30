package main

import "runtime"

func main() {

	func() {
		println("hello")
	}()
	func() {
		println("hello1")
	}()
	println("end of main")

	runtime.Goexit()
}

// 2mb
// read the file
// split into the words
// find the words with maximum number and find the maximum nube
