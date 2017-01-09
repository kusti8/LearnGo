package main

import "fmt"

func main() {
	s := make([][]uint8, 5)
	for i := range s {
		s[i] = make([]uint8, 3)
	}
	fmt.Println(s) // test
}
