package main

import (
	"fmt"
	"strings"
)

func is_pal(str string) bool {
	s := []string{}
	for _, letter := range str {
		if string(letter) != " " {
			s = append(s, string(letter))
		}
	}
	stroka := strings.Join(s, "")
	if len(stroka) <= 1 {
		return true
	}
	if stroka[0] == stroka[len(stroka)-1] {
		return is_pal(stroka[1 : len(stroka)-1])
	}
	return false
}

func main() {
	str := "ab w      ba"
	fmt.Println(is_pal(str))
}
