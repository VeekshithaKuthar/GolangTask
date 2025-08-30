package main

func main() {
	char := 'A'
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is aa vowel")
	default:
		println("its consonent")
	}

	//empty switch
	num := 40
	switch {
	case num >= 0 && num < 50:
		println("the number is between 0 to 50")
	case num >= 50 && num < 100:
		println("the number is between 50 to 100")

	}

	//fallthrough
	num = 32
	switch {
	case num%8 == 0:
		println(num, "is divisble by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisble by 4")

	}

	//false negative while using fallthrough
	//when we remove break in any other language we have tu use fallthrough

}
