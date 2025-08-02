package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			count++
		}
	}
	return count
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SpinWords(str string) string {
	spaceSeparated := strings.Fields(str)
	var result []string
	for i := 0; i < len(spaceSeparated); i++ {
		if len(spaceSeparated[i]) > 4 {
			result = append(result, Reverse(spaceSeparated[i]))
		} else {
			result = append(result, spaceSeparated[i])
		}
	}

	joinedString := strings.Join(result, " ")
	return joinedString
}

func main() {
	fmt.Println(SpinWords("This is a testing of reverse string"))
	fmt.Println(GetCount("this is a test for how many vowel."))
}
