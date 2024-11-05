package main

import (
	"fmt"
	"strings"
)

func removeVowels(str string) string {
	VOWELS := []string{"a", "o", "e", "i", "y", "u"}
	stroka := []string{}
	if len(str) == 0 {
		return ""
	}
	for _, l := range str {
		stroka = append(stroka, string(l))
	}
	if stroka[0] != VOWELS[0] {
		return string(str[0]) + removeVowels(str[1:])
	} else {
		return removeVowels(string(str[1:]))
	}
	s := strings.Join(stroka, "")
	return s
}

func main() {
	str := "apple"
	fmt.Println(removeVowels(str))

}
