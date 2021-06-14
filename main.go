package main

import "fmt"

func main() {
	fmt.Println("hello world")
	s := 3
	switch s {
	case 3:
		fmt.Println(" i am 3 ")
	default:
		fmt.Println("I am not 3")
	}
}
