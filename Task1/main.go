package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	map1 := map[string]float64{
		"AAPL": 172.56,
		"GOOG": 2875.22,
		"INFY": 1400.03,
	}

	chan1 := make(chan map[string]float64)
	done := make(chan struct{})

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				for k, v := range map1 {
					change := rand.Float64() * 2
					newPrice := v + change
					map1[k] = newPrice
				}
				chan1 <- map1
			case <-done:
				close(chan1)
				return
			}
		}
	}()

	go func() {
		for prices := range chan1 {
			now := time.Now().Format("15:04:05")
			for ticker, price := range prices {
				fmt.Printf("[%s] %s: %.2f\n", now, ticker, price)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	ticker.Stop()
	close(done)

}
