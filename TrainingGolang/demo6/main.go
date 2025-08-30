package main

import "fmt"

func main() {
	string1 := "hello world"
	string2 := ""
	for _, v := range string1 {
		string2 += string(v - 32)
	}

	fmt.Println(string2)
}

//array
// var arr1 [5]int //zero value
// arr :=[5]int{1,2,3,4,5}

// func lowerToUpper() {
// 	s := "hello world"
// 	var str1 string
// 	for _, v := range s {
// 		if v >= 'a' && v <= 'z' {
// 			str1 += string(v - 32)
// 		} else {
// 			str1 += string(v)
// 		}
// 	}
// 	println(str1)
// }

//array
