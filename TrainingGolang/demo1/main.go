package main

func main() {
	a, b := 10, 20
	//atomic operation
	c := a + b/2*(a+5) + (b-2)*10 + 20
	println(c)
}
