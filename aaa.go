package main

import (
	"fmt"
	"unicode/utf8"
)

func getMaxLenthStrCount(s string) int {
	list := make(map[rune]int)
	start := 0
	lenth := 0
	for i, ch := range []rune(s) {
		lastI, ok := list[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > lenth {
			lenth = i - start + 1
		}
		list[ch] = i
	}
	return lenth
}
func main() {
	s := "Yes, 我是谁!"
	for i, ch := range s {
		fmt.Println(i, ch)
	}
	fmt.Println(leng("abcaaa"))
	fmt.Println(leng("我是谁我都不知道啊啊留我阿飞咯范围"))

	fmt.Println(utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
}
