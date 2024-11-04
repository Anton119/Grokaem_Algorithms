package main

import "fmt"

func summation(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + summation(n-1)
	}
}

func main() {
	n := 4
	fmt.Println(summation(n))
}
