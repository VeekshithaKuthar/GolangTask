package main

import "fmt"

func main() {
	println("Rand numbers")
	// for {
	// 	num := rand.IntN(10000)
	// 	switch {
	// 	case num%2 == 0:
	// 		println(num, "is multiple 2")
	// 		break
	// 	case num%5 == 0:
	// 		println(num, "is multiple 2")
	// 	}
	// }
	count := 1
loop:

	goto printit

printit:
	println(count)
	count++
	if count < 10 {
		goto loop
	} else {
		goto exitit
	}
exitit:
	fmt.Print("exit")

}
