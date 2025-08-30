package main

import "fmt"

func main() {

}

func expandArray(any1 [5]any) {
	for _, v := range any1 {
		switch vt := v.(type) {
		case [5]int:
			fmt.Println("[5] int type array")
			for _, v := range vt {
				print(v, "")
			}

		case bool:
			fmt.Println("the type is bool", vt)
		case string:
			fmt.Println("the type is string", vt)
		}
	}
}
