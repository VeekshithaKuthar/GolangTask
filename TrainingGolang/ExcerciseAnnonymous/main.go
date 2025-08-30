package main

func main() {
	execute(10, 20, func(i1, i2 int) int { return i1 + i2 }, func(i1, i2 int) int { return i1 - i2 }, func(i1, i2 int) int { return i1 * i2 }, func(i1, i2 int) int { return i1 / i2 })

}

func execute(a, b int, funcs ...any) {
	for _, v := range funcs {
		switch t := v.(type) {
		case func(int, int) int:
			return v(a, b)

		}
	}
}
