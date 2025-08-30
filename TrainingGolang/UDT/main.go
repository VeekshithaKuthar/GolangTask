package main

type Person struct {
}

func main() {
	//pointer objex=ct
	p4 := new(Person)
	p4.Id = 100
	p4.Name = "veekshhitha"

	p5 := &Person{101, "veekshitha", "cfrfr", "86444"}

}
