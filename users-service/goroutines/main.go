// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	wg := new(sync.WaitGroup)
// 	chan1 := make(chan int)
// 	chan2 := make(chan int)
// 	wg.Add(3)

// 	go Producer(chan1, wg)
// 	go Processor(chan1, chan2, wg)
// 	go Consumer(chan2, wg)

// 	// wg.Add(1)
// 	// go func() {
// 	// 	Producer(wg)
// 	// 	wg.Done()
// 	// }()
// 	wg.Wait()

// }

// func Producer(out chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	defer close(out)
// 	for i := 0; i < 20; i++ {
// 		out <- i
// 	}
// }

// func Processor(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	defer close(out)
// 	for data := range in {
// 		out <- data * data
// 	}

// }

// func Consumer(chan3 <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := range chan3 {
// 		if i%5 == 0 {
// 			fmt.Println(i)
// 		}
// 	}

// }
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	chan1 := make(chan int)
// 	wg := new(sync.WaitGroup)
// 	slice1 := []interface{}{10, 10.23, 1047.2, 8, 7, 9}

// 	wg.Add(1)
// 	go RecieveData(chan1, wg)

// 	for i := 0; i < len(slice1); i++ {
// 		switch n := slice1[i].(type) {
// 		case int:
// 			n1 := float64(n)
// 			sqaure := n1 * n1
// 			chan1 <- int(sqaure)
// 		case float64:
// 			squareFloat := n * n
// 			chan1 <- int(squareFloat)

// 		}
// 	}

// }

// func RecieveData(in <-chan int, wg *sync.WaitGroup) {
// 	for i := range in {
// 		fmt.Println(i)
// 	}
// 	defer wg.Done()
// }

package main

import "fmt"

func main() {

	slice1 := []interface{}{10, 10.23, 1047.2, 8, 7, 9}
	for{
		switch {

		}
	}

}
