package main

import "fmt"

type Integer = int

// annoymous fields
type ColourCode struct {
	int
	Integer
	string
}

func (c *ColourCode) Display() {
	println("code", c.int)
	println("Name", c.string)
	println("Integet", c.Integer)

}

type Customer struct {
	Id          Integer
	Name, Email string
	Address     Address //composition
}

type Address struct {
	Line1, pincode string
}

func main() {
	// cd1 := ColourCode(100, 200, "hello")

}

func (a *Address) Display() {
	fmt.Println(a.Line1)
}

func New(name,email,line,pincode string)*Customer{
     return &Cu
}