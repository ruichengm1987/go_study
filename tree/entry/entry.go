package main

import "fmt"

func a() string {
	/* doesn't matter */
	return ""
}

func main() {
	//tmp := a()
	//b := &tmp
	//
	b := &a()
	fmt.Println(b)
}
