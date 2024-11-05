package main

import "fmt"

func doubleString(str string) string {
	if len(str) == 0 {
		return ""
	}
	return string(str[0]) + string(str[0]) + doubleString(str[1:])
}

func main() {
	str := "apple"
	fmt.Println(doubleString(str))
}
