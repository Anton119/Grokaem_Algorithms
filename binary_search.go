package main

import (
	"fmt"
	"sort"
)

func binarySearch(list []int, item int) int {
	sortedArr := sort.IntSlice(list)
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := sortedArr[mid]
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else if guess < item {
			low = mid + 1
		}
	}
	return 0
}

func main() {
	arr := []int{1, 4, 7, 8, 11}
	item := 8
	fmt.Println(binarySearch(arr, item))
}
