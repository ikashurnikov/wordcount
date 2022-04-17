package main

import (
	"fmt"
	"os"
	"unicode"
)

func skip(str []rune, from int, cond func(ch rune) bool) int {
	for i := from; i < len(str); i++ {
		if !cond(str[i]) {
			return i
		}
	}
	return len(str)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid arguments count!")
		os.Exit(1)
	}

	str := []rune(os.Args[1])
	word_count := 0

	for i := 0; i < len(str); {
		i = skip(str, i, unicode.IsSpace)
		if i == len(str) {
			break
		}
		word_count++
		i = skip(str, i, func(ch rune) bool {
			return !unicode.IsSpace(ch)
		})
	}

	fmt.Println(word_count)
}
