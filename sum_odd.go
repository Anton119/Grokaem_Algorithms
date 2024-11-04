package main

import "fmt"

func sum_odd(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if arr[0]%2 != 0 {
		return arr[0] + sum_odd(arr[1:])
	} else {
		return sum_odd(arr[1:])
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(sum_odd(arr))
}
