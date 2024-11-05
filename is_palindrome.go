package main

import "fmt"

func is_paoindrrome(str string) bool {
	if len(str) <= 1 {
		return true
	}
	if str[0] == str[len(str)-1] {
		return is_paoindrrome(str[1 : len(str)-1])
	}

	return false
}

func main() {
	str := "abwba"
	fmt.Println(is_paoindrrome(str))
}
