package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes我爱写代码!" //UTF-8
	fmt.Println(len(s))
	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d, %x) ", i, ch)
	}
	fmt.Println()

	fmt.Println("rune count : ", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
}
