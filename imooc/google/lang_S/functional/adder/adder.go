package main

import "fmt"

func adder() func(int) int {
	// 自由变量sum
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d\n", i, a(i))
	}
}