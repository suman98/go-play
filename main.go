package main

import (
	"fmt"
	"test/technicalindicator"
)

func main() {
	// fmt.Println("Hello, Go Modules!")
	duration, last := technicalindicator.StressTestSampleRSI(1000000, 1000000)
	// panic("Stress test done: last RSI=%.4f")
	// fmt.Println(practice.Main())
	fmt.Printf("Stress test done: duration=%v, last RSI=%.4f\n", duration, last)
}
