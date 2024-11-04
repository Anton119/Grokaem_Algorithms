package main

import "fmt"

func sum_sub(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if arr[0]%2 == 1 {
		return sum_sub(arr[1:]) + arr[0]
	} else {
		return sum_sub(arr[1:]) - arr[0]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(sum_sub(arr))
}
