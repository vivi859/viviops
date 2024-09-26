package main

import "fmt"

func main() {
	fmt.Println("fibonacci 0-10 sum:")
	for i := 1; i <= 10; i++ {
		sum1 := fibonacci(i)
		fmt.Println(sum1)
	}
}

func fibonacci(n int) int {
	// TODO
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
