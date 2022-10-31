package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "好家伙"
	fmt.Println(s)
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("rune count", utf8.RuneCountInString(s))
	fmt.Println()

	for idx, rune := range s {
		fmt.Printf("%d %c", idx, rune)
	}

	for i, w := 0, 0; i < len(s); i += w {
		rune, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c", i, rune)
		w = width

	}
}
